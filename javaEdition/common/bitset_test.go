package common

import (
	"math"
	"testing"
)

func FuzzBitSet_Set(f *testing.F) {

	f.Add(uint(0))
	f.Add(uint(math.MaxInt32))
	f.Add(uint(math.MaxUint32))

	f.Fuzz(func(t *testing.T, index uint) {

		bitSet := NewBitSet(0)

		if index > math.MaxUint32 {
			t.Skip()
		}

		bitSet.Set(index)

		if !bitSet.Get(index) {
			t.Errorf("Expected bit at index %d to be set", index)
		}
		if bitSet.Length() < index {
			t.Errorf("Expected bitset length to be at least %d, got %d", index, bitSet.Length())
		}
		if bitSet.Length()%64 != 0 {
			t.Errorf("Expected bitset length to be a multiple of 64, got %d", bitSet.Length())
		}
	})
}

func FuzzBitSet_Flip(f *testing.F) {

	f.Add(uint(0))
	f.Add(uint(math.MaxInt32))
	f.Add(uint(math.MaxUint32))

	f.Fuzz(func(t *testing.T, index uint) {

		bitSet := NewBitSet(0)

		if index < 0 || index > math.MaxUint32 {
			t.Skip()
		}

		bitSet.Flip(index)

		if !bitSet.Get(index) {
			t.Errorf("Expected bit at index %d to be set %v", index, bitSet.data)
		}
		if bitSet.Length() < index {
			t.Errorf("Expected bitset length to be at least %d, got %d", index, bitSet.Length())
		}
		if bitSet.Length()%64 != 0 {
			t.Errorf("Expected bitset length to be a multiple of 64, got %d", bitSet.Length())
		}
	})
}
