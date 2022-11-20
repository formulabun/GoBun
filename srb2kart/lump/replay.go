package lump

import (
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

func ReadReplay(data io.Reader) (result ReplayHeaderRaw, err error) {
	err = binary.Read(data, binary.LittleEndian, &result)
	if err != nil {
		return result, err
	}
  headerText := string(result.DemoHeader[1:11])
	if demoHeader !=  headerText{
    return result, errors.New(fmt.Sprintf("Not a kart replay file: %s does not equal the demo header", headerText))
	}
  // TODO validate demo flags 
	return result, nil
}

func (r *ReplayHeaderRaw) GetTitle() string {
	return strings.SafeNullTerminated(r.Title[:])
}
