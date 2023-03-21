package nbt

import (
	"github.com/joomcode/errorx"
	"strconv"
	"strings"
)

// https://minecraft.fandom.com/wiki/NBT_format#SNBT_format

func ParseSNBT(input string) (Tag, error) {
	return parseValue(input)
}

func parseValue(value string) (Tag, error) {

	// Handle basic types
	switch value {
	case "true":
		return ByteTag(1), nil
	case "false":
		return ByteTag(0), nil
	}

	// Handle Strings
	if value[0] == '"' && value[len(value)-1] == '"' {
		return StringTag(value[1 : len(value)-1]), nil
	}

	// Handle numbers
	if value[0] == '-' || (value[0] >= '0' && value[0] <= '9') {
		return parseNumber(value)
	}

	// Handle lists and arrays
	if value[0] == '[' && value[len(value)-1] == ']' {
		return parseListOrArray(value)
	}

	// Handle compound
	if value[0] == '{' && value[len(value)-1] == '}' {
		return parseCompound(value)
	}

	return nil, errorx.IllegalState.New("invalid value %s", value)
}

func parseNumber(value string) (Tag, error) {
	switch value[len(value)-1] {
	case 'b', 'B':
		value, err := strconv.ParseInt(value[0:len(value)-1], 10, 8)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to parse byte %q", value)
		}
		return ByteTag(value), nil
	case 's', 'S':
		value, err := strconv.ParseInt(value[0:len(value)-1], 10, 16)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to parse short")
		}
		return ShortTag(value), nil
	case 'l', 'L':
		value, err := strconv.ParseInt(value[0:len(value)-1], 10, 64)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to parse long %q", value)
		}
		return LongTag(value), nil
	case 'f', 'F':
		value, err := strconv.ParseFloat(value[0:len(value)-1], 32)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to parse float %q", value)
		}
		return FloatTag(value), nil
	case 'd', 'D':
		value, err := strconv.ParseFloat(value[0:len(value)-1], 64)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to parse double %q", value)
		}
		return DoubleTag(value), nil
	default:

		// Parse as double if there is a decimal point
		if strings.Contains(value, ".") {
			value, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to parse double %q", value)
			}
			return DoubleTag(value), nil
		}

		// Otherwise parse as int
		value, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to parse int %q", value)
		}
		return IntTag(value), nil
	}
}

func parseListOrArray(value string) (Tag, error) {

	// Remove surrounding brackets
	value = value[1 : len(value)-1]

	// Split the string into a list of values
	values := strings.Split(value, ",")

	listEntries := make([]Tag, len(values))

	var typeID *uint8

	// Parse all values
	for i, value := range values {

		tag, err := parseValue(value)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to parse value %q", value)
		}

		listEntries[i] = tag

		if typeID == nil {
			id := tag.ID()
			typeID = &id
		} else if *typeID != tag.ID() {
			return nil, errorx.IllegalState.New("invalid type %q, expected %q", tag.ID(), *typeID)
		}
	}

	// Default to END type id if the list is empty
	if typeID == nil {
		id := uint8(0)
		typeID = &id
	}

	// Return the list
	return ListTag(listEntries), nil
}

func parseCompound(value string) (NBT, error) {

	// Remove surrounding brackets
	value = value[1 : len(value)-1]

	// Split into key:value pairs
	pairs := strings.Split(value, ",")

	// Create the compound
	compound := NBT{}

}

func (n *NBT) FormatSNBT() string {

}
