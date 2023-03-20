package nbt

import (
	"bytes"
	"github.com/joomcode/errorx"
	"io"
	"strconv"
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
	Size() int // Total size in bytes including the tag ID
	//TODO: String() string                                                            // Returns a string representation of the tag
	PushToWriter(writer io.ByteWriter, endian Endian, includeTagID bool) error // Pushes the tag to the writer
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

func (n *NBT) PushToWriter(writer io.ByteWriter, endian Endian, includeTagID bool) error {

	if includeTagID {
		if err := writer.WriteByte(n.ID()); err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write compound tag ID")
		}
	}

	err := writeUShort(writer, uint16(len(n.Name)), endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write compound name length")
	}

	err = writeBytes(writer, []byte(n.Name))
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write compound name")
	}

	for _, tag := range n.Tags {
		if err := tag.PushToWriter(writer, endian, true); err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write tag")
		}
	}

	if err := writer.WriteByte(0); err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write end tag")
	}

	return nil
}

func readNBT(reader *bytes.Reader, endian Endian) (*NBT, error) {

	tagID, err := reader.ReadByte()
	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "failed to read compound tag ID")
	}
	if tagID != 10 {
		return nil, errorx.IllegalState.New("expected compound tag")
	}

	return readCompound(reader, endian)
}

func readCompound(reader *bytes.Reader, endian Endian) (*NBT, error) {

	nameLength, err := readUShort(reader, endian)
	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "failed to read compound name length")
	}

	nameBytes, err := readNBytes(reader, int(nameLength))
	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "failed to read compound name")
	}

	nbt := NBT{Name: string(nameBytes)}

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

// endregion Compound

func readTag(reader *bytes.Reader, tagID uint8, endian Endian) (Tag, error) {
	switch tagID {
	case 0:
		return EndTag{}, nil
	case 1:
		byteValue, err := reader.ReadByte()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read byte tag")
		}
		return ByteTag(byteValue), nil
	case 2:
		shortValue, err := readShort(reader, endian)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read short tag")
		}
		return ShortTag(shortValue), nil
	case 3:
		intValue, err := readInt(reader, endian)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read int tag")
		}
		return IntTag(intValue), nil
	case 4:
		longValue, err := readLong(reader, endian)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read long tag")
		}
		return LongTag(longValue), nil
	case 5:
		floatValue, err := readFloat(reader, endian)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read float tag")
		}
		return FloatTag(floatValue), nil
	case 6:
		doubleValue, err := readDouble(reader, endian)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read double tag")
		}
		return DoubleTag(doubleValue), nil
	case 7:

		length, err := readInt(reader, endian)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read byte array length")
		}

		bytesArray, err := readNBytes(reader, int(length))
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read byte array tag")
		}

		return ByteArrayTag(bytesArray), nil

	case 8:

		length, err := readShort(reader, endian)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read string length")
		}

		stringArray, err := readNBytes(reader, int(length))
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read string tag")
		}

		return StringTag(stringArray), nil
	case 9:
		return readListTag(reader, endian)
	case 10:
		return readCompound(reader, endian)
	case 11:
		return readIntArrayTag(reader, endian)
	case 12:
		return readLongArrayTag(reader, endian)
	default:
		return nil, errorx.IllegalArgument.New("unknown tag id: " + strconv.FormatUint(uint64(tagID), 10))
	}
}

// region EndTag

type EndTag struct{}

func (e EndTag) ID() uint8 {
	return 0
}

func (e EndTag) Size() int {
	// Data bytes + tag ID size
	return 0 + 1
}

func (e EndTag) PushToWriter(writer io.ByteWriter, endian Endian, includeTagID bool) error {

	if includeTagID {
		err := writer.WriteByte(e.ID())
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write end tag ID")
		}
	}

	return nil // NOOP
}

// endregion EndTag

// region ByteTag

type ByteTag int8

func (b ByteTag) ID() uint8 {
	return 1
}

