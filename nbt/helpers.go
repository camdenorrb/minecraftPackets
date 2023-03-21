package nbt

import (
	"bytes"
	"encoding/binary"
	"github.com/joomcode/errorx"
	"io"
	"math"
)

func writeBytes(writer io.ByteWriter, data []byte) error {
	for _, b := range data {
		err := writer.WriteByte(b)
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write bytes")
		}
	}
	return nil
}

func readNBytes(reader *bytes.Reader, n int) ([]byte, error) {

	if n == 0 {
		return nil, nil
	}

	data := make([]byte, n)

	amount, err := reader.Read(data)
	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "Failed to read %d bytes", n)
	}
	if amount != n {
		return nil, errorx.IllegalState.New("Failed to read %d bytes, only read %d", n, amount)
	}

	return data, nil
}

func readInt16(reader *bytes.Reader, endian Endian) (int16, error) {

	data, err := readNBytes(reader, 2)
	if err != nil {
		return 0, errorx.IllegalState.Wrap(err, "failed to read short")
	}

	if endian == BigEndian {
		return int16(data[0])<<8 | int16(data[1]), nil
	}

	return int16(data[0]) | int16(data[1])<<8, nil
}

func readUint16(reader *bytes.Reader, endian Endian) (uint16, error) {

	data, err := readNBytes(reader, 2)
	if err != nil {
		return 0, errorx.IllegalState.Wrap(err, "failed to read short")
	}

	if endian == BigEndian {
		return binary.BigEndian.Uint16(data), nil
	}

	return binary.LittleEndian.Uint16(data), nil
}

func readInt32(reader *bytes.Reader, endian Endian) (int32, error) {

	data, err := readNBytes(reader, 4)
	if err != nil {
		return 0, errorx.IllegalState.Wrap(err, "failed to read int")
	}

	if endian == BigEndian {
		return int32(data[0])<<24 | int32(data[1])<<16 | int32(data[2])<<8 | int32(data[3]), nil
	}

	return int32(data[0]) | int32(data[1])<<8 | int32(data[2])<<16 | int32(data[3])<<24, nil
}

func readUInt32(reader *bytes.Reader, endian Endian) (uint32, error) {

	data, err := readNBytes(reader, 4)
	if err != nil {
		return 0, errorx.IllegalState.Wrap(err, "failed to read uint")
	}

	if endian == BigEndian {
		return binary.BigEndian.Uint32(data), nil
	}

	return binary.LittleEndian.Uint32(data), nil
}

func readInt64(reader *bytes.Reader, endian Endian) (int64, error) {

	data, err := readNBytes(reader, 8)
	if err != nil {
		return 0, errorx.IllegalState.Wrap(err, "failed to read long")
	}

	if endian == BigEndian {
		return int64(data[0])<<56 | int64(data[1])<<48 | int64(data[2])<<40 | int64(data[3])<<32 |
			int64(data[4])<<24 | int64(data[5])<<16 | int64(data[6])<<8 | int64(data[7]), nil
	}

	return int64(data[0]) | int64(data[1])<<8 | int64(data[2])<<16 | int64(data[3])<<24 |
		int64(data[4])<<32 | int64(data[5])<<40 | int64(data[6])<<48 | int64(data[7])<<56, nil
}

func readUInt64(reader *bytes.Reader, endian Endian) (uint64, error) {

	data, err := readNBytes(reader, 8)
	if err != nil {
		return 0, errorx.IllegalState.Wrap(err, "failed to read long")
	}

	if endian == BigEndian {
		return binary.BigEndian.Uint64(data), nil
	}

	return binary.LittleEndian.Uint64(data), nil
}

func readFloat32(reader *bytes.Reader, endian Endian) (float32, error) {

	data, err := readNBytes(reader, 4)
	if err != nil {
		return 0, errorx.IllegalState.Wrap(err, "failed to read float")
	}

	if endian == BigEndian {
		return math.Float32frombits(binary.BigEndian.Uint32(data)), nil
	}

	return math.Float32frombits(binary.LittleEndian.Uint32(data)), nil
}

func readFloat64(reader *bytes.Reader, endian Endian) (float64, error) {

	data, err := readNBytes(reader, 8)
	if err != nil {
		return 0, errorx.IllegalState.Wrap(err, "failed to read double")
	}

	if endian == BigEndian {
		return math.Float64frombits(binary.BigEndian.Uint64(data)), nil
	}

	return math.Float64frombits(binary.LittleEndian.Uint64(data)), nil
}

