package serverinfo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

var getServerInfoPacket = [...]byte{
	0x58, 0x46, 0x23, 0x01,
	0x00,
	0x00,
	0x0C,
	0x00,

	0x01,
	0x1f, 0x02, 0x00, 0x00,
}

type packetType uint8

const (
	packetTypeServerInfo packetType = 0x0D
	packetTypePlayerInfo packetType = 0x0E
)

type KartPacket interface{
	GetPacketType() packetType
}

type KartPacketHeader struct {
	Checksum   uint32
	Ack        uint8
	Ackret     uint8
	PacketType packetType
	_          [1]byte
}

type kartServerInfoPacketRaw struct {
	Header         KartPacketHeader
	_              uint8 // 255 field
	PacketVersion  uint8
	Application    [16]byte
	Version        uint8
	SubVersion     uint8
	NumberOfPlayer uint8
	MaxPlayers     uint8
	GameType       uint8
	ModifiedGame   bool
	CheatsEnabled  bool
	KartVars       uint8
	FileNeededNum  uint8
	Time           uint32
	LevelTime      uint32
	ServerName     [32]byte
	MapName        [8]byte
	MapTitle       [33]byte
	MapMD5         [16]uint8
	ActNum         uint8
	IsZone         uint8
	HttpSource     [256]byte
	FileNeeded     [915]uint8
}

type KartServerInfoPacket struct {
	Header         KartPacketHeader
	PacketVersion  uint8
	Application    [16]byte
	Version        uint8
	SubVersion     uint8
	NumberOfPlayer uint8
	MaxPlayers     uint8
	GameType       uint8
	ModifiedGame   bool
	CheatsEnabled  bool
	KartVars       uint8
	FileNeededNum  uint8
	Time           uint32
	LevelTime      uint32
	ServerName     string
	MapName        string
	MapTitle       string
	MapMD5         [16]uint8
	ActNum         uint8
	IsZone         uint8
	HttpSource     string
	FileNeeded     []FileNeededEntry
}

func (p *KartServerInfoPacket) GetPacketType() packetType {
	return p.Header.PacketType
}

type FileNeededEntry struct {
	WillSend bool
	TotalSize uint32
	FileName string
	MD5 [16]uint8
}

type kartPlayerInfoEntryRaw struct {
	Node         uint8
	Name         [21 + 1]byte
	Address      [4]uint8
	Team         uint8
	Skin         uint8
	Data         uint8
	Score        uint32
	TimeInServer uint16
}

type kartPlayerInfoPacketRaw struct {
	Header     KartPacketHeader
	PlayerInfo [32]kartPlayerInfoEntryRaw
}

type KartPlayerInfoEntry struct {
	Node         uint8
	Name         string
	Address      [4]uint8
	Team         uint8
	Skin         uint8
	Data         uint8
	Score        uint32
	TimeInServer uint16
}

type KartPlayerInfoPacket struct {
	Header     KartPacketHeader
	PlayerInfo []KartPlayerInfoEntry
}

func (p *KartPlayerInfoPacket) GetPacketType() packetType {
	return p.Header.PacketType
}

func nullTerminated(data []byte) string {
	newBytes := make([]byte, 0)

	// Filter out all bytes above 127
	for i := 0; i < len(data); i++ {
		if data[i] <= 127 {
			newBytes = append(newBytes, data[i])
		}
	}
	before, _, _ := bytes.Cut(newBytes, []byte{0})
	return string(before)
}

func readPacket(data []byte) (KartPacket, error) {
	header := KartPacketHeader{}
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &header); err != nil {
		return nil, err
	}

	switch header.PacketType {
	case packetTypeServerInfo:
		return unpackServerInfoPacket(data)
	case packetTypePlayerInfo:
		return unpackPlayerInfoPacket(data)
	default:
		return nil, fmt.Errorf("unknown packet type: %d", header.PacketType)
	}
}

func parseFileNeeded(data [915]byte, fileNeededCount int) []FileNeededEntry {
	var entries []FileNeededEntry
	buf := bytes.NewBuffer(data[:])
	for i := 0; i < fileNeededCount; i++ {
		entry := FileNeededEntry{}
		binary.Read(buf, binary.LittleEndian, &entry.WillSend)
		b, _ := buf.ReadByte()
		entry.WillSend = (b >> 4) == 1
		binary.Read(buf, binary.LittleEndian, &entry.TotalSize)
		entry.FileName, _ = buf.ReadString(0)
		// Trim off the null terminator
		entry.FileName = entry.FileName[:len(entry.FileName)-1]
		binary.Read(buf, binary.LittleEndian, &entry.MD5)
		entries = append(entries, entry)
	}
	return entries
}

