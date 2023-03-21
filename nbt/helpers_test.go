package nbt

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func FuzzWriteBytes(f *testing.F) {

	f.Add([]byte{0, 1, 2, 3, 4, 5})
	f.Add([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	f.Add([]byte{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0})

	f.Fuzz(func(t *testing.T, values []byte) {
		buffer := bytes.Buffer{}
		err := writeBytes(&buffer, values)
		assert.NoError(t, err)

		if len(values) == 0 {
			assert.Nil(t, buffer.Bytes())
		} else {
			assert.Equal(t, values, buffer.Bytes())
		}
	})
}

func FuzzReadNBytes(f *testing.F) {

	f.Add([]byte{0, 1, 2, 3, 4, 5})
	f.Add([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	f.Add([]byte{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0})

	f.Fuzz(func(t *testing.T, values []byte) {

		reader := bytes.NewReader(values)

		nBytes, err := readNBytes(reader, len(values))
		assert.NoError(t, err)

		if len(values) == 0 {
			assert.Nil(t, nBytes)
		} else {
			assert.Equal(t, values, nBytes)
		}
	})
}

func FuzzReadInt16(f *testing.F) {

	f.Add(int16(0), false)
	f.Add(int16(1), true)
	f.Add(int16(2), false)
	f.Add(int16(3), true)

	f.Fuzz(func(t *testing.T, short int16, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		var data []byte
		if bigEndian {
			data = []byte{byte(short >> 8), byte(short)}
		} else {
			data = []byte{byte(short), byte(short >> 8)}
		}

		reader := bytes.NewReader(data)

		actualShort, err := readInt16(reader, endian)
		assert.NoError(t, err)

		assert.Equal(t, short, actualShort)
	})
}

func FuzzReadUInt16(f *testing.F) {

	f.Add(uint16(0), false)
	f.Add(uint16(1), true)
	f.Add(uint16(2), false)
	f.Add(uint16(3), true)

	f.Fuzz(func(t *testing.T, short uint16, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		var data []byte
		if bigEndian {
			data = []byte{byte(short >> 8), byte(short)}
		} else {
			data = []byte{byte(short), byte(short >> 8)}
		}

		reader := bytes.NewReader(data)

		actualShort, err := readUint16(reader, endian)
		assert.NoError(t, err)

		assert.Equal(t, short, actualShort)
	})
}

func FuzzReadInt32(f *testing.F) {

	f.Add(int32(0), false)
	f.Add(int32(1), true)
	f.Add(int32(2), false)
	f.Add(int32(3), true)

	f.Fuzz(func(t *testing.T, i int32, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		var data []byte
		if bigEndian {
			data = []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
		} else {
			data = []byte{byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24)}
		}

		reader := bytes.NewReader(data)

		actualInt, err := readInt32(reader, endian)
		assert.NoError(t, err)

		assert.Equal(t, i, actualInt)
	})
}

func FuzzReadUInt32(f *testing.F) {

	f.Add(uint32(0), false)
	f.Add(uint32(1), true)
	f.Add(uint32(2), false)
	f.Add(uint32(3), true)

	f.Fuzz(func(t *testing.T, i uint32, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		var data []byte
		if bigEndian {
			data = []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
		} else {
			data = []byte{byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24)}
		}

		reader := bytes.NewReader(data)

		actualInt, err := readUInt32(reader, endian)
		assert.NoError(t, err)

		assert.Equal(t, i, actualInt)
	})
}

func FuzzReadInt64(f *testing.F) {

	f.Add(int64(0), false)
	f.Add(int64(1), true)
	f.Add(int64(2), false)
	f.Add(int64(3), true)

	f.Fuzz(func(t *testing.T, l int64, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		var data []byte
		if bigEndian {
			data = []byte{byte(l >> 56), byte(l >> 48), byte(l >> 40), byte(l >> 32), byte(l >> 24), byte(l >> 16), byte(l >> 8), byte(l)}
		} else {
			data = []byte{byte(l), byte(l >> 8), byte(l >> 16), byte(l >> 24), byte(l >> 32), byte(l >> 40), byte(l >> 48), byte(l >> 56)}
		}

		reader := bytes.NewReader(data)

		actualLong, err := readInt64(reader, endian)
		assert.NoError(t, err)

		assert.Equal(t, l, actualLong)
	})
}

func FuzzReadUInt64(f *testing.F) {

	f.Add(uint64(0), false)
	f.Add(uint64(1), true)
	f.Add(uint64(2), false)
	f.Add(uint64(3), true)

	f.Fuzz(func(t *testing.T, l uint64, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		var data []byte
		if bigEndian {
			data = []byte{byte(l >> 56), byte(l >> 48), byte(l >> 40), byte(l >> 32), byte(l >> 24), byte(l >> 16), byte(l >> 8), byte(l)}
		} else {
			data = []byte{byte(l), byte(l >> 8), byte(l >> 16), byte(l >> 24), byte(l >> 32), byte(l >> 40), byte(l >> 48), byte(l >> 56)}
		}

		reader := bytes.NewReader(data)

		actualLong, err := readUInt64(reader, endian)
		assert.NoError(t, err)

		assert.Equal(t, l, actualLong)
	})
}

func FuzzReadFloat32(f *testing.F) {

	f.Add(float32(0), false)
	f.Add(float32(1), true)
	f.Add(float32(2), false)
	f.Add(float32(3), true)

	f.Fuzz(func(t *testing.T, f32 float32, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		asUint := math.Float32bits(f32)

		var data []byte
		if bigEndian {
			data = []byte{byte(asUint >> 24), byte(asUint >> 16), byte(asUint >> 8), byte(asUint)}
		} else {
			data = []byte{byte(asUint), byte(asUint >> 8), byte(asUint >> 16), byte(asUint >> 24)}
		}

		reader := bytes.NewReader(data)

		actualFloat, err := readFloat32(reader, endian)
		assert.NoError(t, err)

		assert.Equal(t, f32, actualFloat)
	})
}

func FuzzReadFloat64(f *testing.F) {

	f.Add(float64(0), false)
	f.Add(float64(1), true)
	f.Add(float64(2), false)
	f.Add(float64(3), true)

	f.Fuzz(func(t *testing.T, f64 float64, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		asUint := math.Float64bits(f64)

		var data []byte
		if bigEndian {
			data = []byte{byte(asUint >> 56), byte(asUint >> 48), byte(asUint >> 40), byte(asUint >> 32), byte(asUint >> 24), byte(asUint >> 16), byte(asUint >> 8), byte(asUint)}
		} else {
			data = []byte{byte(asUint), byte(asUint >> 8), byte(asUint >> 16), byte(asUint >> 24), byte(asUint >> 32), byte(asUint >> 40), byte(asUint >> 48), byte(asUint >> 56)}
		}

		reader := bytes.NewReader(data)

		actualDouble, err := readFloat64(reader, endian)
		assert.NoError(t, err)

		assert.Equal(t, f64, actualDouble)
	})
}

func FuzzWriteInt16(f *testing.F) {

	f.Add(int16(0), false)
	f.Add(int16(1), true)
	f.Add(int16(2), false)
	f.Add(int16(3), true)

	f.Fuzz(func(t *testing.T, s int16, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		var expected []byte
		if bigEndian {
			expected = []byte{byte(s >> 8), byte(s)}
		} else {
			expected = []byte{byte(s), byte(s >> 8)}
		}

		buffer := new(bytes.Buffer)
		err := writeInt16(buffer, s, endian)
		assert.NoError(t, err)

		assert.Equal(t, expected, buffer.Bytes())
	})
}

func FuzzWriteUInt16(f *testing.F) {

	f.Add(uint16(0), false)
	f.Add(uint16(1), true)
	f.Add(uint16(2), false)
	f.Add(uint16(3), true)

	f.Fuzz(func(t *testing.T, s uint16, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		var expected []byte
		if bigEndian {
			expected = []byte{byte(s >> 8), byte(s)}
		} else {
			expected = []byte{byte(s), byte(s >> 8)}
		}

		buffer := new(bytes.Buffer)
		err := writeUInt16(buffer, s, endian)
		assert.NoError(t, err)

		assert.Equal(t, expected, buffer.Bytes())
	})
}

func FuzzWriteInt32(f *testing.F) {

	f.Add(int32(0), false)
	f.Add(int32(1), true)
	f.Add(int32(2), false)
	f.Add(int32(3), true)

	f.Fuzz(func(t *testing.T, i int32, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		var expected []byte
		if bigEndian {
			expected = []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
		} else {
			expected = []byte{byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24)}
		}

		buffer := new(bytes.Buffer)
		err := writeInt32(buffer, i, endian)
		assert.NoError(t, err)

		assert.Equal(t, expected, buffer.Bytes())
	})
}

func FuzzWriteUInt32(f *testing.F) {

	f.Add(uint32(0), false)
	f.Add(uint32(1), true)
	f.Add(uint32(2), false)
	f.Add(uint32(3), true)

	f.Fuzz(func(t *testing.T, i uint32, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		var expected []byte
		if bigEndian {
			expected = []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
		} else {
			expected = []byte{byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24)}
		}

		buffer := new(bytes.Buffer)
		err := writeUInt32(buffer, i, endian)
		assert.NoError(t, err)

		assert.Equal(t, expected, buffer.Bytes())
	})
}