func (b ByteTag) Size() int {
	// Data bytes + tag ID size
	return 1 + 1
}

func (b ByteTag) PushToWriter(writer io.ByteWriter, endian Endian, includeTagID bool) error {

	if includeTagID {
		err := writer.WriteByte(b.ID())
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write byte tag id")
		}
	}

	err := writer.WriteByte(byte(b))
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write byte tag id")
	}

	return nil
}

// endregion ByteTag

// region ShortTag

type ShortTag int16

func (s ShortTag) ID() uint8 {
	return 2
}

func (s ShortTag) Size() int {
	// Data bytes + tag ID size
	return 2 + 1
}

func (s ShortTag) PushToWriter(writer io.ByteWriter, endian Endian, includeTagID bool) error {

	if includeTagID {
		err := writer.WriteByte(s.ID())
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write short tag id")
		}
	}

	err := writeShort(writer, int16(s), endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write short")
	}

	return nil
}

// endregion ShortTag

// region IntTag

type IntTag int32

func (i IntTag) ID() uint8 {
	return 3
}

func (i IntTag) Size() int {
	// Data bytes + tag ID size
	return 4 + 1
}

func (i IntTag) PushToWriter(writer io.ByteWriter, endian Endian, includeTagID bool) error {

	if includeTagID {
		err := writer.WriteByte(i.ID())
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write int tag id")
		}
	}

	err := writeInt(writer, int32(i), endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write int")
	}

	return nil
}

// endregion IntTag

// region LongTag

type LongTag int64

func (l LongTag) ID() uint8 {
	return 4
}

func (l LongTag) Size() int {
	// Data bytes + tag ID size
	return 8 + 1
}

func (l LongTag) PushToWriter(writer io.ByteWriter, endian Endian, includeTagID bool) error {

	if includeTagID {
		err := writer.WriteByte(l.ID())
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write long tag id")
		}
	}

	err := writeLong(writer, int64(l), endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write long")
	}

	return nil
}

// endregion LongTag

// region FloatTag

type FloatTag float32

func (f FloatTag) ID() uint8 {
	return 5
}

func (f FloatTag) Size() int {
	// Data bytes + tag ID size
	return 4 + 1
}

func (f FloatTag) PushToWriter(writer io.ByteWriter, endian Endian, includeTagID bool) error {

	if includeTagID {
		err := writer.WriteByte(f.ID())
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write float tag id")
		}
	}

	err := writeFloat(writer, float32(f), endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write float")
	}

	return nil
}

// endregion FloatTag

// region DoubleTag

type DoubleTag float64

func (d DoubleTag) ID() uint8 {
	return 6
}

func (d DoubleTag) Size() int {
	return 8
}

func (d DoubleTag) PushToWriter(writer io.ByteWriter, endian Endian, includeTagID bool) error {

	if includeTagID {
		err := writer.WriteByte(d.ID())
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write double tag id")
		}
	}

	err := writeDouble(writer, float64(d), endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write double")
	}

	return nil
}

// endregion DoubleTag

// region ByteArrayTag

type ByteArrayTag []byte

func (b ByteArrayTag) ID() uint8 {
	return 7
}

func (b ByteArrayTag) Size() int {
	// Data bytes + tag ID size + name size (short)
	return len(b) + 1 + 2
}

func (b ByteArrayTag) PushToWriter(writer io.ByteWriter, endian Endian, includeTagID bool) error {

	if includeTagID {
		err := writer.WriteByte(b.ID())
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write byte array tag id")
		}
	}

	err := writeInt(writer, int32(len(b)), endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write byte array length")
	}

	err = writeBytes(writer, b)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write byte array")
	}

	return nil
}

// endregion ByteArrayTag

// region StringTag

type StringTag string

func (s StringTag) ID() uint8 {
	return 8
}

func (s StringTag) Size() int {
	// Data bytes + tag ID size + length (short)
	return len(s) + 1 + 2
}

