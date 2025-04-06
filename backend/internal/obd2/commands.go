package obd2

import (
	"fmt"
	"log"
)

// OBD2Command represents an OBD2 PID command and its parser function
type OBD2Command struct {
	PID         string
	Description string
	ParseFunc   func(string) (float64, error)
}

// Define a lookup map where keys are **logical names**, and values contain OBD2 PIDs
var OBD2Commands = map[string]OBD2Command{
	"Voltage": {
		PID:         "0142",
		Description: "Control Module Voltage (Battery Voltage)",
		ParseFunc:   ParseBatteryVoltage,
	},
	"RPM": {
		PID:         "010C",
		Description: "Engine RPM",
		ParseFunc: func(response string) (float64, error) {
			rpm, err := ParseRPM(response)
			return float64(rpm), err // Convert to float64 for consistency
		},
	},
}

// RunOBD2Command takes a logical name (e.g., "RPM")
func (a *Adapter) RunOBD2Command(commandName string) (float64, error) {
	cmd, exists := OBD2Commands[commandName]
	if !exists {
		return 0, fmt.Errorf("❌ Unknown command: %s", commandName)
	}

	log.Printf("📡 Sending OBD2 command: %s (%s)", cmd.PID, cmd.Description)

	// Send the command
	response, err := a.SendCommand(cmd.PID)
	if err != nil {
		return 0, err
	}

	log.Printf("✅ Raw Response for %s: %q", cmd.PID, response)

	// Parse the response
	value, err := cmd.ParseFunc(response)
	if err != nil {
		return 0, fmt.Errorf("❌ Failed to parse response for %s: %w", cmd.PID, err)
	}

	log.Printf("✅ Parsed Value for %s: %.2f", cmd.PID, value)
	return value, nil
}

// GetBatteryVoltage retrieves the vehicle's battery voltage
func (a *Adapter) GetBatteryVoltage() (float64, error) {
	return a.RunOBD2Command("Voltage")
}

// GetRPM retrieves the engine RPM
func (a *Adapter) GetRPM() (int, error) {
	value, err := a.RunOBD2Command("RPM")
	return int(value), err // Convert back to int
}
