package network

import (
	"fmt"
	"net"
	"time"

	"GoBun/srb2kart/network/info"
)


func GetServerInfo(adress string) (info.ServerInfoPacket, info.PlayerInfoPacket, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", adress)
	if err != nil {
		return info.ServerInfoPacket{}, info.PlayerInfoPacket{}, fmt.Errorf("failed to resolve udp address: %w", err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return info.ServerInfoPacket{}, info.PlayerInfoPacket{}, fmt.Errorf("failed to connect to server: %w", err)
	}
	defer conn.Close()

	var serverInfoPacket info.ServerInfoPacket
	var gotServerInfoPacket bool
	var playerInfoPacket info.PlayerInfoPacket
	var gotPlayerInfoPacket bool

	conn.SetReadBuffer(2048)
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	conn.Write(info.GetServerInfoPacket[:])
	for !gotServerInfoPacket || !gotPlayerInfoPacket {
		buffer := make([]byte, 2048)
		_, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			return info.ServerInfoPacket{}, info.PlayerInfoPacket{}, fmt.Errorf("error getting information from server: %w", err)
		}

		packet, err := info.ParsePacket(buffer)
		if err != nil {
			return info.ServerInfoPacket{}, info.PlayerInfoPacket{}, fmt.Errorf("error reading packet: %w", err)
		}

		switch packetType := packet.GetPacketType(); packetType {
		case info.PacketTypeServerInfo:
			serverInfoPacket = packet.(info.ServerInfoPacket)
			gotServerInfoPacket = true
		case info.PacketTypePlayerInfo:
			playerInfoPacket = packet.(info.PlayerInfoPacket)
			gotPlayerInfoPacket = true
		}
	}

	// Filter out all players from the player info slice that have a node of 255
	playersInGame := []info.PlayerInfoEntry{}
	for _, player := range playerInfoPacket.PlayerInfo {
		if player.Node != 255 {
			playersInGame = append(playersInGame, player)
		}
	}
	playerInfoPacket.PlayerInfo = playersInGame

	if len(playerInfoPacket.PlayerInfo) != int(serverInfoPacket.NumberOfPlayer) {
		return info.ServerInfoPacket{}, info.PlayerInfoPacket{}, fmt.Errorf("number of players in player info packet does not match number of players in server info packet: %w", err)
	}

	return serverInfoPacket, playerInfoPacket, nil
}
