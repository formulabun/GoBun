package main

import (
	"GoBun/srb2kart/network"
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const apiUrl = "https://ms.kartkrew.org/ms/api/games/srb2kart/7/servers?v=2"

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
				output.WriteString("\x1b[38;5;129m")
			case 130: // Yellow
				output.WriteString("\x1b[38;5;226m")
			case 131: // Green
				output.WriteString("\x1b[38;5;46m")
			case 132: // Blue
				output.WriteString("\x1b[38;5;21m")
			case 133: // Red
				output.WriteString("\x1b[38;5;196m")
			case 134: // Gray
				output.WriteString("\x1b[38;5;238m")
			case 135: // Orange
				output.WriteString("\x1b[38;5;208m")
			case 136: // Cyan
				output.WriteString("\x1b[38;5;51m")
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


type serverData struct {
	ipAddress string

	serverName string
	serverNameLength int

	ping int

	players int
	maxPlayers int
}


func main() {
	// Get the server list
	result, err := http.Get(apiUrl)
	if err != nil {
		panic(err)
	}

	// Get the result from the server
	scanner := bufio.NewScanner(result.Body)
	serverEntries := []string{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) >= 2 && len(line) <= 3 {
			serverEntries = append(serverEntries, line[0] + ":" + line[1])
		}
	}

	// Get the server list
	serverInfoEntries := []serverData{}
	serverInfoChan := make(chan struct {serverData; error})
	for _, serverEntry := range serverEntries {
		go getServerInfo(serverEntry, serverInfoChan)
	}

	for i := 0; i < len(serverEntries); i++ {
		serverInfo := <-serverInfoChan
		if serverInfo.error != nil {
			if !os.IsTimeout(errors.Unwrap(errors.Unwrap(serverInfo.error))) {
				panic(serverInfo.error)
			}
		} else {
			serverInfoEntries = append(serverInfoEntries, serverInfo.serverData)
		}
	}

	// Caluclate column widths
	maxServerNameLength := 11
	maxPingLength := 3
	maxPlayersLength := 7
	maxIPLength := 10
	for _, serverInfo := range serverInfoEntries {
		if serverInfo.serverNameLength > maxServerNameLength {
			maxServerNameLength = serverInfo.serverNameLength
		}
		if len(fmt.Sprint(serverInfo.ping)) > maxPingLength {
			maxPingLength = len(fmt.Sprint(serverInfo.ping))
		}
		if len(fmt.Sprintf("%d/%d", serverInfo.players, serverInfo.maxPlayers)) > maxPlayersLength {
			maxPlayersLength = len(fmt.Sprint(serverInfo.players)) + 3
		}
		if len(serverInfo.ipAddress) > maxIPLength {
			maxIPLength = len(serverInfo.ipAddress)
		}
	}

	// Create the format string
	headerFormatString := fmt.Sprintf("%%-%ds\t%%-%ds\t%%-%ds\t%%-%ds", maxIPLength, maxServerNameLength, maxPlayersLength, maxPingLength)
	formatString := fmt.Sprintf("%%-%ds\t%%-s\t%% %ds\t%%%ddms", maxIPLength, maxPlayersLength, maxPingLength)

	// Print the header
	fmt.Printf(headerFormatString, "IP Address", "Server Name", "Players", "Ping")
	fmt.Println()
	for _, serverInfo := range serverInfoEntries {
		serverName := serverInfo.serverName + strings.Repeat(" ", maxServerNameLength - serverInfo.serverNameLength)
		playerCount := fmt.Sprintf("%d/%d", serverInfo.players, serverInfo.maxPlayers)
		fmt.Printf(formatString, serverInfo.ipAddress, serverName, playerCount, serverInfo.ping)
		fmt.Println()
	}
}

func getServerInfo(serverEntry string, dataOut chan<- struct {serverData; error}) {
	pingStart := time.Now()
	serverInfo, _, err := network.GetServerInfo(serverEntry)
	ping := time.Since(pingStart).Milliseconds()
	if err != nil {
		dataOut <- struct {serverData; error}{serverData{}, fmt.Errorf("failed to get serverinfo from goroutine: %w", err)}
		return
	}
	dataOut <- struct {serverData; error}{
		serverData{
			ipAddress: serverEntry,
	
			serverName:       coloriseServerName(serverInfo.ServerNameRaw),
			serverNameLength: len(serverInfo.ServerName),
	
			ping: int(ping),
	
			players:    int(serverInfo.NumberOfPlayer),
			maxPlayers: int(serverInfo.MaxPlayers),

		},
		nil,
	}
}
