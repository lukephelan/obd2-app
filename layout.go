package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// Create Menu View
	if v, err := g.SetView("menu", 0, 0, maxX/3, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " OBD2 Menu "
		v.Wrap = true
		renderMenu(g)
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

	updateDataView(g)
	return nil
}

func renderMenu(g *gocui.Gui) {
	v, err := g.View("menu")
	if err != nil {
		return
	}
	v.Clear()

	fmt.Fprint(v, "\x1b[0m")

	for i, item := range currentMenu {
		prefix := "  "

		if item.IsHeading {
			// Print subheading in green
			fmt.Fprintf(v, "\x1b[32m──%s──\x1b[0m\n", item.Name)
		} else if i == selectedIndex || (i == 0 && selectedIndex == 0) {
			// Ensure highlight is always applied, even on first render
			fmt.Fprintf(v, "\x1b[0m\x1b[7m%s%s\x1b[0m\n", prefix, item.Name)
		} else {
			// Normal menu items
			fmt.Fprintf(v, "\x1b[0m%s%s\n", prefix, item.Name)
		}
	}
}

func updateDataView(g *gocui.Gui) {
	log.Printf("Updating data view (showLiveData: %t)", showLiveData)
	v, err := g.View("data")
	if err != nil {
		return // Avoid crashing if the view isn't available
	}
	v.Clear()

	if showLiveData {
		// Display placeholder until real OBD2 integration
		fmt.Fprintln(v, "🚧 Not Yet Implemented 🚧")
		fmt.Fprintln(v, "This feature will be available in a future update.")
	} else {
		fmt.Fprintln(v, controlsText)
	}
}
