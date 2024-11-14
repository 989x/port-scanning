package ports

import (
	"bytes"
	"fmt"
	"os/exec"
	"port-scanning/internal/colors"
	"strings"
)

func CheckActivePorts() (int, error) {
	cmd := exec.Command("lsof", "-i", "-P", "-n")
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command and check for errors
	err := cmd.Run()
	if err != nil {
		return 0, err
	}

	// Add a newline after loading spinner is cleared
	fmt.Println()

	// Process output
	lines := strings.Split(out.String(), "\n")
	fmt.Printf("%s%-15s %-10s %-10s %-5s %-10s %-10s %-10s %-10s %s%s\n",
		colors.ColorCyan, "COMMAND", "PID", "USER", "FD", "TYPE", "DEVICE", "SIZE/OFF", "NODE", "NAME", colors.ColorReset)
	fmt.Println(colors.ColorGray + strings.Repeat("─", 90) + colors.ColorReset)

	// Continue with the rest of the function...
	activeConnections := 0
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) != "" {
			fields := strings.Fields(line)
			if len(fields) >= 9 {
				fmt.Printf("%s%-15s %s%-10s %s%-10s %s%-5s %s%-10s %s%-10s %s%-10s %s%-10s %s%s\n",
					colors.ColorGreen, fields[0], // COMMAND
					colors.ColorRed, fields[1], // PID
					colors.ColorBlue, fields[2], // USER
					colors.ColorGray, fields[3], // FD
					colors.ColorGray, fields[4], // TYPE
					colors.ColorGray, fields[5], // DEVICE
					colors.ColorGray, fields[6], // SIZE/OFF
					colors.ColorGray, fields[7], // NODE
					colors.ColorReset, fields[8], // NAME
				)
				activeConnections++
			}
		}
	}
	fmt.Println(colors.ColorGray + strings.Repeat("─", 90) + colors.ColorReset)
	return activeConnections, nil
}
