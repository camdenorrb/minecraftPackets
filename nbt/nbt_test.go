package nbt

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
)

func FuzzEncodeNBT_All(f *testing.F) {

	f.Add("MeowName", byte(1), int16(2), int32(3), int64(4), float32(5), float64(6), "MeowString", []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "SubName", true)
	f.Add("NameMeow", byte(6), int16(5), int32(4), int64(3), float32(2), float64(1), "StringMeow", []byte{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, "SubName", false)

	f.Fuzz(func(t *testing.T, name string, byteNum byte, shortNum int16, intNum int32, longNum int64, floatNum float32, doubleNum float64, stringTag string, byteArray []byte, subName string, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		primitiveTags := []Tag{
			ByteTag(byteNum),
			ShortTag(shortNum),
			IntTag(intNum),
			LongTag(longNum),
			FloatTag(floatNum),
			DoubleTag(doubleNum),
			StringTag(stringTag),
			ByteArrayTag(byteArray),
		}

		// Scramble the order of the primitive tags
		for i := range primitiveTags {
			j := rand.Intn(i + 1)
			primitiveTags[i], primitiveTags[j] = primitiveTags[j], primitiveTags[i]
		}

		listTag := ListTag([]Tag{
			ByteTag(byteNum),
			ByteTag(byteNum),
			ByteTag(byteNum),
			ByteTag(byteNum),
			ByteTag(byteNum),
		})

		subNBT := NBT{
			Name: subName,
			Tags: primitiveTags,
		}

		nbt := NBT{
			Name: name,
			Tags: []Tag{
				&subNBT,
				listTag,
			},
		}

		nbt.Tags = append(nbt.Tags, primitiveTags...)

		output := bytes.Buffer{}
		err := nbt.PushToWriter(&output, endian, true)
		assert.NoError(t, err)

		reader := bytes.NewReader(output.Bytes())
		parsedNBT, err := readNBT(reader, endian)
		assert.NoError(t, err)

		// Validate that the NBT we read back is the same as the one we wrote
		assert.Equal(t, nbt.Name, parsedNBT.Name)
		assert.Len(t, nbt.Tags, len(parsedNBT.Tags))

		for i, tag := range nbt.Tags {

			testOutput := bytes.Buffer{}

			err := tag.PushToWriter(&testOutput, endian, true)
			assert.NoError(t, err)

			testOutput2 := bytes.Buffer{}
			err = parsedNBT.Tags[i].PushToWriter(&testOutput2, endian, true)
			assert.NoError(t, err)

			assert.EqualValues(t, testOutput.Bytes(), testOutput2.Bytes())
		}
	})
}

func FuzzEncodeNBT(f *testing.F) {

	f.Add("Meow1", true)
	f.Add("Meow2", false)
	f.Add("Meow3", true)

	f.Fuzz(func(t *testing.T, name string, bigEndian bool) {

		nbt := NBT{Name: name}

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		testOutput := bytes.Buffer{}
		err := nbt.PushToWriter(&testOutput, endian, true)
		assert.NoError(t, err)

		expectedByteBuffer := NewByteBufferWithNames()
		expectedByteBuffer.WriteByte(t, 0xa, "Compound tag")
		expectedByteBuffer.WriteUInt16(t, uint16(len([]byte(name))), "Name length", endian)
		expectedByteBuffer.WriteBytes(t, []byte(name), "Name")
		expectedByteBuffer.WriteByte(t, 0x0, "End tag")

		expectedByteBuffer.Validate(t, testOutput.Bytes())

		reader := bytes.NewReader(testOutput.Bytes())
		parsedNBT, err := readNBT(reader, endian)
		assert.NoError(t, err)

		// Validate that the NBT is the same
		assert.Equal(t, nbt.Name, parsedNBT.Name)
		assert.Len(t, nbt.Tags, len(parsedNBT.Tags))
	})
}

func FuzzEncodeSubNBT_Byte(f *testing.F) {

	f.Add("Meow1", "SubName1", byte(1), true)
	f.Add("Meow2", "SubName2", byte(2), false)
	f.Add("Meow3", "SubName3", byte(3), true)

	f.Fuzz(func(t *testing.T, name, subName string, byteNum byte, bigEndian bool) {

		subNBT := NBT{
			Name: subName,
			Tags: []Tag{ByteTag(byteNum)},
		}

		nbt := NBT{
			Name: name,
			Tags: []Tag{&subNBT},
		}

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		testOutput := bytes.Buffer{}
		err := nbt.PushToWriter(&testOutput, endian, true)
		assert.NoError(t, err)

		expectedByteBuffer := NewByteBufferWithNames()
		expectedByteBuffer.WriteByte(t, 0xa, "Compound Tag")
		expectedByteBuffer.WriteUInt16(t, uint16(len([]byte(name))), "Name Length", endian)
		expectedByteBuffer.WriteBytes(t, []byte(name), "Name")
		expectedByteBuffer.WriteByte(t, 0xa, "Sub Compound Tag")
		expectedByteBuffer.WriteUInt16(t, uint16(len([]byte(subName))), "Sub Name Length", endian)
		expectedByteBuffer.WriteBytes(t, []byte(subName), "Sub Name")
		expectedByteBuffer.WriteByte(t, 0x01, "Byte Tag")
		expectedByteBuffer.WriteByte(t, byteNum, "Byte Value")
		expectedByteBuffer.WriteByte(t, 0x00, "Sub End Tag")
		expectedByteBuffer.WriteByte(t, 0x00, "End Tag")

		expectedByteBuffer.Validate(t, testOutput.Bytes())

		reader := bytes.NewReader(testOutput.Bytes())
		parsedNBT, err := readNBT(reader, endian)
		assert.NoError(t, err)

		// Validate that the NBT is the same
		assert.Equal(t, nbt.Name, parsedNBT.Name)
		assert.Len(t, nbt.Tags, len(parsedNBT.Tags))

	})
}

func FuzzListTag_Single_Byte(f *testing.F) {

	f.Add(byte(1), true, true)
	f.Add(byte(2), false, false)

	f.Fuzz(func(t *testing.T, byteNum byte, bigEndian, includeTagID bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		listTag := ListTag([]Tag{ByteTag(byteNum)})

		testOutput := bytes.Buffer{}
		err := listTag.PushToWriter(&testOutput, endian, includeTagID)
		assert.NoError(t, err)

		expectedByteBuffer := NewByteBufferWithNames()
		if includeTagID {
			expectedByteBuffer.WriteByte(t, 0x09, "List Tag")
		}
		expectedByteBuffer.WriteByte(t, 0x01, "Byte Tag")
		expectedByteBuffer.WriteInt32(t, 1, "Length", endian)
		expectedByteBuffer.WriteByte(t, byteNum, "Byte Value")

		expectedByteBuffer.Validate(t, testOutput.Bytes())

		startIndex := 1
		if !includeTagID {
			startIndex = 0
		}

		reader := bytes.NewReader(testOutput.Bytes()[startIndex:])
		parsedListTag, err := readListTag(reader, endian)
		assert.NoError(t, err)

		assert.Len(t, listTag, len(parsedListTag))
	})
}

type ByteBufferWithNames struct {
	Buffer *bytes.Buffer
	names  []string
}

func NewByteBufferWithNames() *ByteBufferWithNames {
	return &ByteBufferWithNames{
		Buffer: &bytes.Buffer{},
	}
}

func (b *ByteBufferWithNames) Validate(t *testing.T, actual []byte) {

	bufferBytes := b.Buffer.Bytes()

	assert.Len(t, b.names, len(actual), "Expected equal number of names and bytes")

	fmt.Println("Expected:", b.Buffer.Bytes())
	fmt.Println("Actual:  ", actual)

	for i := 0; i < len(actual); i++ {
		assert.Equal(t, bufferBytes[i], actual[i], "Expected equal bytes at index "+strconv.Itoa(i)+" ["+b.names[i]+"]")
	}

	assert.Equal(t, b.Buffer.Bytes(), actual, "Expected equal bytes")
}

func (b *ByteBufferWithNames) Write(t *testing.T, p []byte, name string) {

	for i := 0; i < len(p); i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	wrote, err := b.Buffer.Write(p)
	assert.NoError(t, err, "Expected no error writing bytes")
	assert.Len(t, p, wrote, "Expected to write all bytes")
}

func (b *ByteBufferWithNames) WriteByte(t *testing.T, c byte, name string) {
	b.names = append(b.names, name)
	assert.NoError(t, b.Buffer.WriteByte(c))
}

func (b *ByteBufferWithNames) WriteInt16(t *testing.T, i int16, name string, endian Endian) {

	for i := 1; i <= 2; i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeShort(b.Buffer, i, endian))
}

