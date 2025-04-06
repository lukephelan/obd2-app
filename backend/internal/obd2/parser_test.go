package obd2

import (
	"testing"
)

func TestParseBatteryVoltage_Success(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"12.6V", 12.6},
		{"14.2V", 14.2},
		{"10.0V", 10.0},
	}

	for _, tt := range tests {
		voltage, err := ParseBatteryVoltage(tt.input)
		if err != nil {
			t.Fatalf("Unexpected error for input %q: %v", tt.input, err)
		}
		if voltage != tt.expected {
			t.Errorf("For input %q, expected voltage %f, got %f", tt.input, tt.expected, voltage)
		}
	}
}

func TestParseBatteryVoltage_Failure(t *testing.T) {
	tests := []string{
		"12.6",   // Missing 'V'
		"V12.6",  // Incorrect format
		"abcV",   // Non-numeric input
		"12..6V", // Invalid float format
		"12,6V",  // Comma instead of dot
	}

	for _, input := range tests {
		_, err := ParseBatteryVoltage(input)
		if err == nil {
			t.Errorf("Expected error for input %q, but got none", input)
		}
	}
}

func TestParseRPM_Success(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"41 0C 1F A0", 2024},
		{"41 0C 0F 64", 985},
		{"41 0C 2A 30", 2700},
	}

	for _, tt := range tests {
		rpm, err := ParseRPM(tt.input)
		if err != nil {
			t.Fatalf("Unexpected error for input %q: %v", tt.input, err)
		}
		if rpm != tt.expected {
			t.Errorf("For input %q, expected RPM %d, got %d", tt.input, tt.expected, rpm)
		}
	}
}

func TestParseRPM_Failure(t *testing.T) {
	tests := []string{
		"41 0D 1F A0", // Wrong PID (0D instead of 0C)
		"42 0C 1F A0", // Wrong mode (42 instead of 41)
		"41 0C ZZ XX", // Non-hex characters
		"410C1FA0",    // No spaces
		"41 0C 1F",    // Too short
	}

	for _, input := range tests {
		_, err := ParseRPM(input)
		if err == nil {
			t.Errorf("Expected error for input %q, but got none", input)
		}
	}
}
