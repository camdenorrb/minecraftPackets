package primitive

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func FuzzMCString_ByteLength(f *testing.F) {
	f.Add("")
	f.Add("Meow")
	f.Add("Hello, World!")
	f.Fuzz(func(t *testing.T, s string) {
		length := len([]byte(s))
		assert.Equal(t, len(VarInt(length).Encode())+length, MCString(s).ByteLength())
	})
}

func FuzzMCString_Encode_Decode(f *testing.F) {
	f.Add("")
	f.Add("Meow")
	f.Add("Hello, World!")
	f.Fuzz(func(t *testing.T, s string) {
		encoded := bytes.NewBuffer(MCString(s).Encode())
		decoded, err := DecodeMCString(encoded)
		assert.NoError(t, err)
		assert.Equal(t, *decoded, s)
	})
}

func FuzzVarInt_Encode_Decode(f *testing.F) {
	f.Add(0)
	f.Add(math.MinInt32)
	f.Add(math.MaxInt32)
	f.Fuzz(func(t *testing.T, i int) {
		encoded := bytes.NewBuffer(VarInt(i).Encode())
		decoded, err := DecodeVarInt(encoded)
		assert.NoError(t, err)
		assert.Equal(t, *decoded, VarInt(i))
	})
}

func FuzzVarInt_ByteLength(f *testing.F) {
	f.Add(0)
	f.Add(math.MinInt32)
	f.Add(math.MaxInt32)
	f.Fuzz(func(t *testing.T, i int) {
		expected := VarInt(i).ByteLength()
		encoded := VarInt(i).Encode()
		assert.Equal(t, len(encoded), expected)
	})
}

func FuzzVarLong_Encode_Decode(f *testing.F) {
	f.Add(0)
	f.Add(math.MinInt64)
	f.Add(math.MaxInt64)
	f.Fuzz(func(t *testing.T, i int) {
		encoded := bytes.NewBuffer(VarLong(i).Encode())
		decoded, err := DecodeVarLong(encoded)
		assert.NoError(t, err)
		assert.Equal(t, *decoded, VarLong(i))
	})
}

func FuzzVarLong_ByteLength(f *testing.F) {
	f.Add(0)
	f.Add(math.MinInt64)
	f.Add(math.MaxInt64)
	f.Fuzz(func(t *testing.T, i int) {
		expected := VarLong(i).ByteLength()
		encoded := VarLong(i).Encode()
		assert.Equal(t, len(encoded), expected)
	})
}
