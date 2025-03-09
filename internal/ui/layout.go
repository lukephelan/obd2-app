package ui

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/lukephelan/obd2-tui/internal/obd2"
	"github.com/lukephelan/obd2-tui/internal/state"
)

// Reference to the global adapter and GUI instance
var adapter *obd2.Adapter
var guiInstance *gocui.Gui // ✅ Store reference to gocui instance

func SetAdapter(a *obd2.Adapter) {
	adapter = a
}

// GetGuiInstance returns the current GUI instance
func GetGuiInstance() *gocui.Gui {
	return guiInstance
}

// Fetch battery voltage and update UI
func UpdateBatteryVoltage(g *gocui.Gui) {
	if adapter == nil {
		log.Println("⚠️ No OBD2 adapter available.")
		DisplayMessage(g, "Battery Voltage: N/A")
		return
	}

	voltage, err := adapter.GetBatteryVoltage()
	if err != nil {
		log.Println("❌ Failed to get battery voltage:", err)
		DisplayMessage(g, "Battery Voltage: N/A")
		return
	}

	// ✅ Do NOT call UpdateDataView() again here
	DisplayMessage(g, fmt.Sprintf("Battery Voltage: %.2fV", voltage))
}

// Fetch RPM and update UI
func UpdateRPM(g *gocui.Gui) {
	if adapter == nil {
		log.Println("⚠️ No OBD2 adapter available.")
		DisplayMessage(g, "Engine RPM: N/A")
		return
	}

	rpm, err := adapter.GetRPM()
	if err != nil {
		log.Println("❌ Failed to get RPM:", err)
		DisplayMessage(g, "Engine RPM: N/A")
		return
	}

	DisplayMessage(g, fmt.Sprintf("Engine RPM: %d", rpm))
}

// Display a message in the "data" view
func DisplayMessage(g *gocui.Gui, msg string) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("data")
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, msg)
		return nil
	})
}

func Layout(g *gocui.Gui) error {
	guiInstance = g // ✅ Store GUI instance
	maxX, maxY := g.Size()

	// Create Menu View
	if v, err := g.SetView("menu", 0, 0, maxX/3, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " OBD2 Menu "
		v.Wrap = true
		RenderMenu(g)
		if _, err := g.SetCurrentView("menu"); err != nil {
			return err
		}
	}

	// Create Data View (Right Panel)
	if _, err := g.SetView("data", maxX/3+1, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	return nil
}

func RenderMenu(g *gocui.Gui) {
	v, err := g.View("menu")
	if err != nil {
		return
	}
	v.Clear()

	fmt.Fprint(v, "\x1b[0m")

	for i, item := range state.CurrentMenu {
		prefix := "  "

		if item.IsHeading {
			// Print subheading in green
			fmt.Fprintf(v, "\x1b[32m──%s──\x1b[0m\n", item.Name)
		} else if i == state.SelectedIndex || (i == 0 && state.SelectedIndex == 0) {
			// Ensure highlight is always applied, even on first render
			fmt.Fprintf(v, "\x1b[0m\x1b[7m%s%s\x1b[0m\n", prefix, item.Name)
		} else {
			// Normal menu items
			fmt.Fprintf(v, "\x1b[0m%s%s\n", prefix, item.Name)
		}
	}
}

func UpdateDataView(g *gocui.Gui) {
	v, err := g.View("data")
	if err != nil {
		log.Println("❌ View 'data' not found. Skipping update.")
		return
	}
	v.Clear()

	if state.ShowLiveData {
		log.Printf("Updating data view (state.ShowLiveData: %t)", state.ShowLiveData)
		fmt.Fprintln(v, "Fetching live data...")
		if state.SelectedIndex >= 0 && state.SelectedIndex < len(state.CurrentMenu) {
			selectedMenuItem := state.CurrentMenu[state.SelectedIndex]
			log.Println("📢 Selected Menu Item:", selectedMenuItem.Name)
			if selectedMenuItem.Action != nil {
				selectedMenuItem.Action()
			}
		} else {
			fmt.Fprintln(v, "No valid selection.")
			log.Println("❓ Invalid selection index:", state.SelectedIndex)
		}
	} else {
		// ✅ Keep controls text when not showing live data
		fmt.Fprintln(v, state.ControlsText)
	}
}

func MoveSelection(g *gocui.Gui, delta int) error {
	for {
		state.SelectedIndex = (state.SelectedIndex + delta + len(state.CurrentMenu)) % len(state.CurrentMenu)
		if !state.CurrentMenu[state.SelectedIndex].IsHeading {
			break
		}
	}
	RenderMenu(g)
	return nil
}
