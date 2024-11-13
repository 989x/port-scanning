package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Use the `lsof` command to check active ports and their processes
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

	// Print column headers for clarity
	fmt.Println("COMMAND\tPID\tUSER\tFD\tTYPE\tDEVICE\tSIZE/OFF\tNODE\tNAME")
	fmt.Println("-------------------------------------------------------------")

	// Print each line which includes process info, PID, and port details
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			fmt.Println(line)
		}
	}
}
