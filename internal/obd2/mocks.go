package obd2

import "log"

// getMockResponse returns fake OBD2 responses when in mock mode
func getMockResponse(cmd string) string {
	log.Println("Received command:", cmd)
	switch cmd {
	case "ATRV":
		return "12.6V" // Fake battery voltage
	case "010C":
		return "41 0C 1F A0" // Fake RPM (2000 RPM)
	default:
		return "N/A"
	}
}