func (s StringTag) PushToWriter(writer io.ByteWriter, endian Endian, includeTagID bool) error {

	if includeTagID {
		err := writer.WriteByte(s.ID())
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write string tag id")
		}
	}

	asBytes := []byte(s)

	err := writeUShort(writer, uint16(len(asBytes)), endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write string length")
	}

	err = writeBytes(writer, asBytes)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write string")
	}

	return nil
}

// endregion StringTag

// region ListTag

type ListTag []Tag

func (l ListTag) ID() uint8 {
	return 9
}

func (l ListTag) Size() int {

	size := 0

	for _, tag := range l {
		size += tag.Size()
	}

	// Entry Tags size + Tag ID size + length (int)
	return size + 1 + 4
}

func (l ListTag) PushToWriter(writer io.ByteWriter, endian Endian, includeTagID bool) error {

	if includeTagID {
		err := writer.WriteByte(l.ID())
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write list tag id")
		}
	}

	typeID := 0
	if len(l) > 0 {
		typeID = int(l[0].ID())
	}

	err := writer.WriteByte(byte(typeID))
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write tag id for list type")
	}

	err = writeInt(writer, int32(len(l)), endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to write length of list")
	}

	for _, tag := range l {
		if tag.ID() != byte(typeID) {
			return errorx.IllegalState.New("list tag type mismatch, expected %d, got %d", typeID, tag.ID())
		}
		err := tag.PushToWriter(writer, endian, false)
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write list")
		}
	}

	return nil
}

func readListTag(reader *bytes.Reader, endian Endian) (ListTag, error) {

	var list ListTag

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

// endregion ListTag

// region IntArrayTag

type IntArrayTag []int32

func (i IntArrayTag) ID() uint8 {
	return 11
}

func (i IntArrayTag) Size() int {
	// Data bytes + tag ID size + length (int)
	return len(i)*4 + 1 + 4
}

func (i IntArrayTag) PushToWriter(writer io.ByteWriter, endian Endian, includeTagID bool) error {

	if includeTagID {
		err := writer.WriteByte(i.ID())
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write int array tag id")
		}
	}

	err := writeInt(writer, int32(len(i)), endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to length of int array length")
	}

	for _, v := range i {
		err := writeInt(writer, v, endian)
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write int array")
		}
	}

	return nil
}

func readIntArrayTag(reader *bytes.Reader, endian Endian) (IntArrayTag, error) {

	length, err := readInt(reader, endian)
	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "failed to read length of int array")
	}

	var array IntArrayTag
	for i := 0; i < int(length); i++ {
		v, err := readInt(reader, endian)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read int array")
		}
		array = append(array, v)
	}

	return array, nil
}

// endregion IntArrayTag

// region LongArrayTag

type LongArrayTag []int64

func (l LongArrayTag) ID() uint8 {
	return 12
}

func (l LongArrayTag) Size() int {
	// Data bytes + tag ID size + length (int)
	return len(l)*8 + 1 + 4
}

func (l LongArrayTag) PushToWriter(writer io.ByteWriter, endian Endian, includeTagID bool) error {

	if includeTagID {
		err := writer.WriteByte(l.ID())
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write long array tag id")
		}
	}

	err := writeInt(writer, int32(len(l)), endian)
	if err != nil {
		return errorx.IllegalState.Wrap(err, "failed to length of long array")
	}

	for _, v := range l {
		err := writeLong(writer, v, endian)
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to write long array")
		}
	}

	return nil
}

func readLongArrayTag(reader *bytes.Reader, endian Endian) (LongArrayTag, error) {

	length, err := readInt(reader, endian)
	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "failed to read length of long array")
	}

	var array LongArrayTag
	for i := 0; i < int(length); i++ {
		v, err := readLong(reader, endian)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to read long array")
		}
		array = append(array, v)
	}

	return array, nil
}

// endregion LongArrayTag
