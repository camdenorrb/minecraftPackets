package nbt

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func FuzzEncodeNBT_All(f *testing.F) {

	f.Add("MeowName", byte(1), int16(2), int32(3), int64(4), float32(5), float64(6), "MeowString", []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "SubName")

	f.Fuzz(func(t *testing.T, name string, byteNum byte, shortNum int16, intNum int32, longNum int64, floatNum float32, doubleNum float64, stringTag string, byteArray []byte, subName string) {

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
		err := nbt.PushToWriter(&output, BigEndian, true)
		if err != nil {
			t.Fatal(err)
		}

		reader := bytes.NewReader(output.Bytes())
		parsedNBT, err := readNBT(reader, BigEndian)
		if err != nil {
			t.Error(err)
		}

		// Validate that the NBT we read back is the same as the one we wrote
		if nbt.Name != parsedNBT.Name {
			t.Errorf("Expected Name to be %s, got %s", nbt.Name, parsedNBT.Name)
		}

		if len(nbt.Tags) != len(parsedNBT.Tags) {
			t.Errorf("Expected Tags length to be %d, got %d", len(nbt.Tags), len(parsedNBT.Tags))
		}

		for i, tag := range nbt.Tags {

			testOutput := bytes.Buffer{}

			err := tag.PushToWriter(&testOutput, BigEndian, true)
			if err != nil {
				return
			}

			testOutput2 := bytes.Buffer{}
			err = parsedNBT.Tags[i].PushToWriter(&testOutput2, BigEndian, true)

			if !bytes.Equal(testOutput.Bytes(), testOutput2.Bytes()) {
				t.Errorf("Expected Tag %d to be %v, got %v", i, testOutput.Bytes(), testOutput2.Bytes())
			}
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

		expectedByteBuffer := bytes.Buffer{}
		expectedByteBuffer.WriteByte(0xa) // Compound
		err = writeUShort(&expectedByteBuffer, uint16(len([]byte(name))), endian)
		assert.NoError(t, err)
		expectedByteBuffer.WriteString(name)
		expectedByteBuffer.WriteByte(0x00) // End

		testOutputBytes := testOutput.Bytes()
		expectedBytes := expectedByteBuffer.Bytes()

		assert.Equal(t, expectedBytes[0], testOutputBytes[0], "Expected compound tag")
		assert.Equal(t, expectedBytes[1], testOutputBytes[1], "Expected name length (1)")
		assert.Equal(t, expectedBytes[2], testOutputBytes[2], "Expected name length (2)")
		assert.Equal(t, expectedBytes[3:3+len([]byte(name))], testOutputBytes[3:3+len([]byte(name))], "Expected equal name")
		assert.Equal(t, expectedBytes[len(expectedBytes)-1], testOutputBytes[len(testOutputBytes)-1], "Expected end tag")

		assert.Equal(t, expectedBytes, testOutputBytes, "Expected equal bytes")

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
		testOutputBytes := testOutput.Bytes()

		expectedByteBuffer := bytes.Buffer{}
		expectedByteBuffer.WriteByte(0xa) // Compound
		err = writeUShort(&expectedByteBuffer, uint16(len([]byte(name))), endian)
		assert.NoError(t, err)
		expectedByteBuffer.WriteString(name)
		expectedByteBuffer.WriteByte(0xa) // Compound
		err = writeUShort(&expectedByteBuffer, uint16(len([]byte(subName))), endian)
		assert.NoError(t, err)
		expectedByteBuffer.WriteString(subName)
		expectedByteBuffer.WriteByte(0x01) // Byte Tag
		expectedByteBuffer.WriteByte(byteNum)
		expectedByteBuffer.WriteByte(0x00) // End
		expectedByteBuffer.WriteByte(0x00) // End

		expectedBytes := expectedByteBuffer.Bytes()

		assert.Equal(t, expectedBytes[0], testOutputBytes[0], "Expected compound tag")
		assert.Equal(t, expectedBytes[1], testOutputBytes[1], "Expected name length (1)")
		assert.Equal(t, expectedBytes[2], testOutputBytes[2], "Expected name length (2)")
		assert.Equal(t, expectedBytes[3:3+len([]byte(name))], testOutputBytes[3:3+len([]byte(name))], "Expected equal name")
		nameEndIndex := 3 + len([]byte(name))
		assert.Equal(t, expectedBytes[nameEndIndex], testOutputBytes[nameEndIndex], "Expected compound tag")
		assert.Equal(t, expectedBytes[nameEndIndex+1], testOutputBytes[nameEndIndex+1], "Expected sub name length (1)")
		assert.Equal(t, expectedBytes[nameEndIndex+2], testOutputBytes[nameEndIndex+2], "Expected sub name length (2)")
		assert.Equal(t, expectedBytes[nameEndIndex+3:nameEndIndex+3+len([]byte(subName))], testOutputBytes[nameEndIndex+3:nameEndIndex+3+len([]byte(subName))], "Expected equal name")
		subNameEndIndex := nameEndIndex + 3 + len([]byte(subName))
		assert.Equal(t, expectedBytes[subNameEndIndex], testOutputBytes[subNameEndIndex], "Expected byte tag")
		assert.Equal(t, expectedBytes[subNameEndIndex+1], testOutputBytes[subNameEndIndex+1], "Expected byte value")
		assert.Equal(t, expectedBytes[subNameEndIndex+2], testOutputBytes[subNameEndIndex+2], "Expected end tag")
		assert.Equal(t, expectedBytes[len(expectedBytes)-1], testOutputBytes[len(testOutputBytes)-1], "Expected end tag")

		assert.Equal(t, expectedBytes, testOutputBytes, "Expected equal bytes")

		reader := bytes.NewReader(testOutput.Bytes())
		parsedNBT, err := readNBT(reader, endian)
		assert.NoError(t, err)

		// Validate that the NBT is the same
		assert.Equal(t, nbt.Name, parsedNBT.Name)
		assert.Len(t, nbt.Tags, len(parsedNBT.Tags))

	})
}

func FuzzListTag_Single_Byte(f *testing.F) {

	f.Add(byte(1), true)
	f.Add(byte(2), false)

	f.Fuzz(func(t *testing.T, byteNum byte, bigEndian bool) {

		endian := LittleEndian
		if bigEndian {
			endian = BigEndian
		}

		listTag := ListTag([]Tag{ByteTag(byteNum)})

		testOutput := bytes.Buffer{}
		err := listTag.PushToWriter(&testOutput, endian, true)
		assert.NoError(t, err)

		expectedByteBuffer := bytes.Buffer{}
		expectedByteBuffer.WriteByte(0x09)             // List Tag
		expectedByteBuffer.WriteByte(0x01)             // Byte Tag
		err = writeInt(&expectedByteBuffer, 1, endian) // Length
		assert.NoError(t, err)
		expectedByteBuffer.WriteByte(byteNum) // Byte value

		expectedBytes := expectedByteBuffer.Bytes()

		assert.Equal(t, expectedBytes, testOutput.Bytes(), "Expected equal bytes")

		//reader := bytes.NewReader(testOutput.Bytes())
		testOutputBytes := testOutput.Bytes()
		assert.Equal(t, byte(0x09), testOutputBytes[0], "Expected List tag")
		assert.Equal(t, byte(0x01), testOutputBytes[1], "Expected Byte tag list type")
		// Write size
		lengthReader := bytes.NewReader(testOutputBytes[2 : 2+4])
		actualLength, err := readInt(lengthReader, endian)
		assert.NoError(t, err)
		assert.Equal(t, int32(1), actualLength, "Expected length of 1")
		assert.Equal(t, byteNum, testOutputBytes[6], "Expected byte value")

		reader := bytes.NewReader(testOutput.Bytes()[1:])
		parsedListTag, err := readListTag(reader, endian)
		assert.NoError(t, err)

		// Validate that the NBT is the same
		assert.Len(t, listTag, len(parsedListTag))
	})
}

// Testing helpers

func printOutBytesAsHex(byteArray []byte) {
	for _, b := range byteArray {
		fmt.Printf("%#02x ", b)
	}
	fmt.Println()
}