func FuzzWriteInt64(f *testing.F) {

	f.Add(int64(0), false)
	f.Add(int64(1), true)
	f.Add(int64(2), false)
	f.Add(int64(3), true)

	f.Fuzz(func(t *testing.T, l int64, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		var expected []byte
		if bigEndian {
			expected = []byte{byte(l >> 56), byte(l >> 48), byte(l >> 40), byte(l >> 32), byte(l >> 24), byte(l >> 16), byte(l >> 8), byte(l)}
		} else {
			expected = []byte{byte(l), byte(l >> 8), byte(l >> 16), byte(l >> 24), byte(l >> 32), byte(l >> 40), byte(l >> 48), byte(l >> 56)}
		}

		buffer := new(bytes.Buffer)
		err := writeInt64(buffer, l, endian)
		assert.NoError(t, err)

		assert.Equal(t, expected, buffer.Bytes())
	})
}

func FuzzWriteUInt64(f *testing.F) {

	f.Add(uint64(0), false)
	f.Add(uint64(1), true)
	f.Add(uint64(2), false)
	f.Add(uint64(3), true)

	f.Fuzz(func(t *testing.T, l uint64, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		var expected []byte
		if bigEndian {
			expected = []byte{byte(l >> 56), byte(l >> 48), byte(l >> 40), byte(l >> 32), byte(l >> 24), byte(l >> 16), byte(l >> 8), byte(l)}
		} else {
			expected = []byte{byte(l), byte(l >> 8), byte(l >> 16), byte(l >> 24), byte(l >> 32), byte(l >> 40), byte(l >> 48), byte(l >> 56)}
		}

		buffer := new(bytes.Buffer)
		err := writeUInt64(buffer, l, endian)
		assert.NoError(t, err)

		assert.Equal(t, expected, buffer.Bytes())
	})
}

