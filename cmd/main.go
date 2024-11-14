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

func DisplayTable(processes []Process) {
	commandGroups := make(map[string][]Process)

	// Group processes by command
	for _, p := range processes {
		commandGroups[p.Command] = append(commandGroups[p.Command], p)
	}

	customSymbol := "•"

	brightWhite := "\033[1;37m"  // Bright white for headers
	regularWhite := "\033[0;37m" // Regular white for main text
	lightGreen := "\033[1;32m"   // Light green for details
	resetColor := "\033[0m"      // Reset to default color

	// Display each command group with a summary and detailed information
	for command, group := range commandGroups {
		// Display summary header in bright white
		fmt.Printf("%s%s found %d entries%s\n", brightWhite, command, len(group), resetColor)

		// Display each process in the group with the new format
		for _, p := range group {
			// Determine protocol type from the name or type field
			protocol := "Unknown"
			if strings.Contains(p.Name, "TCP") {
				protocol = "TCP"
			} else if strings.Contains(p.Name, "UDP") {
				protocol = "UDP"
			} else if p.Type == "IPv4" || p.Type == "IPv6" {
				protocol = "IP" // General IP if neither TCP nor UDP is specified
			}

			fmt.Printf(" %s%s%s - %s%s %s on %s%s\n",
				regularWhite, p.Command, resetColor, lightGreen, p.Type, protocol, p.Name, resetColor)

			fmt.Printf(" %s %s PID: %s | User: %s | Node: %s | FD: %s | Size: %s%s\n",
				regularWhite, customSymbol, p.PID, p.User, p.Node, p.FD, p.SizeOff, resetColor)
		}
		fmt.Println() // Add spacing between groups
	}
}

func main() {
	// Record the start time
	startTime := time.Now()

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

	// Enhanced HEAD section with formatting and details
	totalProcesses := len(processes)
	duration := time.Since(startTime)

	// Add a blank line for separation
	fmt.Println()
	fmt.Printf("Program: Process Summary Report\n")
	fmt.Printf("Generated on: %s\n", startTime.Format("Monday, January 2, 2006 - 15:04:05 MST"))
	fmt.Printf("Duration: %v\n", duration)
	fmt.Printf("Total Processes Found: %d\n\n", totalProcesses)

	// Display the table with summaries
	DisplayTable(processes)
}
