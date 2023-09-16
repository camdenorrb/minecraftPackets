package primitive

import (
	"errors"
	"io"
	"regexp"
)

//region MCString

type MCString string

func (s MCString) ByteLength() int {
	length := len([]byte(s))
	return VarInt(length).ByteLength() + len([]byte(s))
}

func (s MCString) Encode() []byte {
	// Length of string bytes + string bytes
	return append(VarInt(len([]byte(s))).Encode(), []byte(s)...)
}

func DecodeMCString(input io.ByteReader) (*string, error) {

	lengthVarInt, err := DecodeVarInt(input)
	if err != nil {
		return nil, err
	}

	length := int(*lengthVarInt)
	read := 0
	stringBytes := make([]byte, length)

	for read < length {

		currentByte, err := input.ReadByte()
		if err != nil {
			return nil, err
		}

		stringBytes[read] = currentByte
		read++
	}

	value := string(stringBytes)
	return &value, nil
}

//endregion

//region Var length integers

const VarSegmentBits = 0x7F
const VarContinueBit = 0x80

type VarInt int32

func (i VarInt) ByteLength() int {

	currentValue := uint32(i)

	var length int

	for {

		length++

		if currentValue&^VarSegmentBits == 0 {
			break
		}

		currentValue >>= 7

	}

	return length
}

func (i VarInt) Encode() []byte {

	currentValue := uint32(i)

	var bytes []byte

	for {

		if currentValue&^VarSegmentBits == 0 {
			bytes = append(bytes, byte(currentValue))
			break
		}

		bytes = append(bytes, byte((currentValue&VarSegmentBits)|VarContinueBit))
		currentValue >>= 7

	}

	return bytes
}

func DecodeVarInt(input io.ByteReader) (*VarInt, error) {

	value := int32(0)
	position := 0

	for {

		currentByte, err := input.ReadByte()
		if err != nil {
			return nil, err
		}

		value |= (int32(currentByte) & VarSegmentBits) << position

		if (currentByte & VarContinueBit) == 0 {
			break
		}

		position += 7

		if position >= 32 {
			return nil, errors.New("VarInt is too big")
		}
	}

	asVarInt := VarInt(value)
	return &asVarInt, nil
}

type VarLong int64

func (l VarLong) ByteLength() int {

	currentValue := uint64(l)

	var length int

	for {

		length++

		if currentValue&^VarSegmentBits == 0 {
			break
		}

		currentValue >>= 7

	}

	return length
}

func (l VarLong) Encode() []byte {

	currentValue := uint64(l)

	var bytes []byte

	for {

		if currentValue&^VarSegmentBits == 0 {
			bytes = append(bytes, byte(currentValue))
			break
		}

		bytes = append(bytes, byte((currentValue&VarSegmentBits)|VarContinueBit))
		currentValue >>= 7

	}

	return bytes

}

func DecodeVarLong(input io.ByteReader) (*VarLong, error) {

	value := int64(0)
	position := 0

	for {

		currentByte, err := input.ReadByte()
		if err != nil {
			return nil, err
		}

		value |= (int64(currentByte) & VarSegmentBits) << position

		if (currentByte & VarContinueBit) == 0 {
			break
		}

		position += 7

		if position >= 64 {
			return nil, errors.New("VarLong is too big")
		}
	}

	asVarLong := VarLong(value)
	return &asVarLong, nil
}

//endregion

type Identifier string

var IdentifierNameSpaceRegex = regexp.MustCompile(`^[a-z0-9._-]+$`)
var IdentifierValueRegex = regexp.MustCompile(`^[a-z0-9/._-]+$`)

func (i Identifier) ValidateNameSpace() error {

	if !IdentifierNameSpaceRegex.MatchString(string(i)) {
		return errors.New("invalid identifier namespace")
	}

	return nil
}

func (i Identifier) ValidateValue() error {

	if !IdentifierValueRegex.MatchString(string(i)) {
		return errors.New("invalid identifier value")
	}

	return nil
}
