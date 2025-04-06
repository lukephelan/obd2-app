package obd2

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tarm/serial"
)

type SerialPort interface {
	Write(p []byte) (n int, err error)
	Read(p []byte) (n int, err error)
	Close() error
}

type Adapter struct {
	port   SerialPort
	isMock bool
}

// NewAdapter opens a serial connection to the OBD2 adapter
func NewAdapter(portName string) (*Adapter, error) {
	// Normalize the input
	if strings.HasPrefix(portName, "serial://") {
		portName = strings.TrimPrefix(portName, "serial://")
	}

	config := &serial.Config{Name: portName, Baud: 9600, ReadTimeout: time.Second * 2}
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Println("⚠️ No OBD2 adapter detected on", portName, "- running in mock mode.")
		return &Adapter{isMock: true}, nil
	}
	log.Println("✅ Connected to OBD2 adapter:", portName)
	return &Adapter{port: SerialPort(port), isMock: false}, nil
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

	_, err := a.port.Write([]byte(cmd + "\r")) // Send command with CR+
	if err != nil {
		return "", errors.New("❌ Failed to send command: " + cmd)
	}

	// Read response in a loop
	var response []byte
	buf := make([]byte, 64)

	for {
		n, err := a.port.Read(buf)
		if err != nil {
			return "", fmt.Errorf("error reading from adapter: %w", err)
		}
		response = append(response, buf[:n]...)

		// Check for end-of-response condition
		// (for example, a trailing \r or \n, or a specific pattern you expect)
		if len(response) > 3 && response[len(response)-1] == '\r' {
			break
		}
	}

	log.Printf("Raw response: %q", response)

	return string(response), nil
}
