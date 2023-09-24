package common

import "github.com/camdenorrb/minecraftPackets/primitive"

type SkyLightArray struct {
	Data []byte
}

type BlockLightArray struct {
	Data []byte
}

type PalettedContainerStructure struct {
	BitsPerEntry uint8
	Palette      Palette
	DataArray    []int64
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
	Palette []primitive.VarInt
}

type DirectPalette struct{}

type ChunkSection struct {
	BlockCount  uint16
	BlockStates PalettedContainerStructure
	Biomes      PalettedContainerStructure
}
