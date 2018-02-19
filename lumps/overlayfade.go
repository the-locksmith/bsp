package lumps

import (
	datatypes "github.com/galaco/bsp/lumps/datatypes/overlayfade"
	"unsafe"
	"encoding/binary"
	"bytes"
	"log"
)
/**
	Lump 60: Overlayfades
 */
type OverlayFade struct {
	LumpInfo
	data []datatypes.OverlayFade
}

func (lump OverlayFade) FromBytes(raw []byte, length int32) ILump {
	if length == 0 {
		return lump
	}
	lump.data = make([]datatypes.OverlayFade, length/int32(unsafe.Sizeof(datatypes.OverlayFade{})))
	err := binary.Read(bytes.NewBuffer(raw[:]), binary.LittleEndian, &lump.data)
	if err != nil {
		log.Fatal(err)
	}
	lump.LumpInfo.SetLength(length)

	return lump
}

func (lump OverlayFade) GetData() interface{} {
	return lump.data
}

func (lump OverlayFade) ToBytes() []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, lump.data)
	return buf.Bytes()
}
