package nbt

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestSNBT_BigTest(t *testing.T) {

	bigNBTBytes, err := os.ReadFile("testdata/bigtest.nbt")
	assert.NoError(t, err)

	reader, err := gzip.NewReader(bytes.NewReader(bigNBTBytes))
	assert.NoError(t, err)

	endian := BigEndian

	nbt, err := readNBT(bufio.NewReader(reader), endian)
	assert.NoError(t, err)

	snbt, err := nbt.FormatSNBT()
	assert.NoError(t, err)
	assert.NotNil(t, snbt)
}

func FuzzParseSNBT(f *testing.F) {

	f.Add(byte(1), int16(2), int32(3), int64(4), float32(5), float64(6), "MeowString", []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	f.Add(byte(6), int16(5), int32(4), int64(3), float32(2), float64(1), "StringMeow", []byte{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})

	f.Fuzz(func(t *testing.T, byteNum byte, shortNum int16, intNum int32, longNum int64, floatNum float32, doubleNum float64, stringTag string, byteArray []byte) {

		intArray := []int32{intNum, intNum, intNum, intNum, intNum}
		longArray := []int64{longNum, longNum, longNum, longNum, longNum}

		int8Array := make([]int8, len(byteArray))
		for i, b := range byteArray {
			int8Array[i] = int8(b)
		}

		if len(int8Array) == 0 {
			int8Array = nil
		}

		// Fix non-printable characters
		stringTag = strings.ReplaceAll(stringTag, `\`, "\\")
		stringTag = fmt.Sprintf("%q", stringTag)
		stringTag = stringTag[1 : len(stringTag)-1]

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

		err := validateString(fmt.Sprintf("%q", stringTag))
		if err != nil {
			t.SkipNow()
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
			Name: "",
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

		snbt, err := nbt.FormatSNBT()
		assert.NoError(t, err)

		parsedNBT, err := ParseSNBT(snbt)
		assert.NoError(t, err)

		assert.Equal(t, nbt.Tags, parsedNBT.Tags)
	})
}
