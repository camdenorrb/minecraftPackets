package types

import "fmt"

type BitSet struct {
	data []uint64
}

func NewBitSet(size uint) BitSet {
	return BitSet{
		data: make([]uint64, (size+63)/64, (size+63)/64),
	}
}

func (b *BitSet) Get(index uint) bool {
	return b.data[index/64]&(1<<(index%64)) != 0
}

func (b *BitSet) Set(index uint) {
	b.expandFor(index)
	b.data[index/64] |= 1 << (index % 64)
}

func (b *BitSet) Flip(index uint) {
	b.expandFor(index)
	b.data[index/64] ^= 1 << (index % 64)
}

func (b *BitSet) Length() uint {
	return uint(len(b.data) * 64)
}

func (b *BitSet) expandFor(index uint) {
	if index >= uint(len(b.data)*64) {
		difference := (int(index)-len(b.data))/64 + 1
		b.data = append(b.data, make([]uint64, difference, difference)...)
	}
}

func (b *BitSet) String() string {
	var str string
	for _, v := range b.data {
		str += fmt.Sprintf("%064b", v)
	}
	return str
}
