package nbt

import (
	"fmt"
	"github.com/joomcode/errorx"
	"math"
	"regexp"
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
	if isString(value) {

		err := validateString(value)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "invalid string %s", value)
		}

		return StringTag(value[1 : len(value)-1]), nil
	}

	// Handle numbers
	if value[0] == '-' || (value[0] >= '0' && value[0] <= '9') {
		return parseNumber(value)
	}

	if strings.HasPrefix(value, "[B;") && strings.HasSuffix(value, "]") {
		return parseByteArray(value)
	}

	if strings.HasPrefix(value, "[I;") && strings.HasSuffix(value, "]") {
		return parseIntArray(value)
	}

	if strings.HasPrefix(value, "[L;") && strings.HasSuffix(value, "]") {
		return parseLongArray(value)
	}

	// Handle lists
	if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
		return parseList(value)
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
			return nil, errorx.IllegalState.Wrap(err, "failed to parse byte %d", value)
		}
		return ByteTag(value), nil
	case 's', 'S':
		value, err := strconv.ParseInt(value[0:len(value)-1], 10, 16)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to parse short '%d'", value)
		}
		return ShortTag(value), nil
	case 'l', 'L':
		value, err := strconv.ParseInt(value[0:len(value)-1], 10, 64)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to parse long '%d'", value)
		}
		return LongTag(value), nil
	case 'f', 'F':
		value, err := strconv.ParseFloat(value[0:len(value)-1], 32)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to parse float '%f'", value)
		}
		return FloatTag(value), nil
	case 'd', 'D':
		value, err := strconv.ParseFloat(value[0:len(value)-1], 64)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to parse double '%f'", value)
		}
		return DoubleTag(value), nil
	default:

		// Parse as double if there is a decimal point
		if strings.Contains(value, ".") {
			value, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to parse double '%f'", value)
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

func parseList(value string) (ListTag, error) {

	// Remove surrounding brackets
	value = value[1 : len(value)-1]

	// Split the string into a list of values
	values := strings.Split(value, ",")

	if len(values) == 1 && values[0] == "" {
		return nil, nil
	}

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
	return listEntries, nil
}

func parseCompound(value string) (*CompoundTag, error) {

	// Remove surrounding brackets
	value = value[1 : len(value)-1]

	pairs := parseCompoundPairs(value)

	// Create the compound
	compound := CompoundTag{}

	// Parse all pairs
	for _, pair := range pairs {

		// Split into key and value
		parts := strings.SplitN(pair, ":", 2)

		// Validate the key
		err := validateCompoundKey(parts[0])
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "invalid key %q", parts[0])
		}

		// Parse the value
		tag, err := parseValue(parts[1])
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to parse value %q", parts[1])
		}

		compound[parts[0]] = tag
	}

	return &compound, nil
}

func parseCompoundPairs(value string) []string {
	// Only split if not in {} or [] and not in a string
	var pairs []string
	squareBracketDepth := 0
	curlyBracketDepth := 0
	start := 0
	isInQuotes := false
	for i, c := range value {
		switch c {
		case '[':
			if !isInQuotes {
				squareBracketDepth++
			}
		case ']':
			if !isInQuotes {
				squareBracketDepth--
			}
		case '{':
			if !isInQuotes {
				curlyBracketDepth++
			}
		case '}':
			if !isInQuotes {
				curlyBracketDepth--
			}
		case '"':

			isEscaped := false

			// Check for escaped quotes
			if isInQuotes && len(value) >= 2 {

				backslashCount := 0

				// Count \ before the quote
				for j := i - 1; j >= 0; j-- {
					if value[j] == '\\' {
						backslashCount++
					} else {
						break
					}
				}

				if backslashCount%2 == 1 {
					isEscaped = true
				}
			}

			if !isEscaped {
				isInQuotes = !isInQuotes
			}

		case ',':
			if !isInQuotes && squareBracketDepth == 0 && curlyBracketDepth == 0 {
				pairs = append(pairs, value[start:i])
				start = i + 1
			}
		}
	}

	// Append the last pair
	pairs = append(pairs, value[start:])

	return pairs
}

