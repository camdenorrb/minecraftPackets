package common

type Angle float32

type ChunkPosition struct {
	X int32
	Z int32
}

type BlockPosition struct {
	Encoded int64
}

func NewBlockPosition(x, y, z int32) *BlockPosition {
	return &BlockPosition{
		Encoded: (int64(x&0x3FFFFFF) << 38) | (int64(z&0x3FFFFFF) << 12) | int64(y&0xFFF),
	}
}

func (p *BlockPosition) SetPosition(x, y, z int32) {
	p.Encoded = (int64(x&0x3FFFFFF) << 38) | (int64(z&0x3FFFFFF) << 12) | int64(y&0xFFF)
}

func (p *BlockPosition) GetX() int32 {
	return int32(p.Encoded >> 38)
}

func (p *BlockPosition) GetY() int32 {
	return int32(p.Encoded & 0xFFF)
}

func (p *BlockPosition) GetZ() int32 {
	return int32(p.Encoded << 26 >> 38)
}

func (p *BlockPosition) GetChunkPosition() *ChunkPosition {
	return &ChunkPosition{
		X: int32(p.Encoded >> 38),
		Z: int32(p.Encoded << 26 >> 38),
	}
}
