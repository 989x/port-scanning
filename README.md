# Go Port Monitor

A simple Go application to monitor active ports and the processes using them on a server running Ubuntu.

## Features

- Lists all active ports on the server.
- Displays process details including the command, PID, user, and port information.
- Uses the `lsof` and `ss` commands to gather port and process data.
- Works on Ubuntu and other Linux distributions.

## Requirements

- Go 1.18 or higher
- `lsof` and `ss` commands should be available on the system

## Installation

1. Build the Go application:
   ```bash
   go build -o port-monitor main.go
   ```

2. Run the application:
   ```bash
   ./port-monitor
   ```

## Usage

- The application will display a list of active ports along with the process details such as the command, PID, user, and port information.

Example output:
```
COMMAND   PID   USER   FD     TYPE    DEVICE   SIZE/OFF   NODE     NAME
sshd      1256  root   3u     IPv4    12345    0t0        TCP      *:22
apache2   2246  www    4u     IPv6    67891    0t0        TCP      *:80
```
