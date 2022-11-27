package lump

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"time"

	"GoBun/functional/array"
	"GoBun/functional/strings"
)

var demoHeader = "KartReplay"

type ReplayRaw struct {
	HeaderPreFileEntries
	WadEntries
	HeaderPostFileEntries
	CVarEntries
	PlayerEntries
	PlayerListingEnd byte
}

type HeaderPreFileEntries struct {
	DemoHeader  [12]byte
	Version     uint8
	SubVersion  uint8
	DemoVersion uint16

	Title    [64]byte
	Checksum [16]byte

	Play    [4]byte
	GameMap uint16
	MapMd5  [16]byte

	DemoFlags uint8
	GameType  uint8

	FileCount byte
}

type WadEntries []WadEntry

type WadEntry struct {
	FileName string
	WadMd5   [16]byte
}

type HeaderPostFileEntries struct {
	Time uint32
	Lap  uint32

	Seed     uint32
	Reserved uint32

	CVarCount int16
}

type CVarEntries []CVarEntry

type CVarEntry struct {
	CVarId uint16
	Value  string
	False  uint8
}

type PlayerEntries []PlayerEntry

type PlayerEntry struct {
	Spec uint8
	PlayerEntryData
}

type PlayerEntryData struct {
	Name   [16]byte
	Skin   [16]byte
	Color  [16]byte
	Score  uint32
	Speed  byte
	Weight byte
}

func ReadReplayData(data []byte) (result ReplayRaw, err error) {
	dataReader := bytes.NewReader(data)
	return ReadReplay(dataReader)
}

func ReadReplay(data io.Reader) (result ReplayRaw, err error) {
	var headerPreReplays HeaderPreFileEntries
	err = binary.Read(data, binary.LittleEndian, &headerPreReplays)
	if err != nil {
		return result, fmt.Errorf("Could not read the replay header before addons: %s", err)
	}
	result.HeaderPreFileEntries = headerPreReplays

	fileCount := int(result.FileCount)
	result.WadEntries = make([]WadEntry, fileCount)
	readCount := 0
	for readCount < fileCount {
		entry, err := readWadEntry(data)
		if err != nil {
			return result, fmt.Errorf("Could not read the file entry number %d: %s", readCount+1, err)
		}
		result.WadEntries[readCount] = entry
		readCount++
	}

	var headerPostReplays HeaderPostFileEntries
	err = binary.Read(data, binary.LittleEndian, &headerPostReplays)
	if err != nil {
		return result, fmt.Errorf("Could not read the replay header before addons: %s", err)
	}
	result.HeaderPostFileEntries = headerPostReplays

	cvarCount := int(result.CVarCount)
	readCount = 0
	for readCount < cvarCount {
		result.CVarEntries = make([]CVarEntry, cvarCount)
		entry, err := readCVarEntry(data)
		if err != nil {
			return result, fmt.Errorf("Could not read cvar entry number %d: %s", readCount+1, err)
		}
		result.CVarEntries[readCount] = entry
		readCount++
	}

	players, end, err := readPlayerEntries(data)
	if err != nil {
		return result, fmt.Errorf("Could not read player entries from file: %s", err)
	}
	result.PlayerEntries = players
	result.PlayerListingEnd = end

	return result, validate(result)
}

func (R *ReplayRaw) Write(writer io.Writer) error {
	err := binary.Write(writer, binary.LittleEndian, R.HeaderPreFileEntries)
	if err != nil {
		return fmt.Errorf("Could not write the replay header: %s", err)
	}
	for _, replayEntry := range R.WadEntries {
		_, err = io.WriteString(writer, replayEntry.FileName)
		if err != nil {
			return fmt.Errorf("Could not write replay file name: %s", err)
		}
		_, err = writer.Write(replayEntry.WadMd5[:])
		if err != nil {
			return fmt.Errorf("Could not write replay file checksum: %s", err)
		}
	}

	err = binary.Write(writer, binary.LittleEndian, R.HeaderPostFileEntries)
	if err != nil {
		return err
	}

	for _, cvarEntry := range R.CVarEntries {
		err = binary.Write(writer, binary.LittleEndian, cvarEntry.CVarId)
		_, err = io.WriteString(writer, cvarEntry.Value)
		err = binary.Write(writer, binary.LittleEndian, cvarEntry.False)
		if err != nil {
			return err
		}
	}

	for _, playerEntry := range R.PlayerEntries {
		err = binary.Write(writer, binary.LittleEndian, playerEntry)
		if err != nil {
			return err
		}
	}

	err = binary.Write(writer, binary.LittleEndian, R.PlayerListingEnd)
	return err
}

