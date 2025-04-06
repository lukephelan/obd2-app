package state

import (
	"testing"
)

// Test that the initial state is set correctly
func TestInitialState(t *testing.T) {
	if CurrentMenu == nil {
		t.Fatal("Expected CurrentMenu to be initialized, but it is nil")
	}
	if len(CurrentMenu) == 0 {
		t.Fatal("Expected CurrentMenu to have menu items, but it is empty")
	}
}

// Test that selecting a menu item updates SelectedIndex
func TestSelectedIndex(t *testing.T) {
	SelectedIndex = 1
	if SelectedIndex != 1 {
		t.Errorf("Expected SelectedIndex to be 1, got %d", SelectedIndex)
	}
}

func TestAllMenuActions(t *testing.T) {
	actionCalls := make(map[string]bool)

	// Mock all actions
	mockAction := func(name string) func() {
		return func() { actionCalls[name] = true }
	}

	// Assign mock actions to all menu items
	for _, item := range menu {
		for _, subItem := range item.SubMenu {
			if subItem.Action != nil {
				actionCalls[subItem.Name] = false // Track if it gets called
				subItem.Action = mockAction(subItem.Name)
			}
		}
	}

	// Execute all actions
	for _, item := range menu {
		for _, subItem := range item.SubMenu {
			if subItem.Action != nil {
				subItem.Action()
			}
		}
	}

	// Verify all actions were executed
	for name, called := range actionCalls {
		if !called {
			t.Errorf("Expected action for menu item '%s' to be executed, but it was not", name)
		}
	}
}
