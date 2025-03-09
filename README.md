# OBD2 Terminal UI (gocui-based)

This is a **Terminal UI (TUI)** application for interacting with an **OBD2 adapter**, built using [gocui](https://github.com/jroimartin/gocui). It allows you to navigate through menus and view real-time vehicle data.

## 🚀 Features

- **OBD2 Menu Navigation** – Browse vehicle diagnostics data
- **Live Sensor Data** – View RPM, voltage, and speed in real-time
- **Interactive Controls** – Navigate using arrow keys & Enter
- **Submenus** – Drill down into detailed sensor readings

## 🎮 Controls

```
↑ ↓         : Move Up / Down
←           : Go Back
→ / ENTER   : Select Item
ESC         : Exit Menu
CTRL+C      : Quit Program
```

## 📚 Installation

### 1. Clone the Repository

```sh
git clone https://github.com/lukephelan/obd2-tui.git
cd obd2-tui
```

### 2. Install Dependencies

Ensure you have **Go installed** (version 1.18+):

```sh
go version
```

Then install dependencies:

```sh
go mod tidy
```

### 3. Run the App

Run the application with:

```sh
go run ./cmd
```

## 🛠 Future Improvements

- ✅ Add support for reading OBD2 sensor data via USB/Bluetooth
- ✅ Implement real-time data graphing
- ✅ Add error code lookup functionality

## 🐝 License

This project is open-source under the **MIT License**.
