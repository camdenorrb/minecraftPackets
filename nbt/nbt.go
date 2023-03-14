package nbt

import (
	"bytes"
	"encoding/binary"
	"github.com/joomcode/errorx"
	"io"
	"math"
)

// region Endian

type Endian int

const (
	BigEndian Endian = iota
	LittleEndian
)

// endregion Endian

type Tag interface {
	ID() uint8
	Size() int                                              // Total size in bytes including the tag ID
	PushToWriter(writer io.ByteWriter, endian Endian) error // Pushes the tag to the writer including the tag ID
}

// NBT We use NBT to represent Compound
type NBT struct {
	Name string
	Tags []Tag
}

// region Compound

func (n *NBT) ID() uint8 {
	return 10
}

func (n *NBT) Size() int {

	size := 0
	size += 2 // Name size
	size += 1 // Tag ID size

	for _, tag := range n.Tags {
		size += tag.Size()
	}

	return size
}

func (n *NBT) PushToWriter(writer io.ByteWriter, endian Endian) error {

	if err := String(n.Name).PushToWriter(writer, endian); err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write compound name")
	}

	for _, tag := range n.Tags {
		if err := writer.WriteByte(tag.ID()); err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write tag ID")
		}
		if err := tag.PushToWriter(writer, endian); err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write tag")
		}
	}

	if err := writer.WriteByte(0); err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write end tag")
	}

	return nil
}

// endregion Compound

func readNBT(reader bytes.Reader, endian Endian) (*NBT, error) {
	nbt := NBT{}
	for {
		tagID, err := reader.ReadByte()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read tag ID")
		}
		if tagID == 0 {
			break
		}
		tag, err := readTag(reader, tagID, endian)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read tag")
		}
		nbt.Tags = append(nbt.Tags, tag)
	}
	return &nbt, nil
}

func readTag(reader bytes.Reader, tagID uint8, endian Endian) (Tag, error) {
	switch tagID {
	case 0:
		return End{}, nil
	case 1:
		return readByte(reader)
	case 2:
		return readShort(reader, endian)
	case 3:
		return readInt(reader, endian)
	case 4:
		return readLong(reader, endian)
	case 5:
		return readFloat(reader, endian)
	case 6:
		return readDouble(reader, endian)
	case 7:
		return readByteArray(reader, endian)
	case 8:
		return readString(reader, endian)
	case 9:
		return readList(reader, endian)
	case 10:
		return readNBT(reader, endian)
	case 11:
		return readIntArray(reader, endian)
	case 12:
		return readLongArray(reader, endian)
	default:
		return nil, errorx.IllegalArgument.New("unknown tag id: " + string(tagID))
	}
}

// region End

type End struct{}

func (e End) ID() uint8 {
	return 0
}

func (e End) Size() int {
	// Data bytes + tag ID size
	return 0 + 1
}

func (e End) PushToWriter(writer io.ByteWriter, endian Endian) error {

	err := writer.WriteByte(e.ID())
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write end tag ID")
	}

	return nil // NOOP
}

// endregion End

// region Byte

type Byte int8

func (b Byte) ID() uint8 {
	return 1
}

func (b Byte) Size() int {
	// Data bytes + tag ID size
	return 1 + 1
}

func (b Byte) PushToWriter(writer io.ByteWriter, endian Endian) error {

	err := writer.WriteByte(byte(b))
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write byte tag id")
	}

	return nil
}

func readByte(reader bytes.Reader) (Byte, error) {
	b, err := reader.ReadByte()
	if err != nil {
		return 0, errorx.IllegalState.Wrap(err, "failed to read byte")
	}
	return Byte(b), nil
}

// endregion Byte

// region Short

type Short int16

func (s Short) ID() uint8 {
	return 2
}

func (s Short) Size() int {
	// Data bytes + tag ID size
	return 2 + 1
}

func (s Short) PushToWriter(writer io.ByteWriter, endian Endian) error {

	err := writer.WriteByte(s.ID())
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write short tag id")
	}

	if endian == BigEndian {
		err = writeBytes(writer, []byte{byte(s >> 8), byte(s)})
	} else {
		err = writeBytes(writer, []byte{byte(s), byte(s >> 8)})
	}

	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write short")
	}

	return nil
}

func readShort(reader bytes.Reader, endian Endian) (Short, error) {

	data, err := readNBytes(reader, 2)
	if err != nil {
		return 0, errorx.IllegalState.Wrap(err, "failed to read short")
	}

	if endian == BigEndian {
		return Short(binary.BigEndian.Uint16(data)), nil
	}
	return Short(binary.LittleEndian.Uint16(data)), nil
}

// endregion Short

// region Int

type Int int32

func (i Int) ID() uint8 {
	return 3
}

func (i Int) Size() int {
	// Data bytes + tag ID size
	return 4 + 1
}