func readString(reader *bytes.Reader, endian Endian) (string, error) {

	length, err := readUint16(reader, endian)
	if err != nil {
		return "", errorx.IllegalState.Wrap(err, "failed to read string length")
	}

	data, err := readNBytes(reader, int(length))
	if err != nil {
		return "", errorx.IllegalState.Wrap(err, "failed to read string")
	}

	return string(data), nil
}

func writeInt16(writer io.ByteWriter, value int16, endian Endian) error {

	var err error

	if endian == BigEndian {
		err = writeBytes(writer, []byte{byte(value >> 8), byte(value)})
	} else {
		err = writeBytes(writer, []byte{byte(value), byte(value >> 8)})
	}

	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write short")
	}

	return nil
}

func writeUInt16(writer io.ByteWriter, value uint16, endian Endian) error {

	var err error

	if endian == BigEndian {
		err = writeBytes(writer, []byte{byte(value >> 8), byte(value)})
	} else {
		err = writeBytes(writer, []byte{byte(value), byte(value >> 8)})
	}

	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write ushort")
	}

	return nil
}

func writeInt32(writer io.ByteWriter, value int32, endian Endian) error {

	var err error

	if endian == BigEndian {
		err = writeBytes(writer, []byte{byte(value >> 24), byte(value >> 16), byte(value >> 8), byte(value)})
	} else {
		err = writeBytes(writer, []byte{byte(value), byte(value >> 8), byte(value >> 16), byte(value >> 24)})
	}

	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write int")
	}

	return nil
}

func writeUInt32(writer io.ByteWriter, value uint32, endian Endian) error {

	var err error

	if endian == BigEndian {
		err = writeBytes(writer, []byte{byte(value >> 24), byte(value >> 16), byte(value >> 8), byte(value)})
	} else {
		err = writeBytes(writer, []byte{byte(value), byte(value >> 8), byte(value >> 16), byte(value >> 24)})
	}

	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write uint")
	}

	return nil
}

func writeInt64(writer io.ByteWriter, value int64, endian Endian) error {

	var err error

	if endian == BigEndian {
		err = writeBytes(writer, []byte{byte(value >> 56), byte(value >> 48), byte(value >> 40), byte(value >> 32),
			byte(value >> 24), byte(value >> 16), byte(value >> 8), byte(value)})
	} else {
		err = writeBytes(writer, []byte{byte(value), byte(value >> 8), byte(value >> 16), byte(value >> 24),
			byte(value >> 32), byte(value >> 40), byte(value >> 48), byte(value >> 56)})
	}

	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write long")
	}

	return nil
}

func writeUInt64(writer io.ByteWriter, value uint64, endian Endian) error {

	var err error

	if endian == BigEndian {
		err = writeBytes(writer, []byte{byte(value >> 56), byte(value >> 48), byte(value >> 40), byte(value >> 32),
			byte(value >> 24), byte(value >> 16), byte(value >> 8), byte(value)})
	} else {
		err = writeBytes(writer, []byte{byte(value), byte(value >> 8), byte(value >> 16), byte(value >> 24),
			byte(value >> 32), byte(value >> 40), byte(value >> 48), byte(value >> 56)})
	}

	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write ulong")
	}

	return nil
}

func writeFloat32(writer io.ByteWriter, value float32, endian Endian) error {

	asUint := math.Float32bits(value)

	var err error
	if endian == BigEndian {
		err = writeBytes(writer, []byte{byte(asUint >> 24), byte(asUint >> 16), byte(asUint >> 8), byte(asUint)})

	} else {
		err = writeBytes(writer, []byte{byte(asUint), byte(asUint >> 8), byte(asUint >> 16), byte(asUint >> 24)})
	}

	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write float")
	}

	return nil
}

func writeFloat64(writer io.ByteWriter, value float64, endian Endian) error {

	asUint := math.Float64bits(value)

	var err error
	if endian == BigEndian {
		err = writeBytes(writer, []byte{byte(asUint >> 56), byte(asUint >> 48), byte(asUint >> 40), byte(asUint >> 32),
			byte(asUint >> 24), byte(asUint >> 16), byte(asUint >> 8), byte(asUint)})
	} else {
		err = writeBytes(writer, []byte{byte(asUint), byte(asUint >> 8), byte(asUint >> 16), byte(asUint >> 24),
			byte(asUint >> 32), byte(asUint >> 40), byte(asUint >> 48), byte(asUint >> 56)})
	}

	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write double")
	}

	return nil
}

// WriteString writes a string to the writer.
// First, the length of the string is written as a uint32.
// Then, the string is written as a sequence of bytes.
func writeString(writer io.ByteWriter, value string, endian Endian) error {

	if err := writeUInt32(writer, uint32(len(value)), endian); err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write string length")
	}

	if err := writeBytes(writer, []byte(value)); err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write string")
	}

	return nil
}
