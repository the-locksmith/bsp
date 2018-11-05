package lumps

/**
Lump interface.
Organise Lump data in a cleaner and more accessible manner
*/
type ILump interface {
	/**
	Import a []byte to a defined lump structure(s).
	*/
	FromBytes([]byte, int32)

	/**
	Export lump structure back to []byte.
	*/
	ToBytes() []byte
}

type LumpGeneric struct {
	LumpInfo
	data []byte
}

func (lump *LumpGeneric) FromBytes(data []byte, length int32) {
	lump.length = length
	lump.data = data
}

func (lump *LumpGeneric) ToBytes() []byte {
	return lump.data
}

/**
Helper info for a lump
*/
type LumpInfo struct {
	length  int32
	version int32
}

/**
Return lump import length in bytes.
*/
func (info LumpInfo) GetLength() int32 {
	return info.length
}

/**
Set lump import length in bytes
*/
func (info LumpInfo) SetLength(length int32) {
	info.length = length
}

/**
Set bsp version of lump
*/
func (info LumpInfo) SetVersion(version int32) {
	info.version = version
}
