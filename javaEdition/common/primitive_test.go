package common

import (
	"math"
	"testing"
)

func FuzzVarInt(f *testing.F) {
	f.Add(0)
	f.Add(math.MinInt32)
	f.Add(math.MaxInt32)
	f.Fuzz(func(t *testing.T, i int) {
		encoded := VarInt(i).Encode()
		decoded, err := DecodeVarInt(encoded)
		if err != nil {
			t.Error(err)
		}
		if *decoded != int32(i) {
			t.Errorf("Expected %d, got %d", i, *decoded)
		}
	})
}

func FuzzVarLong(f *testing.F) {
	f.Add(0)
	f.Add(math.MinInt64)
	f.Add(math.MaxInt64)
	f.Fuzz(func(t *testing.T, i int) {
		encoded := VarLong(i).Encode()
		decoded, err := DecodeVarLong(encoded)
		if err != nil {
			t.Error(err)
		}
		if *decoded != int64(i) {
			t.Errorf("Expected %d, got %d", i, *decoded)
		}
	})
}
