package common

import (
	"errors"
	"regexp"
)

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

func DecodeVarInt(input []byte) (*int32, error) {

	value := int32(0)
	position := 0

	for _, currentByte := range input {

		value |= (int32(currentByte) & VarSegmentBits) << position

		if (currentByte & VarContinueBit) == 0 {
			break
		}

		position += 7

		if position >= 32 {
			return nil, errors.New("VarInt is too big")
		}
	}

	return &value, nil
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

func DecodeVarLong(input []byte) (*int64, error) {

	value := int64(0)
	position := 0

	for _, currentByte := range input {

		value |= (int64(currentByte) & VarSegmentBits) << position

		if (currentByte & VarContinueBit) == 0 {
			break
		}

		position += 7

		if position >= 64 {
			return nil, errors.New("VarLong is too big")
		}
	}

	return &value, nil
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