func parseByteArray(input string) (ByteArrayTag, error) {

	// Assert it starts with [B;
	if !strings.HasPrefix(input, "[B;") {
		return nil, errorx.IllegalState.New("invalid byte array: %s", input)
	}

	// Remove surrounding brackets, [B; and ]
	input = input[3 : len(input)-1]

	// Split the string into a list of values
	values := strings.Split(input, ",")

	if len(values) == 1 && values[0] == "" {
		return nil, nil
	}

	// Create the byte array
	byteArray := make([]int8, len(values))

	// Parse all values
	for i, value := range values {

		if !strings.HasSuffix(value, "b") && !strings.HasSuffix(value, "B") {
			return nil, errorx.IllegalState.New("invalid byte array value: %s", input)
		}

		value = value[0 : len(value)-1]

		byteValue, err := strconv.ParseInt(value, 10, 8)
		if err != nil {
			panic(errorx.IllegalState.Wrap(err, "failed to parse byte %q", value))
		}

		byteArray[i] = int8(byteValue)
	}

	return byteArray, nil
}

func parseIntArray(input string) (IntArrayTag, error) {

	// Assert it starts with [I;
	if !strings.HasPrefix(input, "[I;") {
		return nil, errorx.IllegalState.New("invalid int array: %s", input)
	}

	// Remove surrounding brackets, [I; and ]
	input = input[3 : len(input)-1]

	// Split the string into a list of values
	values := strings.Split(input, ",")

	if len(values) == 1 && values[0] == "" {
		return nil, nil
	}

	// Create the int array
	intArray := make([]int32, len(values))

	// Parse all values
	for i, value := range values {

		intValue, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			panic(errorx.IllegalState.Wrap(err, "failed to parse int %q", value))
		}

		intArray[i] = int32(intValue)
	}

	return intArray, nil
}

func parseLongArray(input string) (LongArrayTag, error) {

	// Assert it starts with [L;
	if !strings.HasPrefix(input, "[L;") {
		return nil, errorx.IllegalState.New("invalid long array: %s", input)
	}

	// Remove surrounding brackets, [L; and ]
	input = input[3 : len(input)-1]

	// Split the string into a list of values
	values := strings.Split(input, ",")

	if len(values) == 1 && values[0] == "" {
		return nil, nil
	}

	// Create the long array
	longArray := make([]int64, len(values))

	// Parse all values
	for i, value := range values {

		if !strings.HasSuffix(value, "l") && !strings.HasSuffix(value, "L") {
			return nil, errorx.IllegalState.New("invalid long array value: %s", input)
		}

		value = value[0 : len(value)-1]

		longValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			panic(errorx.IllegalState.Wrap(err, "failed to parse long %q", value))
		}

		longArray[i] = longValue
	}

	return longArray, nil
}

var stringValidationRegex = regexp.MustCompile(`[a-zA-Z0-9_\-.+]+`)

func validateString(input string) error {

	text := input[1 : len(input)-1]

	// Validate every " has a \ before it
	for i := 0; i < len(text); i++ {
		if text[i] == '"' && (i == 0 || text[i-1] != '\\') {
			return errorx.IllegalState.New("invalid string, no \\ before \": %s", input)
		}
	}

	if !stringValidationRegex.MatchString(text) {
		return errorx.IllegalState.New("invalid string: %s", input)
	}

	return nil
}

var compoundKeyValidationRegex = regexp.MustCompile(`[a-zA-Z0-9_\-.+]+`)

func validateCompoundKey(input string) error {

	if isString(input) {
		return nil
	}

	if !compoundKeyValidationRegex.MatchString(input) {
		return errorx.IllegalState.New("invalid compound key: %s", input)
	}

	return nil
}

func isString(input string) bool {
	return len(input) > 2 &&
		(input[0] == '"' && input[len(input)-1] == '"') ||
		(input[0] == '\'' && input[len(input)-1] == '\'')
}

func (n *NBT) FormatSNBT() (string, error) {
	return formatNBT(n)
}

