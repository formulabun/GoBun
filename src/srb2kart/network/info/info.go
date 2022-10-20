package info

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"GoBun/functional/strings"
)

var GetServerInfoPacket = [...]byte{
	0x58, 0x46, 0x23, 0x01,
	0x00,
	0x00,
	0x0C,
	0x00,

	0x01,
	0x1f, 0x02, 0x00, 0x00,
}

type PacketType uint8

const (
	PacketTypeServerInfo PacketType = 0x0D
	PacketTypePlayerInfo PacketType = 0x0E
)

type packetHeader struct {
	Checksum   uint32
	Ack        uint8
	Ackret     uint8
	PacketType PacketType
	_          [1]byte
}

type serverInfoPacketRaw struct {
	Header         packetHeader
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

type ServerInfoPacket struct {
	header         packetHeader
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
	ServerNameRaw  []byte
	ServerName     string
	MapName        string
	MapTitle       string
	MapMD5         [16]uint8
	ActNum         uint8
	IsZone         uint8
	HttpSource     string
	FileNeeded     []FileNeededEntry
}

func (p ServerInfoPacket) GetPacketType() PacketType {
	return p.header.PacketType
}

type playerInfoPacketRaw struct {
	Header     packetHeader
	PlayerInfo [32]playerInfoEntryRaw
}

type playerInfoEntryRaw struct {
	Node         uint8
	Name         [21 + 1]byte
	Address      [4]uint8
	Team         uint8
	Skin         uint8
	Data         uint8
	Score        uint32
	TimeInServer uint16
}

type PlayerInfoPacket struct {
	header     packetHeader
	PlayerInfo []PlayerInfoEntry
}

func (p PlayerInfoPacket) GetPacketType() PacketType {
	return p.header.PacketType
}

type PlayerInfoEntry struct {
	Node         uint8
	Name         string
	Address      [4]uint8
	Team         uint8
	Skin         uint8
	Data         uint8
	Score        uint32
	TimeInServer uint16
}

type packet interface {
	GetPacketType() PacketType
}

type FileNeededEntry struct {
	WillSend  bool
	TotalSize uint32
	FileName  string
	MD5       [16]uint8
}

func ParsePacket(data []byte) (packet, error) {
	header := packetHeader{}
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &header); err != nil {
		return nil, err
	}

	switch header.PacketType {
	case PacketTypeServerInfo:
		return parseServerInfoPacket(data)
	case PacketTypePlayerInfo:
		return parsePlayerInfoPacket(data)
	default:
		return nil, fmt.Errorf("unknown packet type: %d", header.PacketType)
	}
}

func parseFileNeeded(data []byte, fileNeededCount int) []FileNeededEntry {
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

func parseServerInfoPacket(data []byte) (ServerInfoPacket, error) {
	var packetRaw serverInfoPacketRaw
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &packetRaw); err != nil {
		return ServerInfoPacket{}, fmt.Errorf("failed to unpack server info packet: %v", err)
	}
	packet := ServerInfoPacket{
		header:         packetRaw.Header,
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
		ServerNameRaw: strings.NullTerminated(packetRaw.ServerName[:]),
		ServerName:     strings.SafeNullTerminated(packetRaw.ServerName[:]),
		MapName:        strings.SafeNullTerminated(packetRaw.MapName[:]),
		MapTitle:       strings.SafeNullTerminated(packetRaw.MapTitle[:]),
		MapMD5:         packetRaw.MapMD5,
		ActNum:         packetRaw.ActNum,
		IsZone:         packetRaw.IsZone,
		HttpSource:     strings.SafeNullTerminated(packetRaw.HttpSource[:]),
		FileNeeded:     parseFileNeeded(packetRaw.FileNeeded[:], int(packetRaw.FileNeededNum)),
	}
	return packet, nil
}

func parsePlayerInfoPacket(data []byte) (PlayerInfoPacket, error) {
	var packetRaw playerInfoPacketRaw
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &packetRaw); err != nil {
		return PlayerInfoPacket{}, fmt.Errorf("failed to unpack playerinfo packet: %w", err)
	}
	packet := PlayerInfoPacket{
		header: packetRaw.Header,
	}
	for i := 0; i < len(packetRaw.PlayerInfo); i++ {
		// Skin nodes if they are not used
		if packetRaw.PlayerInfo[i].Node == 255 {
			continue
		}
		entry := packetRaw.PlayerInfo[i]
		packet.PlayerInfo = append(packet.PlayerInfo, PlayerInfoEntry{
			Node: entry.Node,
			Name: strings.SafeNullTerminated(entry.Name[:]),
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
	return packet, nil
}
