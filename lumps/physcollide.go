package lumps

import (
	"bytes"
	"encoding/binary"
	primitives "github.com/galaco/bsp/primitives/physcollide"
	"unsafe"
)

// PhysCollide is Lump 20: PhysCollide
type PhysCollide struct {
	Generic
	data []primitives.PhysCollideEntry
}

// Unmarshall Imports this lump from raw byte data
func (lump *PhysCollide) Unmarshall(raw []byte) (err error) {
	length := len(raw)
	lump.Metadata.SetLength(length)
	if length == 0 {
		return
	}

	lump.data = make([]primitives.PhysCollideEntry, length/int(unsafe.Sizeof(primitives.PhysCollideEntry{})))
	err = binary.Read(bytes.NewBuffer(raw), binary.LittleEndian, &lump.data)

	return err
}

// GetData gets internal format structure data
func (lump *PhysCollide) GetData() []primitives.PhysCollideEntry {
	return lump.data
}

// Marshall dumps this lump back to raw byte data
func (lump *PhysCollide) Marshall() ([]byte, error) {
	var buf bytes.Buffer
	for _, entry := range lump.data {
		err := binary.Write(&buf, binary.LittleEndian, entry.ModelHeader)
		if err != nil {
			return nil, err
		}
		for _, solid := range entry.Solids {
			if err = binary.Write(&buf, binary.LittleEndian, solid.Size); err != nil {
				return nil, err
			}
			if err = binary.Write(&buf, binary.LittleEndian, solid.CollisionBinary); err != nil {
				return nil, err
			}
		}
		if err = binary.Write(&buf, binary.LittleEndian, []byte(entry.TextBuffer)); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}
