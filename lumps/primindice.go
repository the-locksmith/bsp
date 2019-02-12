package lumps

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Lump 39: PrimIndice
type PrimIndice struct {
	LumpGeneric
	data []uint16
}

// Import this lump from raw byte data
func (lump *PrimIndice) FromBytes(raw []byte, length int32) {
	lump.LumpInfo.SetLength(length)
	if length == 0 {
		return
	}

	lump.data = make([]uint16, length/int32(2))
	err := binary.Read(bytes.NewBuffer(raw), binary.LittleEndian, &lump.data)
	if err != nil {
		log.Fatal(err)
	}
}

// Get internal format structure data
func (lump *PrimIndice) GetData() []uint16 {
	return lump.data
}

// Dump this lump back to raw byte data
func (lump *PrimIndice) ToBytes() ([]byte,error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, lump.data)
	return buf.Bytes(),err
}
