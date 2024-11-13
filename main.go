package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// ANSI color codes
const (
	ColorReset = "\033[0m"
	ColorRed   = "\033[31m"
	ColorGreen = "\033[32m"
	ColorBlue  = "\033[34m"
	ColorCyan  = "\033[36m"
	ColorGray  = "\033[90m"
)

func main() {
	// Use `lsof` to check active ports and their processes
	cmd := exec.Command("lsof", "-i", "-P", "-n")
	var out bytes.Buffer
	cmd.Stdout = &out

	// Check for errors when running the command
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// Split the output into lines
	lines := strings.Split(out.String(), "\n")

	// Print a colorized header for better readability
	fmt.Printf("%s%-15s %-10s %-10s %-5s %-10s %-10s %-10s %-10s %s%s\n",
		ColorCyan, "COMMAND", "PID", "USER", "FD", "TYPE", "DEVICE", "SIZE/OFF", "NODE", "NAME", ColorReset)
	fmt.Println(ColorGray + strings.Repeat("─", 90) + ColorReset)

	// Count the number of active connections
	activeConnections := 0

	// Print each line with aligned columns and colors
	for _, line := range lines[1:] { // Skip the first line since it's already a header
		if strings.TrimSpace(line) != "" {
			fields := strings.Fields(line)
			if len(fields) >= 9 {
				fmt.Printf("%s%-15s %s%-10s %s%-10s %s%-5s %s%-10s %s%-10s %s%-10s %s%-10s %s%s\n",
					ColorGreen, fields[0], // COMMAND
					ColorRed, fields[1], // PID
					ColorBlue, fields[2], // USER
					ColorGray, fields[3], // FD
					ColorGray, fields[4], // TYPE
					ColorGray, fields[5], // DEVICE
					ColorGray, fields[6], // SIZE/OFF
					ColorGray, fields[7], // NODE
					ColorReset, fields[8], // NAME
				)
				activeConnections++ // Increase count for each valid connection line
			}
		}
	}

	// Print footer with summary of active connections
	fmt.Println(ColorGray + strings.Repeat("─", 90) + ColorReset)
	fmt.Printf("%sTotal Active Connections: %d%s\n", ColorCyan, activeConnections, ColorReset)
}
