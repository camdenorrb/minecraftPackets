package nbt

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"sort"
	"strconv"
	"testing"
)

func TestNBT_SubCompound(t *testing.T) {

	tags := ListTag{}

	for i := 0; i < 10; i++ {
		tags = append(tags, CompoundTag{
			"Byte": ByteTag(byte(i)),
			"NBT": CompoundTag{
				"String": StringTag("Meow"),
				"Description": CompoundTag{
					"String": StringTag("Meow"),
				},
			},
		})
	}

	nbt := NBT{
		Name: nil,
		Tags: CompoundTag{
			"List": tags,
		},
	}

	output := bytes.Buffer{}
	err := nbt.PushToWriter(&output, BigEndian, true)
	assert.NoError(t, err)

	parsedNBT, err := ReadNBT(bufio.NewReader(bytes.NewReader(output.Bytes())), BigEndian, false)
	assert.NoError(t, err)

	validateEqualTagBytes(t, &nbt, parsedNBT, BigEndian, true)
}

func TestNBT_HelloWorld(t *testing.T) {

	helloWorldBytes, err := os.ReadFile("testdata/hello_world.nbt")
	assert.NoError(t, err)

	helloWorldFile, err := os.Open("testdata/hello_world.nbt")
	assert.NoError(t, err)

	nbt, err := ReadNBT(bufio.NewReader(helloWorldFile), BigEndian, true)
	assert.NoError(t, err)

	snbt, err := nbt.FormatSNBT()
	assert.NoError(t, err)

	assert.Equal(t, len(helloWorldBytes), nbt.Size(true))
	assert.Equal(t, "{name:\"Bananrama\"}", snbt)
}

func TestNBT_RegistryCodec(t *testing.T) {

	registryCodecFile, err := os.ReadFile("testdata/registry_codec.nbt")
	assert.NoError(t, err)

	endian := BigEndian

	gzipReader, err := gzip.NewReader(bytes.NewReader(registryCodecFile))
	assert.NoError(t, err)

	nbt, err := ReadNBT(bufio.NewReader(gzipReader), endian, true)
	assert.NoError(t, err)

	output := bytes.Buffer{}
	err = nbt.PushToWriter(&output, endian, true)
	assert.NoError(t, err)

	parsedNBT, err := ReadNBT(bufio.NewReader(bytes.NewReader(output.Bytes())), endian, true)
	assert.NoError(t, err)

	assert.Equal(t, parsedNBT.Size(true), output.Len())
	assert.Equal(t, nbt.Size(true), parsedNBT.Size(true))
	assert.Equal(t, nbt.Size(false), parsedNBT.Size(false))
	assert.Len(t, nbt.Tags, len(parsedNBT.Tags))
	validateEqualTagBytes(t, nbt, parsedNBT, endian, true)
}

func TestNBT_BigTest(t *testing.T) {

	bigNBTBytes, err := os.ReadFile("testdata/bigtest.nbt")
	assert.NoError(t, err)

	reader, err := gzip.NewReader(bytes.NewReader(bigNBTBytes))
	assert.NoError(t, err)

	endian := BigEndian

	nbt, err := ReadNBT(bufio.NewReader(reader), endian, true)
	assert.NoError(t, err)

	output := bytes.Buffer{}
	err = nbt.PushToWriter(&output, endian, true)
	assert.NoError(t, err)

	parsedNBT, err := ReadNBT(bufio.NewReader(bytes.NewReader(output.Bytes())), endian, true)
	assert.NoError(t, err)

	assert.Equal(t, parsedNBT.Size(true), output.Len())
	assert.Equal(t, nbt.Size(true), parsedNBT.Size(true))
	assert.Equal(t, nbt.Size(false), parsedNBT.Size(false))
	assert.Len(t, nbt.Tags, len(parsedNBT.Tags))
	validateEqualTagBytes(t, nbt, parsedNBT, endian, true)
}

