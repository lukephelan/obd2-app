package main

import (
	"fmt"

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
		v.Highlight = true
		v.SelFgColor = gocui.ColorGreen
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
	for i, item := range currentMenu {
		prefix := "  "
		if i == selectedIndex {
			prefix = "→ "
		}

		if item.IsHeading {
			fmt.Fprintln(v, fmt.Sprintf("── %s ──", item.Name))
		} else {
			fmt.Fprintln(v, prefix+item.Name)
		}
	}
}

func updateDataView(g *gocui.Gui) {
	v, err := g.View("data")
	if err != nil {
		return // Avoid crashing if the view isn't available
	}
	v.Clear()

	if showLiveData {
		// Display OBD2 live data
		fmt.Fprintf(v, " RPM: %d\n", rpm)
		fmt.Fprintf(v, " Voltage: %.1fV\n", voltage)
		fmt.Fprintf(v, " Speed: %dkm/h\n", speed)
	} else {
		fmt.Fprintln(v, controlsText)
	}
}
