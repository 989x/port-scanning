package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
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

// ParseData converts a line of process data into a Process struct
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

// DisplayTable groups and formats process information
func DisplayTable(processes []Process) {
	commandGroups := make(map[string][]Process)

	// Group processes by command
	for _, p := range processes {
		commandGroups[p.Command] = append(commandGroups[p.Command], p)
	}

	// Color codes for formatting
	brightWhite := "\033[1;37m"
	regularWhite := "\033[0;37m"
	lightGreen := "\033[1;32m"
	resetColor := "\033[0m"

	// Display grouped command summaries with detailed info
	for command, group := range commandGroups {
		fmt.Printf("%s%s found %d entries%s\n", brightWhite, command, len(group), resetColor)

		for _, p := range group {
			// Determine protocol type from name or type field
			protocol := "Unknown"
			if strings.Contains(p.Name, "TCP") {
				protocol = "TCP"
			} else if strings.Contains(p.Name, "UDP") {
				protocol = "UDP"
			} else if p.Type == "IPv4" || p.Type == "IPv6" {
				protocol = "IP"
			}

			fmt.Printf(" %s%s%s - %s%s %s on %s%s\n",
				regularWhite, p.Command, resetColor, lightGreen, p.Type, protocol, p.Name, resetColor)

			fmt.Printf(" %s • PID: %s | User: %s | Node: %s | FD: %s | Size: %s%s\n",
				regularWhite, p.PID, p.User, p.Node, p.FD, p.SizeOff, resetColor)
		}
		fmt.Println() // Space between groups
	}
}

func main() {
	startTime := time.Now() // Start time for duration calculation

	// Run `lsof` to gather process information
	out, err := exec.Command("lsof", "-i").Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// Process output into structured data
	lines := strings.Split(string(out), "\n")
	var processes []Process
	for _, line := range lines[1:] {
		if line != "" {
			processes = append(processes, ParseData(line))
		}
	}

	// Display header with report details
	totalProcesses := len(processes)
	duration := time.Since(startTime)
	fmt.Println()
	fmt.Printf("Program: Process Summary Report\n")
	fmt.Printf("Generated on: %s\n", startTime.Format("Monday, January 2, 2006 - 15:04:05 MST"))
	fmt.Printf("Duration: %v\n", duration)
	fmt.Printf("Total Processes Found: %d\n\n", totalProcesses)

	// Output process table
	DisplayTable(processes)
}
