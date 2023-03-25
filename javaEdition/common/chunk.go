package common

type SkyLightArray struct {
	Length VarInt
	Data   []byte
}

type BlockLightArray struct {
	Length VarInt
	Data   []byte
}

type PalettedContainerStructure struct {
	BitsPerEntry    uint8
	Palette         Palette
	DataArrayLength VarInt
	DataArray       []int64
}

type Palette struct {
	*SingleValuedPalette
	*IndirectPalette
	*DirectPalette
}

type SingleValuedPalette struct {
	Value VarInt
}

type IndirectPalette struct {
	PaletteLength VarInt
	Palette       []VarInt
}

type DirectPalette struct{}

type ChunkSection struct {
	BlockCount  int16
	BlockStates PalettedContainerStructure
	Biomes      PalettedContainerStructure
}
