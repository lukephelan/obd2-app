package obd2

import (
	"errors"
	"testing"
)

type MockSerialPort struct {
	response string
	writeErr error
	readErr  error
}

func (m *MockSerialPort) Write(p []byte) (n int, err error) {
	if m.writeErr != nil {
		return 0, m.writeErr
	}
	return len(p), nil
}

func (m *MockSerialPort) Read(p []byte) (n int, err error) {
	if m.readErr != nil {
		return 0, m.readErr
	}
	copy(p, m.response)
	return len(m.response), nil
}

func (m *MockSerialPort) Close() error {
	return nil
}

func TestNewAdapter_MockMode(t *testing.T) {
	adapter, err := NewAdapter("COM1")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !adapter.isMock {
		t.Errorf("Expected adapter to be in mock mode, got isMock=%v", adapter.isMock)
	}
}

func TestAdapter_Close(t *testing.T) {
	adapter := &Adapter{isMock: true}
	adapter.Close()
}

func TestAdapter_SendCommand_MockMode(t *testing.T) {
	adapter := &Adapter{isMock: true}
	resp, err := adapter.SendCommand("010C")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := getMockResponse("010C")
	if resp != expected {
		t.Errorf("Expected mock response %q, got %q", expected, resp)
	}
}

func TestAdapter_SendCommand_NoPort(t *testing.T) {
	adapter := &Adapter{isMock: false, port: nil}
	_, err := adapter.SendCommand("010C")

	if err == nil || err.Error() != "❌ OBD2 adapter is not connected" {
		t.Errorf("Expected 'OBD2 adapter is not connected' error, got %v", err)
	}
}

func TestAdapter_SendCommand_RealPort(t *testing.T) {
	mockPort := &MockSerialPort{response: "41 0C 1F A0"}
	adapter := &Adapter{isMock: false, port: mockPort}

	resp, err := adapter.SendCommand("010C")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := "41 0C 1F A0"
	if resp != expected {
		t.Errorf("Expected response %q, got %q", expected, resp)
	}
}

func TestAdapter_SendCommand_WriteError(t *testing.T) {
	mockPort := &MockSerialPort{writeErr: errors.New("write failed")}
	adapter := &Adapter{isMock: false, port: mockPort}

	_, err := adapter.SendCommand("010C")
	if err == nil || err.Error() != "❌ Failed to send command: 010C" {
		t.Errorf("Expected 'Failed to send command' error, got %v", err)
	}
}

func TestAdapter_SendCommand_ReadError(t *testing.T) {
	mockPort := &MockSerialPort{readErr: errors.New("read failed")}
	adapter := &Adapter{isMock: false, port: mockPort}

	_, err := adapter.SendCommand("010C")
	if err == nil || err.Error() != "❌ No response received from OBD2 adapter" {
		t.Errorf("Expected 'No response received' error, got %v", err)
	}
}
