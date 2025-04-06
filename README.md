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
git clone https://github.com/lukephelan/obd2-app.git
cd obd2-app
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

### 3. Build & Run the Application

#### Using Makefile (Recommended)

If you have `make` installed, use:

```sh
make build      # Compiles the binary
make run        # Runs the application
make clean      # Removes build artifacts
make test       # Runs unit tests
make coverage   # Runs unit tests and provides a coverage report
```

#### Manual Build (Without Make)

If you don’t have `make`, you can build manually:

```sh
go build -o obd2-app ./cmd
./obd2-app
```

## 🔽 Download the Latest Binary

You can download a prebuilt version of `obd2-app` without needing to install Go.

### **1️⃣ Go to the GitHub Actions Page**

1. Click on the **"Actions"** tab in this repository.
2. Select **"Build and Upload Binary"** from the list of workflows.
3. Click on the latest run (usually at the top).

### **2️⃣ Download the Binary**

1. Scroll down to the **"Artifacts"** section.
2. Click **"obd2-app-binary"** to download it.
3. If needed, make it executable:
   ```sh
   chmod +x obd2-app
   ```

### **3️⃣ Run the App**

Now you can run the app without installing anything:

```sh
./obd2-app
```

🚀 **You're ready to use OBD2-APP!**

## 🛠 Future Improvements

- ✅ Add support for reading OBD2 sensor data via USB/Bluetooth
- ✅ Implement real-time data graphing
- ✅ Add error code lookup functionality

## 🐝 License

This project is open-source under the **MIT License**.