func formatNBT(nbt *NBT) (string, error) {

	var builder strings.Builder

	builder.WriteString("{")

	index := 0
	length := len(nbt.Tags)

	for key, value := range nbt.Tags {

		builder.WriteString(key)
		builder.WriteString(":")
		formattedValue, err := formatTag(value)
		if err != nil {
			return "", errorx.IllegalState.Wrap(err, "failed to format value tag %s:%s", key, value)
		}
		builder.WriteString(formattedValue)

		if index < length-1 {
			builder.WriteString(",")
		}

		index++
	}

	builder.WriteString("}")

	return builder.String(), nil
}

func formatCompound(compound CompoundTag) (string, error) {

	var builder strings.Builder

	builder.WriteString("{")

	index := 0
	length := len(compound)

	for key, value := range compound {

		builder.WriteString(key)
		builder.WriteString(":")
		formattedValue, err := formatTag(value)
		if err != nil {
			return "", errorx.IllegalState.Wrap(err, "failed to format value tag %s:%s", key, value)
		}
		builder.WriteString(formattedValue)

		if index < length-1 {
			builder.WriteString(",")
		}

		index++
	}

	builder.WriteString("}")

	return builder.String(), nil
}

func formatTag(tag Tag) (string, error) {

	switch tag.ID() {
	case 1:
		return fmt.Sprintf("%db", tag.(ByteTag)), nil
	case 2:
		return fmt.Sprintf("%ds", int16(tag.(ShortTag))), nil
	case 3:
		return strconv.Itoa(int(tag.(IntTag))), nil
	case 4:
		return fmt.Sprintf("%dl", int64(tag.(LongTag))), nil
	case 5:
		return fmt.Sprintf("%vf", float32(tag.(FloatTag))), nil
	case 6:
		double := float64(tag.(DoubleTag))
		if math.Mod(double, 1.0) == 0 {
			return fmt.Sprintf("%.2f", double), nil
		}
		return fmt.Sprintf("%v", double), nil
	case 7:
		return formatByteArray(tag.(ByteArrayTag)), nil
	case 8:
		return "\"" + string(tag.(StringTag)) + "\"", nil
	case 9:
		return formatList(tag.(ListTag))
	case 10:
		return formatCompound(*tag.(*CompoundTag))
	case 11:
		return formatIntArray(tag.(IntArrayTag)), nil
	case 12:
		return formatLongArray(tag.(LongArrayTag)), nil
	default:
		return "", errorx.IllegalState.New("invalid tag id: %d", tag.ID())
	}
}

func formatByteArray(array ByteArrayTag) string {

	var builder strings.Builder

	builder.WriteString("[B;")

	for i, value := range array {

		builder.WriteString(strconv.Itoa(int(value)) + "b")

		if i < len(array)-1 {
			builder.WriteString(",")
		}
	}

	builder.WriteString("]")

	return builder.String()
}

func formatList(list ListTag) (string, error) {

	var builder strings.Builder

	builder.WriteString("[")

	typeID := byte(0)
	if len(list) > 0 {
		typeID = list[0].ID()
	}

	for i, value := range list {

		if value.ID() != typeID {
			return "", errorx.IllegalState.New("invalid list value type: %d, expected %d", value, typeID)
		}

		formattedTag, err := formatTag(value)
		if err != nil {
			return "", err
		}
		builder.WriteString(formattedTag)

		if i < len(list)-1 {
			builder.WriteString(",")
		}
	}

	builder.WriteString("]")

	return builder.String(), nil
}

func formatIntArray(array IntArrayTag) string {

	var builder strings.Builder

	builder.WriteString("[I;")

	for i, value := range array {

		builder.WriteString(strconv.Itoa(int(value)))

		if i < len(array)-1 {
			builder.WriteString(",")
		}
	}

	builder.WriteString("]")

	return builder.String()
}

func formatLongArray(array LongArrayTag) string {

	var builder strings.Builder

	builder.WriteString("[L;")

	for i, value := range array {

		builder.WriteString(strconv.FormatInt(value, 10) + "l")

		if i < len(array)-1 {
			builder.WriteString(",")
		}
	}

	builder.WriteString("]")

	return builder.String()
}