func (b *ByteBufferWithNames) WriteUInt16(t *testing.T, i uint16, name string, endian Endian) {

	for i := 1; i <= 2; i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeUShort(b.Buffer, i, endian))
}

func (b *ByteBufferWithNames) WriteInt32(t *testing.T, i int32, name string, endian Endian) {

	for i := 1; i <= 4; i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeInt(b.Buffer, i, endian))
}

func (b *ByteBufferWithNames) WriteUInt32(t *testing.T, i uint32, name string, endian Endian) {

	for i := 1; i <= 4; i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeUInt(b.Buffer, i, endian))
}

func (b *ByteBufferWithNames) WriteInt64(t *testing.T, i int64, name string, endian Endian) {

	for i := 1; i <= 8; i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeLong(b.Buffer, i, endian))
}

func (b *ByteBufferWithNames) WriteUInt64(t *testing.T, i uint64, name string, endian Endian) {

	for i := 1; i <= 8; i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeULong(b.Buffer, i, endian))
}

func (b *ByteBufferWithNames) WriteFloat32(t *testing.T, f float32, name string, endian Endian) {

	for i := 1; i <= 4; i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeFloat(b.Buffer, f, endian))
}

func (b *ByteBufferWithNames) WriteFloat64(t *testing.T, f float64, name string, endian Endian) {

	for i := 1; i <= 8; i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeDouble(b.Buffer, f, endian))
}

func (b *ByteBufferWithNames) WriteBytes(t *testing.T, array []byte, name string) {

	for i := 0; i < len(array); i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeBytes(b.Buffer, array))
}

// Testing helpers

func printOutBytesAsHex(byteArray []byte) {
	for _, b := range byteArray {
		fmt.Printf("%#02x ", b)
	}
	fmt.Println()
}
