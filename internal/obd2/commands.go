package obd2

// GetBatteryVoltage retrieves the vehicle's battery voltage
func (a *Adapter) GetBatteryVoltage() (float64, error) {
	response, err := a.SendCommand("ATRV") // ELM327 command for voltage
	if err != nil {
		return 0, err
	}

	return ParseBatteryVoltage(response)
}

// GetRPM retrieves the engine RPM using OBD2 PID 010C.
func (a *Adapter) GetRPM() (int, error) {
	response, err := a.SendCommand("010C") // OBD2 PID for RPM
	if err != nil {
		return 0, err
	}

	return ParseRPM(response)
}
