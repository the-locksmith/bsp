package texdata

import "github.com/galaco/bsp/primitives/common"

type TexData struct {
	Reflectivity common.Vector
	NameStringTableID int32
	Width int32
	Height int32
	ViewWidth int32
	ViewHeight int32
}