func (i Int) PushToWriter(writer io.ByteWriter, endian Endian) error {

	err := writer.WriteByte(i.ID())
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write int tag id")
	}

	if endian == BigEndian {
		err = writeBytes(writer, []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)})
	} else {
		err = writeBytes(writer, []byte{byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24)})
	}

	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write int")
	}

	return nil
}

func readInt(reader bytes.Reader, endian Endian) (Int, error) {

	data, err := readNBytes(reader, 4)
	if err != nil {
		return 0, errorx.IllegalState.Wrap(err, "failed to read int")
	}

	if endian == BigEndian {
		return Int(binary.BigEndian.Uint32(data)), nil
	}
	return Int(binary.LittleEndian.Uint32(data)), nil
}

// endregion Int

// region Long

type Long int64

func (l Long) ID() uint8 {
	return 4
}

func (l Long) Size() int {
	// Data bytes + tag ID size
	return 8 + 1
}

func (l Long) PushToWriter(writer io.ByteWriter, endian Endian) error {

	err := writer.WriteByte(l.ID())
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write long tag id")
	}

	if endian == BigEndian {
		err = writeBytes(writer, []byte{byte(l >> 56), byte(l >> 48), byte(l >> 40), byte(l >> 32), byte(l >> 24), byte(l >> 16), byte(l >> 8), byte(l)})
	} else {
		err = writeBytes(writer, []byte{byte(l), byte(l >> 8), byte(l >> 16), byte(l >> 24), byte(l >> 32), byte(l >> 40), byte(l >> 48), byte(l >> 56)})
	}

	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write long")
	}

	return nil
}

func readLong(reader bytes.Reader, endian Endian) (Long, error) {

	data, err := readNBytes(reader, 8)
	if err != nil {
		return 0, errorx.IllegalState.Wrap(err, "failed to read long")
	}

	if endian == BigEndian {
		return Long(binary.BigEndian.Uint64(data)), nil
	}

	return Long(binary.LittleEndian.Uint64(data)), nil
}

// endregion Long

// region Float

type Float float32

func (f Float) ID() uint8 {
	return 5
}

func (f Float) Size() int {
	// Data bytes + tag ID size
	return 4 + 1
}

func (f Float) PushToWriter(writer io.ByteWriter, endian Endian) error {

	err := writer.WriteByte(f.ID())
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write float tag id")
	}

	asUint := math.Float32bits(float32(f))
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

func readFloat(reader bytes.Reader, endian Endian) (Float, error) {

	data, err := readNBytes(reader, 4)
	if err != nil {
		return 0, errorx.IllegalState.Wrap(err, "failed to read float")
	}

	if endian == BigEndian {
		return Float(math.Float32frombits(binary.BigEndian.Uint32(data))), nil
	}

	return Float(math.Float32frombits(binary.LittleEndian.Uint32(data))), nil
}

// endregion Float

// region Double

type Double float64

func (d Double) ID() uint8 {
	return 6
}

func (d Double) Size() int {
	return 8
}

func (d Double) PushToWriter(writer io.ByteWriter, endian Endian) error {

	err := writer.WriteByte(d.ID())
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write double tag id")
	}

	asUint := math.Float64bits(float64(d))
	if endian == BigEndian {
		err = writeBytes(writer, []byte{byte(asUint >> 56), byte(asUint >> 48), byte(asUint >> 40), byte(asUint >> 32), byte(asUint >> 24), byte(asUint >> 16), byte(asUint >> 8), byte(asUint)})
	} else {
		err = writeBytes(writer, []byte{byte(asUint), byte(asUint >> 8), byte(asUint >> 16), byte(asUint >> 24), byte(asUint >> 32), byte(asUint >> 40), byte(asUint >> 48), byte(asUint >> 56)})
	}

	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write double")
	}

	return nil
}

func readDouble(reader bytes.Reader, endian Endian) (Double, error) {

	data, err := readNBytes(reader, 8)
	if err != nil {
		return 0, errorx.IllegalState.Wrap(err, "failed to read double")
	}

	if endian == BigEndian {
		return Double(math.Float64frombits(binary.BigEndian.Uint64(data))), nil
	}

	return Double(math.Float64frombits(binary.LittleEndian.Uint64(data))), nil
}

// endregion Double

// region ByteArray

type ByteArray []byte

func (b ByteArray) ID() uint8 {
	return 7
}

func (b ByteArray) Size() int {
	// Data bytes + tag ID size + name size (short)
	return len(b) + 1 + 2
}

func (b ByteArray) PushToWriter(writer io.ByteWriter, endian Endian) error {

	err := writer.WriteByte(b.ID())
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write byte array tag id")
	}

	err = Int(len(b)).PushToWriter(writer, endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write byte array length")
	}

	err = writeBytes(writer, b)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write byte array")
	}

	return nil
}

