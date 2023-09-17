package primitive

import "fmt"

type BitSet struct {
	Data []uint64 // Can't inline type, since we need to change this value in the methods
}

func NewBitSet(size uint) BitSet {
	return BitSet{
		Data: make([]uint64, (size+63)/64),
	}
}

func (b *BitSet) Get(index uint) bool {
	return b.Data[index/64]&(1<<(index%64)) != 0
}

func (b *BitSet) Set(index uint) {
	b.expandFor(index)
	b.Data[index/64] |= 1 << (index % 64)
}

func (b *BitSet) Flip(index uint) {
	b.expandFor(index)
	b.Data[index/64] ^= 1 << (index % 64)
}

func (b *BitSet) Length() uint {
	return uint(len(b.Data) * 64)
}

func (b *BitSet) expandFor(index uint) {
	if index >= uint(len(b.Data)*64) {
		difference := (int(index)-len(b.Data))/64 + 1
		b.Data = append(b.Data, make([]uint64, difference)...)
	}
}

func (b *BitSet) String() string {
	var str string
	for _, v := range b.Data {
		str += fmt.Sprintf("%064b", v)
	}
	return str
}
