package lumps

import (
	datatypes "github.com/galaco/bsp/lumps/datatypes/texdata"
	"encoding/binary"
	"bytes"
	"log"
	"unsafe"
)

/**
	Lump 2: TexData
 */
type TexData struct {
	LumpInfo
	data []datatypes.TexData
}
func (lump TexData) FromBytes(raw []byte, length int32) ILump {
	lump.data = make([]datatypes.TexData, length/int32(unsafe.Sizeof(datatypes.TexData{})))
	err := binary.Read(bytes.NewBuffer(raw[:]), binary.LittleEndian, &lump.data)
	if err != nil {
		log.Fatal(err)
	}
	lump.LumpInfo.SetLength(length)

	return lump
}

func (lump TexData) GetData() interface{} {
	return lump.data
}

func (lump TexData) ToBytes() []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, lump.data)
	return buf.Bytes()
}