func readByteArray(reader bytes.Reader, endian Endian) (ByteArray, error) {

	length, err := readInt(reader, endian)
	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "failed to read byte array length")
	}

	data, err := readNBytes(reader, int(length))
	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "failed to read byte array")
	}

	return data, nil
}

// endregion ByteArray

// region String

type String string

func (s String) ID() uint8 {
	return 8
}

func (s String) Size() int {
	// Data bytes + tag ID size + length (short)
	return len(s) + 1 + 2
}

func (s String) PushToWriter(writer io.ByteWriter, endian Endian) error {

	err := writer.WriteByte(s.ID())
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write string tag id")
	}

	err = Short(len(s)).PushToWriter(writer, endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write string length")
	}

	err = writeBytes(writer, []byte(s))
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write string")
	}

	return nil
}

func readString(reader bytes.Reader, endian Endian) (String, error) {

	length, err := readShort(reader, endian)
	if err != nil {
		return "", errorx.IllegalState.Wrap(err, "failed to read string length")
	}

	data, err := readNBytes(reader, int(length))
	if err != nil {
		return "", errorx.IllegalState.Wrap(err, "failed to read string")
	}

	return String(data), nil
}

// endregion String

// region List

type List []Tag

func (l List) ID() uint8 {
	return 9
}

func (l List) Size() int {

	size := 0

	for _, tag := range l {
		size += tag.Size()
	}

	// Entry Tags size + Tag ID size + length (int)
	return size + 1 + 4
}

func (l List) PushToWriter(writer io.ByteWriter, endian Endian) error {

	err := writer.WriteByte(l.ID())
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write list tag id")
	}

	err = Byte(l[0].ID()).PushToWriter(writer, endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write tag id for list type")
	}

	err = Int(len(l)).PushToWriter(writer, endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write length of list")
	}

	for _, tag := range l {
		err := tag.PushToWriter(writer, endian)
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write list")
		}
	}

	return nil
}

func readList(reader bytes.Reader, endian Endian) (List, error) {

	var list List

	tagID, err := reader.ReadByte()
	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "failed to tag id for list type")
	}

	length, err := readInt(reader, endian)
	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "failed to length of list")
	}

	for i := 0; i < int(length); i++ {
		tag, err := readTag(reader, tagID, endian)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read list")
		}
		list = append(list, tag)
	}

	return list, nil

}

// endregion List

// region IntArray

type IntArray []int32

func (i IntArray) ID() uint8 {
	return 11
}

func (i IntArray) Size() int {
	// Data bytes + tag ID size + length (int)
	return len(i)*4 + 1 + 4
}

func (i IntArray) PushToWriter(writer io.ByteWriter, endian Endian) error {

	err := writer.WriteByte(i.ID())
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write int array tag id")
	}

	err = Int(len(i)).PushToWriter(writer, endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to length of int array length")
	}

	for _, v := range i {
		err := Int(v).PushToWriter(writer, endian)
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write int array")
		}
	}

	return nil
}

func readIntArray(reader bytes.Reader, endian Endian) (IntArray, error) {

	length, err := readInt(reader, endian)
	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "failed to read length of int array")
	}

	var array IntArray
	for i := 0; i < int(length); i++ {
		v, err := readInt(reader, endian)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read int array")
		}
		array = append(array, int32(v))
	}

	return array, nil
}

// endregion IntArray

// region LongArray

type LongArray []int64

func (l LongArray) ID() uint8 {
	return 12
}

func (l LongArray) Size() int {
	// Data bytes + tag ID size + length (int)
	return len(l)*8 + 1 + 4
}

func (l LongArray) PushToWriter(writer io.ByteWriter, endian Endian) error {

	err := writer.WriteByte(l.ID())
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write long array tag id")
	}

	err = Int(len(l)).PushToWriter(writer, endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to length of long array")
	}

	for _, v := range l {
		err := Long(v).PushToWriter(writer, endian)
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write long array")
		}
	}

	return nil
}

func readLongArray(reader bytes.Reader, endian Endian) (LongArray, error) {

	length, err := readInt(reader, endian)
	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "failed to read length of long array")
	}

	var array LongArray
	for i := 0; i < int(length); i++ {
		v, err := readLong(reader, endian)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read long array")
		}
		array = append(array, int64(v))
	}

	return array, nil
}

// endregion LongArray

func writeBytes(writer io.ByteWriter, data []byte) error {
	for _, b := range data {
		err := writer.WriteByte(b)
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write bytes")
		}
	}
	return nil
}

func readNBytes(buffer bytes.Reader, n int) ([]byte, error) {

	data := make([]byte, n)

	amount, err := buffer.Read(data)
	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "Failed to read %d bytes", n)
	}
	if amount != n {
		return nil, errorx.IllegalState.New("Failed to read %d bytes, only read %d", n, amount)
	}

	return data, nil
}
