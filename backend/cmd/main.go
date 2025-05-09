package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jroimartin/gocui"
	"github.com/lukephelan/obd2-app/backend/internal/obd2"
	"github.com/lukephelan/obd2-app/backend/internal/state"
	"github.com/lukephelan/obd2-app/backend/internal/ui"
)

var logFile *os.File
var adapter *obd2.Adapter

func init() {
	var err error
	logFile, err = os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(logFile)
	log.Println("===== Application Started =====")
}

func enterMenu(g *gocui.Gui) error {
	item := state.CurrentMenu[state.SelectedIndex]

	if item.SubMenu != nil {
		// Save current menu state before navigating deeper
		state.MenuHistory = append(state.MenuHistory, state.CurrentMenu)
		state.IndexHistory = append(state.IndexHistory, state.SelectedIndex)

		// Enter submenu
		state.CurrentMenu = item.SubMenu
		state.SelectedIndex = 0 // Reset selection

		// Ensure we don't land on a heading
		for state.SelectedIndex < len(state.CurrentMenu) && state.CurrentMenu[state.SelectedIndex].IsHeading {
			state.SelectedIndex++ // Move to first non-heading item
		}

		state.ShowLiveData = false

	} else {
		state.ShowLiveData = true
	}
	ui.RenderMenu(g)
	ui.UpdateDataView(g)
	return nil
}

func exitMenu(g *gocui.Gui) error {
	if len(state.MenuHistory) > 0 {
		state.CurrentMenu = state.MenuHistory[len(state.MenuHistory)-1]
		state.SelectedIndex = state.IndexHistory[len(state.IndexHistory)-1]

		state.MenuHistory = state.MenuHistory[:len(state.MenuHistory)-1]
		state.IndexHistory = state.IndexHistory[:len(state.IndexHistory)-1]

		state.ShowLiveData = false // Restore controls view when going back
	}

	ui.RenderMenu(g)
	ui.UpdateDataView(g)
	return nil
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("menu", gocui.KeyArrowUp, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return ui.MoveSelection(g, -1)
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("menu", gocui.KeyArrowDown, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return ui.MoveSelection(g, 1)
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("menu", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return enterMenu(g)
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("menu", gocui.KeyArrowRight, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return enterMenu(g)
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("menu", gocui.KeyArrowLeft, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return exitMenu(g)
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("menu", gocui.KeyEsc, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return exitMenu(g)
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	}); err != nil {
		return err
	}

	return nil
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func startHTTPServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Hello from OBD2 TUI!"}`))
	})

	log.Println("Starting HTTP server on :8080")
	if err := http.ListenAndServe(":8080", enableCORS(mux)); err != nil {
		log.Fatal("HTTP server failed:", err)
	}
}

func main() {
	// Can remove separate go routine once we remove TUI functionality
	go startHTTPServer()

	// Try connecting to the OBD2 adapter
	var err error
	adapter, err = obd2.NewAdapter("/dev/tty.usbserial-A79B4CMW") // FIXME: Need a reusable solution for setting portName
	if err != nil {
		log.Panicln("❌ Failed to initialize OBD2 adapter:", err)
	}

	defer adapter.Close()  // Ensure cleanup on exit
	ui.SetAdapter(adapter) // Pass the adapter to UI

	// TODO: This needs to be scalable
	state.ReadBatteryVoltage = func() { ui.UpdateBatteryVoltage(ui.GetGuiInstance()) }
	state.ReadRPM = func() { ui.UpdateRPM(ui.GetGuiInstance()) }

	// Start TUI
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(ui.Layout)
	ui.UpdateDataView(g)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
