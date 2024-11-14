# Go Port Monitor

A lightweight and efficient Go application for monitoring active ports and displaying process details on a Linux server, optimized for Ubuntu.

## Features

- **Displays Active Ports and Processes**: Lists all active ports and associated process information on the server.
- **Process Details**:
  - **Command**: Process name
  - **PID**: Process ID
  - **User**: Process owner
  - **File Descriptor (FD)**, **Type**, **Device**, **Size/Offset**, **Node**, and **Port Name**.
- **Commands Used**: Utilizes `lsof` and optionally `ss` to gather port and process information.
- **Platform Compatibility**: Primarily optimized for Ubuntu but compatible with other Linux distributions.
- **Summarized Output**: Provides a summary section including total active processes, execution time, and timestamp.

## Project Structure

```plaintext
project-root/
├── cmd/
│   └── main.go         # Main entry point of the application
└── go.mod              # Go module file
```

- **`cmd/main.go`**: Coordinates application functionalities.

## Requirements

- **Go** 1.18 or higher
- **lsof**: Essential for retrieving port information (pre-installed on most Linux distributions).
- **ss** (optional): Recommended for enhanced network diagnostics.

## Installation & Usage

1. **Build the Application**:
   ```bash
   go build -o port-monitor cmd/main.go
   ```

2. **Run the Application**:
   ```bash
   ./port-monitor
   ```

## Output Format

The application presents process and port information in the **"Bracketed Process Summary"** format with an informative **HEAD** section.

### Example Output:
```plaintext
Program: Process Summary Report
Generated on: Thursday, November 14, 2024 - 11:38:11 +07
Duration: 135.760208ms
Total Processes Found: 62

sshd found 7 entries
 sshd - IPv4 TCP on *:22
  • PID: 1256 | User: root | Node: 12345 | FD: 3u | Size: 0t0

apache2 found 7 entries
 apache2 - IPv6 TCP on *:80
  • PID: 2246 | User: www-data | Node: 67891 | FD: 4u | Size: 0t0
```

### Output Sections
1. **HEAD Section**: 
   - Displays program name, generation time, execution duration, and total number of processes found.
   
2. **Bracketed Process Summary**: 
   - Lists each process with bracketed command information, followed by PID, user, node, FD, and size.

## Troubleshooting

- **Missing `lsof` Command**: Install with:
  ```bash
  sudo apt-get install lsof
  ```
- **Optional `ss` Command**: Install for extended network details:
  ```bash
  sudo apt-get install iproute2
  ```

This tool provides a straightforward approach for monitoring active connections on your server, making port management efficient and insightful.