func FuzzEncodeNBT_All(f *testing.F) {

	f.Add("MeowName", byte(1), int16(2), int32(3), int64(4), float32(5), float64(6), "MeowString", []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "SubName", true)
	f.Add("NameMeow", byte(6), int16(5), int32(4), int64(3), float32(2), float64(1), "StringMeow", []byte{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, "SubName", false)

	f.Fuzz(func(t *testing.T, name string, byteNum byte, shortNum int16, intNum int32, longNum int64, floatNum float32, doubleNum float64, stringTag string, byteArray []byte, subName string, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		intArray := []int32{intNum, intNum, intNum, intNum, intNum}
		longArray := []int64{longNum, longNum, longNum, longNum, longNum}

		int8Array := make([]int8, len(byteArray))
		for i, b := range byteArray {
			int8Array[i] = int8(b)
		}

		primitiveTags := map[string]Tag{
			"Byte":      ByteTag(byteNum),
			"Short":     ShortTag(shortNum),
			"Int":       IntTag(intNum),
			"Long":      LongTag(longNum),
			"Float":     FloatTag(floatNum),
			"Double":    DoubleTag(doubleNum),
			"String":    StringTag(stringTag),
			"ByteArray": ByteArrayTag(int8Array),
			"IntArray":  IntArrayTag(intArray),
			"LongArray": LongArrayTag(longArray),
		}

		listTag := ListTag([]Tag{
			ByteTag(byteNum),
			ByteTag(byteNum),
			ByteTag(byteNum),
			ByteTag(byteNum),
			ByteTag(byteNum),
		})

		subNBT := CompoundTag{}

		nbt := NBT{
			Name: &name,
			Tags: map[string]Tag{
				"SubNBT":  subNBT,
				"ListTag": listTag,
			},
		}

		// Add all the primitive tags to the NBT
		for name, tag := range primitiveTags {
			nbt.Tags[name] = tag
			subNBT[name] = tag
		}

		output := bytes.Buffer{}
		err := nbt.PushToWriter(&output, endian, true)
		assert.NoError(t, err)

		reader := bytes.NewReader(output.Bytes())
		parsedNBT, err := ReadNBT(bufio.NewReader(reader), endian, true)
		assert.NoError(t, err)

		validateEqualTagBytes(t, &nbt, parsedNBT, endian, true)
		assert.Equal(t, nbt.Size(true), output.Len())

		// Validate that the NBT we read back is the same as the one we wrote
		assert.Equal(t, nbt.Name, parsedNBT.Name)
		assert.Len(t, nbt.Tags, len(parsedNBT.Tags))

		validateEqualTagBytes(t, &nbt, parsedNBT, endian, true)
	})
}

func FuzzEncodeNBT(f *testing.F) {

	f.Add("Meow1", true)
	f.Add("Meow2", false)
	f.Add("Meow3", true)

	f.Fuzz(func(t *testing.T, name string, bigEndian bool) {

		nbt := NBT{Name: &name}

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		testOutput := bytes.Buffer{}
		err := nbt.PushToWriter(&testOutput, endian, true)
		assert.NoError(t, err)

		expectedByteBuffer := newByteBufferWithNames()
		expectedByteBuffer.WriteByte(t, 0xa, "Compound tag")
		expectedByteBuffer.WriteUInt16(t, uint16(len([]byte(name))), "Name length", endian)
		expectedByteBuffer.WriteBytes(t, []byte(name), "Name")
		expectedByteBuffer.WriteByte(t, 0x0, "End tag")

		expectedByteBuffer.Validate(t, testOutput.Bytes())

		reader := bytes.NewReader(testOutput.Bytes())
		parsedNBT, err := ReadNBT(bufio.NewReader(reader), endian, true)
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

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		subNBT := CompoundTag{
			"Byte": ByteTag(byteNum),
		}

		nbt := NBT{
			Name: &name,
			Tags: map[string]Tag{"SubNBT": &subNBT},
		}

		testOutput := bytes.Buffer{}
		err := nbt.PushToWriter(&testOutput, endian, true)
		assert.NoError(t, err)

		assert.Equal(t, nbt.Size(true), len(testOutput.Bytes()))

		expectedByteBuffer := newByteBufferWithNames()
		expectedByteBuffer.WriteByte(t, 0xa, "Compound Tag")
		expectedByteBuffer.WriteString(t, name, endian)
		expectedByteBuffer.WriteByte(t, 0xa, "Sub Compound Tag")
		expectedByteBuffer.WriteString(t, "SubNBT", endian)
		expectedByteBuffer.WriteByte(t, 0x01, "Byte Tag")
		expectedByteBuffer.WriteString(t, "Byte", endian)
		expectedByteBuffer.WriteByte(t, byteNum, "Byte Value")
		expectedByteBuffer.WriteByte(t, 0x00, "Sub End Tag")
		expectedByteBuffer.WriteByte(t, 0x00, "End Tag")

		expectedByteBuffer.Validate(t, testOutput.Bytes())

		reader := bytes.NewReader(testOutput.Bytes())
		parsedNBT, err := ReadNBT(bufio.NewReader(reader), endian, true)
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

		assert.Equal(t, listTag.Size(includeTagID), len(testOutput.Bytes()))

		expectedByteBuffer := newByteBufferWithNames()
		if includeTagID {
			expectedByteBuffer.WriteByte(t, 0x09, "List Tag")
		}
		expectedByteBuffer.WriteByte(t, 0x01, "Byte Tag")
		expectedByteBuffer.WriteInt32(t, 1, "Length", endian)
		expectedByteBuffer.WriteByte(t, byteNum, "Byte Value")

		expectedByteBuffer.Validate(t, testOutput.Bytes())

		startOfListIndex := 1
		if !includeTagID {
			startOfListIndex = 0
		}

		reader := bytes.NewReader(testOutput.Bytes()[startOfListIndex:])
		parsedListTag, err := readListTag(bufio.NewReader(reader), endian)
		assert.NoError(t, err)

		assert.Len(t, listTag, len(parsedListTag))
	})
}

type byteBufferWithNames struct {
	Buffer *bytes.Buffer
	names  []string
}

func newByteBufferWithNames() *byteBufferWithNames {
	return &byteBufferWithNames{
		Buffer: &bytes.Buffer{},
	}
}

func (b *byteBufferWithNames) Validate(t *testing.T, actual []byte) {

	bufferBytes := b.Buffer.Bytes()

	assert.Len(t, b.names, len(actual), "Expected equal number of names and bytes")

	if len(b.Buffer.Bytes()) != len(actual) {
		fmt.Println("Expected:", b.Buffer.Bytes())
		fmt.Println("Actual:  ", actual)
	}

	for i := 0; i < len(actual); i++ {
		assert.Equal(t, bufferBytes[i], actual[i], "Expected equal bytes at index "+strconv.Itoa(i)+" ["+b.names[i]+"]")
	}

	assert.Equal(t, b.Buffer.Bytes(), actual, "Expected equal bytes")
}

func (b *byteBufferWithNames) Write(t *testing.T, p []byte, name string) {

	for i := 0; i < len(p); i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	wrote, err := b.Buffer.Write(p)
	assert.NoError(t, err, "Expected no error writing bytes")
	assert.Len(t, p, wrote, "Expected to write all bytes")
}

func (b *byteBufferWithNames) WriteByte(t *testing.T, c byte, name string) {
	b.names = append(b.names, name)
	assert.NoError(t, b.Buffer.WriteByte(c))
}

func (b *byteBufferWithNames) WriteInt16(t *testing.T, i int16, name string, endian Endian) {

	for i := 1; i <= 2; i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeInt16(b.Buffer, i, endian))
}

