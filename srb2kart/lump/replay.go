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

const headerLength = 12 + 8 + 8 + 16 + 64 + 16

var demoHeader = "KartReplay"

type ReplayRaw struct {
	HeaderPreFileEntries
	WadEntries
	HeaderPostFileEntries
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
		writer.Write([]byte{0x0})
		_, err = writer.Write(replayEntry.WadMd5[:])
		if err != nil {
			return fmt.Errorf("Could not write replay file checksum: %s", err)
		}
	}
	return nil
}

func readWadEntry(data io.Reader) (result WadEntry, err error) {
	var filename bytes.Buffer
	buffer := make([]byte, 16)

	for {
		n, err := data.Read(buffer)
		if err != nil {
			return result, fmt.Errorf("Could not read a file entry from the replay: ", err)
		}
		if n < 16 {
			return result, fmt.Errorf("Unexpected end to the replay file.")
		}

		found := array.FindFirstIndexMatching(buffer, 0x00)
		if found >= 0 {
			filename.Write(buffer[:found+1])
			result.FileName = filename.String()
			copy(result.WadMd5[:len(buffer)-found-1], buffer[found+1:])
			n, err = data.Read(result.WadMd5[:found+1]) // TODO don't ignore errors
			if err != nil {
				return result, fmt.Errorf("Could not read a file entry from the replay: ", err)
			}
			if n < found {
				return result, fmt.Errorf("Unexpected end to the replay file.")
			}
			return result, nil
		}
		filename.Write(buffer)
	}
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

func (r *HeaderPreFileEntries) GetTitle() string {
	return strings.SafeNullTerminated(r.Title[:])
}

func (r *HeaderPostFileEntries) GetTime() time.Duration {
	return time.Millisecond * time.Duration(1000 * r.Time / 35)
}

func (r *HeaderPostFileEntries) GetBestLapTime() time.Duration {
	return time.Millisecond * time.Duration(1000 * r.Lap / 35)
}
