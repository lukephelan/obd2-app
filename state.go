package main

type MenuItem struct {
	Name      string
	SubMenu   []*MenuItem
	IsHeading bool // Non-selectable if true
}

var showLiveData = false

var controlsText = `
OBD2 Menu Navigation:
---------------------
↑ ↓       : Move Up / Down
←         : Go Back
→ / ENTER : Select Item
ESC       : Exit Menu
CTRL+C    : Quit Program
`

var (
	selectedIndex = 0
	currentMenu   []*MenuItem
	menuHistory   [][]*MenuItem
	indexHistory  []int
)

var menu = []*MenuItem{
	{Name: "General OBD2 Options", SubMenu: []*MenuItem{
		{Name: "Check Engine Light Status"},
		{Name: "Read Engine Trouble Codes"},
		{Name: "Read Pending Trouble Codes"},
		{Name: "Read Permanent Trouble Codes"},
		{Name: "Clear Trouble Codes"},
		{Name: "Reset Check Engine Light"},
		{Name: "Freeze Frame Data"},
	}},
	{Name: "Live Sensor Data", SubMenu: []*MenuItem{
		{Name: "Engine & Performance", IsHeading: true},
		{Name: "Engine RPM"},
		{Name: "Vehicle Speed"},
		{Name: "Throttle Position"},
		{Name: "Mass Air Flow (MAF)"},
		{Name: "Short-Term Fuel Trim (STFT)"},
		{Name: "Long-Term Fuel Trim (LTFT)"},

		{Name: "Temperatures", IsHeading: true},
		{Name: "Engine Coolant Temperature"},
		{Name: "Intake Air Temperature"},
		{Name: "Oil Temperature"},
		{Name: "Transmission Temperature"},
		{Name: "Exhaust Gas Temperature"},
	}},
	{Name: "Advanced Diagnostics & Controls", SubMenu: []*MenuItem{
		{Name: "Actuator & System Tests", IsHeading: true},
		{Name: "Evaporative System Test"},
		{Name: "Oxygen Sensor Response Test"},
		{Name: "Throttle Body Relearn"},
		{Name: "ECU Reset/Reboot"},

		{Name: "Hybrid/Electric Vehicle Data", IsHeading: true},
		{Name: "Battery Pack Voltage"},
		{Name: "Battery Cell Temperature"},
		{Name: "Charging Status"},
	}},
	{Name: "Transmission & Drivetrain", SubMenu: []*MenuItem{
		{Name: "Transmission Gear Position"},
		{Name: "Torque Converter Lockup Status"},
		{Name: "Wheel Speed Sensors"},
	}},
	{Name: "Suspension, Brakes & Chassis", SubMenu: []*MenuItem{
		{Name: "Brake Pedal Position"},
		{Name: "ABS Activation Status"},
		{Name: "Tire Pressure Monitoring System (TPMS)"},
	}},
	{Name: "Vehicle Information & Configuration", SubMenu: []*MenuItem{
		{Name: "VIN & ECU Details", IsHeading: true},
		{Name: "Vehicle Identification Number (VIN)"},
		{Name: "ECU Software Version"},
		{Name: "ECU Manufacturer"},
	}},
}

func init() {
	currentMenu = menu
}
