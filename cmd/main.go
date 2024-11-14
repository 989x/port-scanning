package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// Process represents each row in the table
type Process struct {
	Command string
	PID     string
	User    string
	FD      string
	Type    string
	Device  string
	SizeOff string
	Node    string
	Name    string
}

// ParseData parses each line of data into a Process struct
func ParseData(line string) Process {
	fields := strings.Fields(line)
	return Process{
		Command: fields[0],
		PID:     fields[1],
		User:    fields[2],
		FD:      fields[3],
		Type:    fields[4],
		Device:  fields[5],
		SizeOff: fields[6],
		Node:    fields[7],
		Name:    fields[8],
	}
}

// DisplayTable displays the processes in the "Bracketed Process Summary" style with grouping by command
func DisplayTable(processes []Process) {
	var lastCommand string

	for _, p := range processes {
		// Determine protocol type from the command name and network type
		protocol := "Unknown"
		if strings.Contains(p.Name, "TCP") {
			protocol = "TCP"
		} else if strings.Contains(p.Name, "UDP") {
			protocol = "UDP"
		}

		// Check if we are in a new group (new command)
		if p.Command != lastCommand {
			// Add a newline to separate groups if this is not the first group
			if lastCommand != "" {
				fmt.Println()
			}
			lastCommand = p.Command
		}

		// Print the bracketed summary line
		fmt.Printf("[ %s - %s %s on %s ]\n", p.Command, p.Type, protocol, p.Name)

		// Print the detailed information line
		fmt.Printf("  PID: %s | User: %s | Node: %s | FD: %s | Size: %s\n",
			p.PID, p.User, p.Node, p.FD, p.SizeOff)
	}
}

func main() {
	// Run the lsof command to get process information
	out, err := exec.Command("lsof", "-i").Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// Split output into lines
	lines := strings.Split(string(out), "\n")

	// Skip header and parse each line
	var processes []Process
	for _, line := range lines[1:] {
		if line != "" {
			processes = append(processes, ParseData(line))
		}
	}

	// Display the table in "Bracketed Process Summary" style with grouping
	DisplayTable(processes)
}
