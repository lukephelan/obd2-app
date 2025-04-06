package obd2

import "testing"

func TestGetMockResponse(t *testing.T) {
	tests := []struct {
		cmd      string
		expected string
	}{
		{"ATRV", "12.6V"},
		{"010C", "41 0C 1F A0"},
		{"UNKNOWN", "N/A"},
	}

	for _, tt := range tests {
		result := getMockResponse(tt.cmd)
		if result != tt.expected {
			t.Errorf("getMockResponse(%s) = %s; want %s", tt.cmd, result, tt.expected)
		}
	}
}
