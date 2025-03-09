package obd2

import (
	"errors"
	"log"
	"time"

	"github.com/tarm/serial"
)

type Adapter struct {
	port   *serial.Port
	isMock bool // Flag for mock mode
}

// NewAdapter opens a serial connection to the OBD2 adapter
func NewAdapter(portName string) (*Adapter, error) {
	config := &serial.Config{Name: portName, Baud: 9600, ReadTimeout: time.Second * 2}
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Println("⚠️ No OBD2 adapter detected on", portName, "- running in mock mode.")
		return &Adapter{isMock: true}, nil
	}
	log.Println("✅ Connected to OBD2 adapter:", portName)
	return &Adapter{port: port, isMock: false}, nil
}

// Close closes the connection to the OBD2 adapter
func (a *Adapter) Close() {
	if a.isMock {
		log.Println("🟡 Mock adapter closed.")
	} else if a.port != nil {
		a.port.Close()
		log.Println("✅ Closed OBD2 adapter connection")
	}
}

// SendCommand sends an OBD2 command and returns the response (mock mode returns fake data)
func (a *Adapter) SendCommand(cmd string) (string, error) {
	if a.isMock {
		log.Println("🟡 Mock response for command:", cmd)
		return getMockResponse(cmd), nil
	}

	if a.port == nil {
		return "", errors.New("❌ OBD2 adapter is not connected")
	}

	_, err := a.port.Write([]byte(cmd + "\r")) // Send command with CR
	if err != nil {
		return "", errors.New("❌ Failed to send command: " + cmd)
	}

	// Read response
	buf := make([]byte, 64)
	n, err := a.port.Read(buf)
	if err != nil || n == 0 {
		return "", errors.New("❌ No response received from OBD2 adapter")
	}

	response := string(buf[:n])
	log.Printf("✅ OBD2 Response: %s", response)
	return response, nil
}
