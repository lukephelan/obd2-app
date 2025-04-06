package obd2

import (
	"errors"
	"testing"
)

type MockAdapter struct {
	mockSendCommand func(cmd string) (string, error)
	response        string
}

func (m *MockAdapter) SendCommand(cmd string) (string, error) {
	return m.mockSendCommand(cmd)
}

func (m *MockAdapter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *MockAdapter) Read(p []byte) (n int, err error) {
	if m.response == "" {
		return 0, errors.New("no data to read") // Simulate empty response
	}
	copy(p, m.response)
	return len(m.response), nil
}

func (m *MockAdapter) Close() error {
	return nil
}

func MockParseBatteryVoltage(response string) (float64, error) {
	if response == "12.6V" {
		return 12.6, nil
	}
	return 0, errors.New("invalid voltage response")
}

func MockParseRPM(response string) (int, error) {
	if response == "41 0C 1F A0" {
		return 2024, nil
	}
	return 0, errors.New("invalid RPM response")
}

func TestGetBatteryVoltage_Success(t *testing.T) {
	mockAdapter := &MockAdapter{
		mockSendCommand: func(cmd string) (string, error) {
			if cmd == "ATRV" {
				return "12.6V", nil
			}
			return "", errors.New("unexpected command")
		},
		response: "12.6V",
	}
	adapter := &Adapter{port: mockAdapter}

	voltage, err := adapter.GetBatteryVoltage()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedVoltage := 12.6
	if voltage != expectedVoltage {
		t.Errorf("Expected voltage %f, got %f", expectedVoltage, voltage)
	}
}

func TestGetBatteryVoltage_ParseError(t *testing.T) {
	mockAdapter := &MockAdapter{
		mockSendCommand: func(cmd string) (string, error) {
			return "invalid response", nil
		},
	}
	adapter := &Adapter{port: mockAdapter}

	_, err := adapter.GetBatteryVoltage()
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}
}

func TestGetRPM_Success(t *testing.T) {
	mockAdapter := &MockAdapter{
		mockSendCommand: func(cmd string) (string, error) {
			if cmd == "010C" {
				return "41 0C 1F A0", nil
			}
			return "", errors.New("unexpected command")
		},
		response: "41 0C 1F A0",
	}
	adapter := &Adapter{port: mockAdapter}

	rpm, err := adapter.GetRPM()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedRPM := 2024
	if rpm != expectedRPM {
		t.Errorf("Expected RPM %d, got %d", expectedRPM, rpm)
	}
}

func TestGetRPM_ParseError(t *testing.T) {
	mockAdapter := &MockAdapter{
		mockSendCommand: func(cmd string) (string, error) {
			return "invalid response", nil
		},
	}
	adapter := &Adapter{port: mockAdapter}

	_, err := adapter.GetRPM()
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}
}
