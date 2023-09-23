package common

import "github.com/camdenorrb/minecraftPackets/primitive"

type SkyLightArray struct {
	Length primitive.VarInt
	Data   []byte
}

type BlockLightArray struct {
	Length primitive.VarInt
	Data   []byte
}

type PalettedContainerStructure struct {
	BitsPerEntry    uint8
	Palette         Palette
	DataArrayLength primitive.VarInt
	DataArray       []int64
}

type Palette struct {
	SingleValued *SingleValuedPalette
	Indirect     *IndirectPalette
	Direct       *DirectPalette
}

type SingleValuedPalette struct {
	Value primitive.VarInt
}

type IndirectPalette struct {
	PaletteLength primitive.VarInt
	Palette       []primitive.VarInt
}

type DirectPalette struct{}

type ChunkSection struct {
	BlockCount  int16
	BlockStates PalettedContainerStructure
	Biomes      PalettedContainerStructure
}