func readWadEntry(data io.Reader) (result WadEntry, err error) {
	filename, extra, err := readNullTerminatedString(data, 16)
	if err != nil {
		return result, err
	}

	result.FileName = filename
	copy(result.WadMd5[:len(extra)], extra)

	n, err := data.Read(result.WadMd5[len(extra):])
	if err != nil {
		return result, fmt.Errorf("Could not read a file entry from the replay: ", err)
	}
	if n < 16-len(extra) {
		return result, fmt.Errorf("Unexpected end to the replay file.")
	}
	return result, nil
}

func readCVarEntry(data io.Reader) (result CVarEntry, err error) {
	err = binary.Read(data, binary.LittleEndian, &result.CVarId)
	if err != nil {
		return result, fmt.Errorf("Could not read CVar ID: %s", err)
	}

	cvarValue, extra, err := readNullTerminatedString(data, 2)
	if err != nil {
		return result, err
	}

	result.Value = cvarValue
	if len(extra) != 2 {
		binary.Read(data, binary.LittleEndian, &result.False)
	} else {
		result.False = extra[1]
	}
	return result, err
}

func readPlayerEntries(data io.Reader) (result PlayerEntries, end byte, err error) {
	var spec byte

	err = binary.Read(data, binary.LittleEndian, &spec)
	if err != nil {
		return result, spec, fmt.Errorf("Could not read player spec value: %s", err)
	}

	for spec != 0xff {
		var playerEntry PlayerEntry

		err = binary.Read(data, binary.LittleEndian, &playerEntry.PlayerEntryData)
		if err != nil {
			return result, spec, fmt.Errorf("Could not read player entry: %s", err)
		}
		playerEntry.Spec = spec
		result = append(result, playerEntry)

		err = binary.Read(data, binary.LittleEndian, &spec)
		if err != nil {
			return result, spec, fmt.Errorf("Could not read player spec value: %s", err)
		}
	}
	return result, spec, nil
}

func validate(replay ReplayRaw) error {
	headerText := string(replay.DemoHeader[1:11])
	badFileError := errors.New("Not a kart replay file")

	if demoHeader != headerText && replay.DemoHeader[0] == 0xf0 && replay.DemoHeader[11] == 0x0f {
		return badFileError
	}
	if string(replay.Play[:]) != "PLAY" {
		return badFileError
	}
	if replay.DemoFlags&0x2 == 0 {
		return badFileError
	}
	return nil
}

func readNullTerminatedString(data io.Reader, bufferSize int) (string, []byte, error) {
	var result bytes.Buffer
	buffer := make([]byte, bufferSize)

	for {
		n, err := data.Read(buffer)
		if err != nil {
			return result.String(), buffer, fmt.Errorf("Could not read from the replay: ", err)
		}
		if n < bufferSize {
			return result.String(), buffer, fmt.Errorf("Unexpected end to the replay file.")
		}

		found := array.FindFirstIndexMatching(buffer, 0x00)
		if found >= 0 {
			result.Write(buffer[:found+1])
			return result.String(), buffer[found+1:], nil
		}
		result.Write(buffer)
	}
}

func (r *HeaderPreFileEntries) GetTitle() string {
	return strings.SafeNullTerminated(r.Title[:])
}

func (r *HeaderPostFileEntries) GetTime() time.Duration {
	return time.Millisecond * time.Duration(1000*r.Time/35)
}

func (r *HeaderPostFileEntries) GetBestLapTime() time.Duration {
	return time.Millisecond * time.Duration(1000*r.Lap/35)
}
