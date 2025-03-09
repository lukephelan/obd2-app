package obd2

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// ParseBatteryVoltage extracts the battery voltage from an OBD2 response
func ParseBatteryVoltage(response string) (float64, error) {
	response = strings.TrimSpace(response)

	// Validate that the response ends with 'V' (e.g., "12.5V")
	matched, _ := regexp.MatchString(`^\d+(\.\d+)?V$`, response)
	if !matched {
		return 0, errors.New("invalid battery voltage response: " + response)
	}

	// Extract the numeric part and parse it
	voltageStr := strings.TrimSuffix(response, "V")
	var voltage float64
	_, err := fmt.Sscanf(voltageStr, "%f", &voltage)
	if err != nil {
		return 0, fmt.Errorf("failed to parse voltage: %s", response)
	}

	return voltage, nil
}

// ParseRPM extracts engine RPM from an OBD2 response
func ParseRPM(response string) (int, error) {
	// Expected response: "41 0C A B"
	var mode, pid int
	var A, B byte

	_, err := fmt.Sscanf(response, "%X %X %X %X", &mode, &pid, &A, &B)
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
