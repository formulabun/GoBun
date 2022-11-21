package lump

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"GoBun/functional/array"
	"GoBun/functional/strings"
)

const headerLength = 12 + 8 + 8 + 16 + 64 + 16

var demoHeader = "KartReplay"

type ReplayRaw struct {
	ReplayHeaderRaw
	WadEntries
}

type ReplayHeaderRaw struct {
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

func ReadReplayData(data []byte) (result ReplayRaw, err error) {
	dataReader := bytes.NewReader(data)
	return ReadReplay(dataReader)
}

func ReadReplay(data io.Reader) (result ReplayRaw, err error) {
	var headerPreReplays ReplayHeaderRaw
	err = binary.Read(data, binary.LittleEndian, &headerPreReplays)
	if err != nil {
		return result, fmt.Errorf("Could not read the replay header before addons: %s", err)
	}
	result.ReplayHeaderRaw = headerPreReplays
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
	return result, validate(result)
}

func (R *ReplayRaw) Write(writer io.Writer) error {
  err := binary.Write(writer, binary.LittleEndian, R.ReplayHeaderRaw)
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
	var fileName bytes.Buffer

	buffer := make([]byte, 17) // length of checksum + null character
	for n, err := data.Read(buffer); err == nil; {
		if err != nil {
			return result, err
		}
		if n < len(buffer) {
			return result, errors.New("Unexpected end of replay file")
		}

		i := array.FindFirstIndex(buffer, func(b byte) bool {
			return b == 0x0
		})
		if i < 0 || i >= len(buffer) {
			fileName.Write(buffer)
		} else {
			fileName.Write(buffer[:i])
			result.FileName = fileName.String()
			copy(result.WadMd5[:], buffer[i+1:])
			md5 := make([]byte, i)
			data.Read(md5)
			copy(result.WadMd5[len(buffer)-i-1:], md5)
			return result, nil
		}
	}
	return result, err
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

func (r *ReplayHeaderRaw) GetTitle() string {
	return strings.SafeNullTerminated(r.Title[:])
}
