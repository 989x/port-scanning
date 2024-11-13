# Go Port Monitor

A lightweight and efficient Go application designed to monitor active ports and display details about the processes using them on a Linux server, especially optimized for Ubuntu.

## Features

- Lists all active ports and associated process information on the server.
- Displays process details, including:
  - **Command**: The name of the process
  - **PID**: Process ID
  - **User**: User who owns the process
  - **File Descriptor (FD)**, **Type**, **Device**, **Size/Offset**, **Node**, and **Port Name**.
- Uses `lsof` and optionally `ss` commands to retrieve port and process data.
- Works seamlessly on Ubuntu and other Linux distributions.
- Summarizes total active connections at the end of the output for quick analysis.

## Requirements

- Go 1.18 or higher
- `lsof` command should be available on the system (included in most Linux distributions)
- `ss` command is optional but recommended for enhanced network diagnostics

## Installation

1. Clone the repository and navigate to the project directory:
   ```bash
   git clone https://github.com/yourusername/go-port-monitor.git
   cd go-port-monitor
   ```

2. Build the Go application:
   ```bash
   go build -o port-monitor main.go
   ```

3. Run the application:
   ```bash
   ./port-monitor
   ```

## Usage

- When executed, the application will list all active ports along with detailed process information such as the command, PID, user, and port specifics.

- The output is color-coded to enhance readability:
  - **Command** names are highlighted in green.
  - **PID** is displayed in red.
  - **User** details are shown in blue.
  - Other fields, such as **FD**, **Type**, and **Device**, use gray for differentiation.

- **Footer Summary**: The total number of active connections is shown at the bottom for a quick summary.

Example output:
```
COMMAND       PID        USER       FD     TYPE     DEVICE    SIZE/OFF   NODE       NAME
sshd          1256       root       3u     IPv4     12345     0t0        TCP        *:22
apache2       2246       www        4u     IPv6     67891     0t0        TCP        *:80
───────────────────────────────────────────────────────────────────────────────────────────
Total Active Connections: 2
```

## Troubleshooting

- Ensure `lsof` is installed if you encounter errors about missing commands:
  ```bash
  sudo apt-get install lsof
  ```
- If `ss` is also required, install it with:
  ```bash
  sudo apt-get install iproute2
  ```