func (b *byteBufferWithNames) WriteUInt16(t *testing.T, i uint16, name string, endian Endian) {

	for i := 1; i <= 2; i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeUInt16(b.Buffer, i, endian))
}

func (b *byteBufferWithNames) WriteInt32(t *testing.T, i int32, name string, endian Endian) {

	for i := 1; i <= 4; i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeInt32(b.Buffer, i, endian))
}

func (b *byteBufferWithNames) WriteUInt32(t *testing.T, i uint32, name string, endian Endian) {

	for i := 1; i <= 4; i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeUInt32(b.Buffer, i, endian))
}

func (b *byteBufferWithNames) WriteInt64(t *testing.T, i int64, name string, endian Endian) {

	for i := 1; i <= 8; i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeInt64(b.Buffer, i, endian))
}

func (b *byteBufferWithNames) WriteUInt64(t *testing.T, i uint64, name string, endian Endian) {

	for i := 1; i <= 8; i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeUInt64(b.Buffer, i, endian))
}

func (b *byteBufferWithNames) WriteFloat32(t *testing.T, f float32, name string, endian Endian) {

	for i := 1; i <= 4; i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeFloat32(b.Buffer, f, endian))
}

func (b *byteBufferWithNames) WriteFloat64(t *testing.T, f float64, name string, endian Endian) {

	for i := 1; i <= 8; i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeFloat64(b.Buffer, f, endian))
}

