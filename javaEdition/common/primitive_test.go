package common

import (
	"bytes"
	"math"
	"testing"
)

func FuzzVarInt_Encode_Decode(f *testing.F) {
	f.Add(0)
	f.Add(math.MinInt32)
	f.Add(math.MaxInt32)
	f.Fuzz(func(t *testing.T, i int) {
		encoded := bytes.NewBuffer(VarInt(i).Encode())
		decoded, err := DecodeVarInt(encoded)
		if err != nil {
			t.Error(err)
		}
		if *decoded != int32(i) {
			t.Errorf("Expected %d, got %d", i, *decoded)
		}
	})
}

func FuzzVarInt_ByteLength(f *testing.F) {
	f.Add(0)
	f.Add(math.MinInt32)
	f.Add(math.MaxInt32)
	f.Fuzz(func(t *testing.T, i int) {
		expected := VarInt(i).ByteLength()
		encoded := VarInt(i).Encode()
		if len(encoded) != expected {
			t.Errorf("Expected %d, got %d", expected, len(encoded))
		}
	})
}

func FuzzVarLong_Encode_Decode(f *testing.F) {
	f.Add(0)
	f.Add(math.MinInt64)
	f.Add(math.MaxInt64)
	f.Fuzz(func(t *testing.T, i int) {
		encoded := bytes.NewBuffer(VarLong(i).Encode())
		decoded, err := DecodeVarLong(encoded)
		if err != nil {
			t.Error(err)
		}
		if *decoded != int64(i) {
			t.Errorf("Expected %d, got %d", i, *decoded)
		}
	})
}

func FuzzVarLong_ByteLength(f *testing.F) {
	f.Add(0)
	f.Add(math.MinInt64)
	f.Add(math.MaxInt64)
	f.Fuzz(func(t *testing.T, i int) {
		expected := VarLong(i).ByteLength()
		encoded := VarLong(i).Encode()
		if len(encoded) != expected {
			t.Errorf("Expected %d, got %d", expected, len(encoded))
		}
	})
}