func FuzzWriteFloat32(f *testing.F) {

	f.Add(float32(0), false)
	f.Add(float32(1), true)
	f.Add(float32(2), false)
	f.Add(float32(3), true)

	f.Fuzz(func(t *testing.T, f32 float32, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		asUint := math.Float32bits(f32)

		var expected []byte
		if bigEndian {
			expected = []byte{byte(asUint >> 24), byte(asUint >> 16), byte(asUint >> 8), byte(asUint)}
		} else {
			expected = []byte{byte(asUint), byte(asUint >> 8), byte(asUint >> 16), byte(asUint >> 24)}
		}

		buffer := new(bytes.Buffer)
		err := writeFloat32(buffer, f32, endian)
		assert.NoError(t, err)

		assert.Equal(t, expected, buffer.Bytes())
	})
}

func FuzzWriteFloat64(f *testing.F) {

	f.Add(float64(0), false)
	f.Add(float64(1), true)
	f.Add(float64(2), false)
	f.Add(float64(3), true)

	f.Fuzz(func(t *testing.T, f64 float64, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		asUint := math.Float64bits(f64)

		var expected []byte
		if bigEndian {
			expected = []byte{byte(asUint >> 56), byte(asUint >> 48), byte(asUint >> 40), byte(asUint >> 32), byte(asUint >> 24), byte(asUint >> 16), byte(asUint >> 8), byte(asUint)}
		} else {
			expected = []byte{byte(asUint), byte(asUint >> 8), byte(asUint >> 16), byte(asUint >> 24), byte(asUint >> 32), byte(asUint >> 40), byte(asUint >> 48), byte(asUint >> 56)}
		}

		buffer := new(bytes.Buffer)
		err := writeFloat64(buffer, f64, endian)
		assert.NoError(t, err)

		assert.Equal(t, expected, buffer.Bytes())
	})
}
