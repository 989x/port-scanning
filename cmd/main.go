package main

import (
	"fmt"
	"port-watcher/internal/colors"
	"port-watcher/internal/ports"
)

func main() {
	// Display active ports and processes
	activeConnections, err := ports.CheckActivePorts()

	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// Print footer with summary of active connections
	fmt.Printf("%sTotal Active Connections: %d%s\n", colors.ColorCyan, activeConnections, colors.ColorReset)
}
