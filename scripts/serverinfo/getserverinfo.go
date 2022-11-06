package main

import (
	"GoBun/srb2kart/network"
	"fmt"
	"os"
	"strings"
)

func coloriseServerName(name []byte) string {
	output := strings.Builder{}

	// Set the text color to white
	output.WriteString("\x1b[37m")

	// Loop through the bytes
	for i, b := range name {
		// If the byte is over 127, append the equivalent ansi escape code
		if b >= 128 {
			switch name[i] {
			case 128: // White
				output.WriteString("\x1b[37m")
			case 129: // Purple
				output.WriteString("\x1b[35m")
			case 130: // Yellow
				output.WriteString("\x1b[33m")
			case 131: // Green
				output.WriteString("\x1b[32m")
			case 132: // Blue
				output.WriteString("\x1b[34m")
			case 133: // Red
				output.WriteString("\x1b[31m")
			case 134: // Gray
				output.WriteString("\x1b[90m")
			case 135: // Orange
				output.WriteString("\x1b[38;5;208m")
			case 136: // Cyan
				output.WriteString("\x1b[96m")
			case 137: // Lavender
				output.WriteString("\x1b[38;5;183m")
			case 138: // Gold
				output.WriteString("\x1b[38;5;220m")
			case 139: // Lime
				output.WriteString("\x1b[38;5;154m")
			case 140: // Steel
				output.WriteString("\x1b[38;5;145m")
			case 141: // Pink
				output.WriteString("\x1b[38;5;217m")
			case 142: // Brown
				output.WriteString("\x1b[38;5;94m")
			case 143: // Peach
				output.WriteString("\x1b[38;5;215m")
			}
		} else {
			// Otherwise, append the byte
			output.WriteByte(b)
		}
	}

	// Reset the ansi escape code
	output.WriteString("\x1b[0m")

	return output.String()
}

func main() {
	if len(os.Args) < 2 {
		println("Usage: getserverinfo <ip:port>")
		os.Exit(1)
	}

	serverIp := os.Args[1]

	// Add a port if it's not there
	if !strings.Contains(serverIp, ":") {
		serverIp += ":5029"
	}

	// Get the server info
	serverInfo, _, err := network.GetServerInfo(serverIp)
	if err != nil {
		panic(err)
	}

	// Print the server info
	// fmt.Printf("Server info: %+v\n", serverInfo)
	// fmt.Printf("Player info: %+v\n", playerInfo)

	// Print the server name
	fmt.Printf("Server name: %s\n", coloriseServerName(serverInfo.ServerNameRaw))
}
