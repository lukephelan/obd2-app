package obd2

import (
	"fmt"
	"log"
	"strings"
)

// ParseBatteryVoltage extracts the battery voltage from an OBD2 response
func ParseBatteryVoltage(response string) (float64, error) {
	response = strings.TrimSpace(response)

	// Expected response: "41 42 XX YY"
	var mode, pid int
	var XX, YY byte

	_, err := fmt.Sscanf(response, "%X %X %X %X", &mode, &pid, &XX, &YY)
	if err != nil {
		return 0, fmt.Errorf("invalid battery voltage response: %s", response)
	}

	// Validate mode and PID
	if mode != 0x41 || pid != 0x42 {
		return 0, fmt.Errorf("unexpected response: mode=%X, pid=%X", mode, pid)
	}

	// Convert to voltage: (XX * 256 + YY) / 1000
	voltage := float64((int(XX)*256 + int(YY))) / 1000
	return voltage, nil
}

// ParseRPM extracts engine RPM from an OBD2 response
func ParseRPM(response string) (int, error) {
	log.Printf("Raw RPM response: %q", response)

	response = strings.Trim(response, ">")
	response = strings.TrimSpace(response)

	// Expected response: "41 0C A B"
	var mode, pid int
	var A, B byte

	_, err := fmt.Sscanf(response, "%X %X %X %X", &mode, &pid, &A, &B)
	log.Printf("Parsed mode: %X, pid: %X, A: %X, B: %X", mode, pid, A, B)

	if err != nil {
		return 0, fmt.Errorf("invalid RPM response format: %s", response)
	}

	// Validate mode and PID
	if mode != 0x41 || pid != 0x0C {
		return 0, fmt.Errorf("unexpected response: mode=%X, pid=%X", mode, pid)
	}

	rpm := ((int(A) * 256) + int(B)) / 4
	return rpm, nil
}
