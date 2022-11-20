package lump

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"GoBun/functional/strings"
)

const headerLength = 12 + 8 + 8 + 16 + 64 + 16

var demoHeader = "KartReplay"

type ReplayHeaderRaw struct {
	DemoHeader  [12]byte
	Version     uint8
	SubVersion  uint8
	DemoVersion uint16

	Title    [64]byte
	Checksum [16]byte
}

func ReadReplayData(data []byte) (result ReplayHeaderRaw, err error) {
	dataReader := bytes.NewReader(data)
	return ReadReplay(dataReader)
}

func ReadReplay(data io.Reader) (result ReplayHeaderRaw, err error) {
	err = binary.Read(data, binary.LittleEndian, &result)
	if err != nil {
		return result, err
	}
	return result, validate(result)
}

func validate(replay ReplayHeaderRaw) error {
	headerText := string(replay.DemoHeader[1:11])
	if demoHeader != headerText {
		return errors.New(fmt.Sprintf("Not a kart replay file: %s does not equal the demo header", headerText))
	}
	// TODO validate demo flags
	return nil
}

func (r *ReplayHeaderRaw) GetTitle() string {
	return strings.SafeNullTerminated(r.Title[:])
}