func (b *byteBufferWithNames) WriteBytes(t *testing.T, array []byte, name string) {

	for i := 0; i < len(array); i++ {
		b.names = append(b.names, name+" ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeBytes(b.Buffer, array))
}

func (b *byteBufferWithNames) WriteString(t *testing.T, input string, endian Endian) {

	// Short for length
	for i := 0; i < 2; i++ {
		b.names = append(b.names, "length of string '"+input+"' ("+strconv.Itoa(i)+")")
	}

	for i := 0; i < len([]byte(input)); i++ {
		b.names = append(b.names, "string '"+input+"' ("+strconv.Itoa(i)+")")
	}

	assert.NoError(t, writeMCString(b.Buffer, input, endian))
}

// Testing helpers

func validateEqualTagBytes(t *testing.T, expected, actual Tag, endian Endian, includeTagID bool) {

	assert.Equal(t, expected.ID(), actual.ID(), "Expected equal tag IDs")

	_, isCompound := expected.(CompoundTag)
	_, isNBT := expected.(*NBT)

	// If NBT or Compound
	if isCompound || isNBT {

		var expectedCompound CompoundTag
		var actualCompound CompoundTag

		if isCompound {
			expectedCompound = expected.(CompoundTag)
			actualCompound = actual.(CompoundTag)
		} else {
			expectedCompound = expected.(*NBT).Tags
			actualCompound = actual.(*NBT).Tags
		}

		// Assert same Tags keys
		var expectedKeys []string
		for k := range expectedCompound {
			expectedKeys = append(expectedKeys, k)
		}

		var actualKeys []string
		for k := range actualCompound {
			actualKeys = append(actualKeys, k)
		}

		// Sort both
		sort.Strings(expectedKeys)
		sort.Strings(actualKeys)

		assert.Equal(t, expectedKeys, actualKeys, "Expected equal keys in NBT tags")

		for name := range expectedCompound {
			validateEqualTagBytes(t, expectedCompound[name], actualCompound[name], endian, true)
		}

		return
	}

	expectedList, isList := expected.(ListTag)
	actualList, _ := actual.(ListTag)

	if isList && len(expectedList) != 0 {
		// If compound list, validate each compound
		if _, isCompound := (expectedList)[0].(CompoundTag); isCompound {
			for i := range expectedList {
				validateEqualTagBytes(t, (expectedList)[i], (actualList)[i], endian, true)
			}
			return
		}
	}

	// Buffers
	expectedBuffer := bytes.Buffer{}
	actualBuffer := bytes.Buffer{}

	// Write expected
	assert.NoError(t, expected.PushToWriter(&expectedBuffer, endian, includeTagID))
	assert.NoError(t, actual.PushToWriter(&actualBuffer, endian, includeTagID))

	// Assert equal
	assert.Equal(t, expectedBuffer.Bytes(), actualBuffer.Bytes(), "Expected equal bytes, tag: "+strconv.Itoa(int(expected.ID())))
}

func printOutBytesAsHex(byteArray []byte) {
	for _, b := range byteArray {
		fmt.Printf("%#02x ", b)
	}
	fmt.Println()
}