func unpackServerInfoPacket(data []byte) (*KartServerInfoPacket, error) {
	var packetRaw kartServerInfoPacketRaw
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &packetRaw); err != nil {
		return nil, err
	}
	packet := KartServerInfoPacket{
		Header:         packetRaw.Header,
		PacketVersion:  packetRaw.PacketVersion,
		Application:    packetRaw.Application,
		Version:        packetRaw.Version,
		SubVersion:     packetRaw.SubVersion,
		NumberOfPlayer: packetRaw.NumberOfPlayer,
		MaxPlayers:     packetRaw.MaxPlayers,
		GameType:       packetRaw.GameType,
		ModifiedGame:   packetRaw.ModifiedGame,
		CheatsEnabled:  packetRaw.CheatsEnabled,
		KartVars:       packetRaw.KartVars,
		FileNeededNum:  packetRaw.FileNeededNum,
		Time:           packetRaw.Time,
		LevelTime:      packetRaw.LevelTime,
		ServerName:     nullTerminated(packetRaw.ServerName[:]),
		MapName:        nullTerminated(packetRaw.MapName[:]),
		MapTitle:       nullTerminated(packetRaw.MapTitle[:]),
		MapMD5:         packetRaw.MapMD5,
		ActNum:         packetRaw.ActNum,
		IsZone:         packetRaw.IsZone,
		HttpSource:     nullTerminated(packetRaw.HttpSource[:]),
		FileNeeded:     parseFileNeeded(packetRaw.FileNeeded, int(packetRaw.FileNeededNum)),
	}
	return &packet, nil
}

func unpackPlayerInfoPacket(data []byte) (*KartPlayerInfoPacket, error) {
	var packetRaw kartPlayerInfoPacketRaw
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &packetRaw); err != nil {
		return nil, err
	}
	packet := KartPlayerInfoPacket{
		Header: packetRaw.Header,
	}
	for i := 0; i < len(packetRaw.PlayerInfo); i++ {
		entry := packetRaw.PlayerInfo[i]
		packet.PlayerInfo = append(packet.PlayerInfo, KartPlayerInfoEntry{
			Node: entry.Node,
			Name: nullTerminated(entry.Name[:]),
			Address: [4]uint8{
				entry.Address[0],
				entry.Address[1],
				entry.Address[2],
				entry.Address[3],
			},
			Team:         entry.Team,
			Skin:         entry.Skin,
			Data:         entry.Data,
			Score:        entry.Score,
			TimeInServer: entry.TimeInServer,
		})
	}
	return &packet, nil
}

func GetSRB2Info(adress string) (*KartServerInfoPacket, *KartPlayerInfoPacket, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", adress)
	if err != nil {
		return nil, nil, err
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()

	var serverInfoPacket *KartServerInfoPacket
	var playerInfoPacket *KartPlayerInfoPacket

	conn.SetReadBuffer(2048)
	conn.Write(getServerInfoPacket[:])
	for serverInfoPacket == nil || playerInfoPacket == nil {
		buffer := make([]byte, 2048)
		_, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			return nil, nil, fmt.Errorf("error getting information from udp: %w", err)
		}

		packet, err := readPacket(buffer)
		if err != nil {
			return nil, nil, fmt.Errorf("error reading packet: %w", err)
		}

		switch packetType := packet.GetPacketType(); packetType {
		case packetTypeServerInfo:
			serverInfoPacket = packet.(*KartServerInfoPacket)
		case packetTypePlayerInfo:
			playerInfoPacket = packet.(*KartPlayerInfoPacket)
		}
	}

	// Filter out all players from the player info slice that have a node of 255
	playersInGame := []KartPlayerInfoEntry{}
	for _, player := range playerInfoPacket.PlayerInfo {
		if player.Node != 255 {
			playersInGame = append(playersInGame, player)
		}
	}
	playerInfoPacket.PlayerInfo = playersInGame

	if len(playerInfoPacket.PlayerInfo) != int(serverInfoPacket.NumberOfPlayer) {
		return nil, nil, fmt.Errorf("number of players in player info packet does not match number of players in server info packet: %w", err)
	}

	return serverInfoPacket, playerInfoPacket, nil
}
