package common

type Angle float32

type ChunkPosition struct {
	X int32
	Z int32
}

type BlockPosition uint64

func NewBlockPosition(x, y, z int32) BlockPosition {
	return BlockPosition((uint64(x&0x3FFFFFF) << 38) | (uint64(z&0x3FFFFFF) << 12) | uint64(y&0xFFF))
}

func (p BlockPosition) X() int32 {

	x := int32(p >> 38)
	if x >= 1<<25 {
		x -= 1 << 26
	}

	return x
}

func (p BlockPosition) Y() int32 {
	y := int32(p & 0xFFF)
	if y >= 1<<11 {
		y -= 1 << 12
	}
	return y
}

func (p BlockPosition) Z() int32 {

	z := int32((p >> 12) & 0x3FFFFFF)
	if z >= 1<<25 {
		z -= 1 << 26
	}

	return z
}

func (p BlockPosition) GetChunkPosition() *ChunkPosition {
	return &ChunkPosition{
		X: p.X() >> 4,
		Z: p.Z() >> 4,
	}
}
