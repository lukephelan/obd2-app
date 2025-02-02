package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

func moveSelection(g *gocui.Gui, delta int) error {
	for {
		selectedIndex = (selectedIndex + delta + len(currentMenu)) % len(currentMenu)
		if !currentMenu[selectedIndex].IsHeading {
			break
		}
	}
	renderMenu(g)
	return nil
}

func enterMenu(g *gocui.Gui) error {
	item := currentMenu[selectedIndex]
	if item.SubMenu != nil {
		// Navigating into a submenu
		menuHistory = append(menuHistory, currentMenu)
		indexHistory = append(indexHistory, selectedIndex)
		currentMenu = item.SubMenu
		selectedIndex = 0
		showLiveData = false // Keep coontrols view when entering a submenu
	} else {
		// Selecting an actual option → Switch to OBD2 data
		showLiveData = true
	}

	renderMenu(g)
	updateDataView(g)
	return nil
}

func exitMenu(g *gocui.Gui) error {
	if len(menuHistory) > 0 {
		currentMenu = menuHistory[len(menuHistory)-1]
		selectedIndex = indexHistory[len(indexHistory)-1]

		menuHistory = menuHistory[:len(menuHistory)-1]
		indexHistory = indexHistory[:len(indexHistory)-1]

		showLiveData = false // Restore controls view when going back
	}

	renderMenu(g)
	updateDataView(g)
	return nil
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("menu", gocui.KeyArrowUp, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return moveSelection(g, -1)
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("menu", gocui.KeyArrowDown, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return moveSelection(g, 1)
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

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
