package palette

import (
	"encoding/json"
	"fmt"
	"github.com/joomcode/errorx"
	"github.com/mitchellh/hashstructure/v2"
	"reflect"
	"strconv"
	"strings"
)

func (p *BlockStateProperties) Hash() (*uint64, error) {

	hash, err := hashstructure.Hash(p, hashstructure.FormatV2, nil)
	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "failed to hash block properties")
	}

	return &hash, nil
}

func (p *BlockPaletteJSON) UnmarshalJSON(data []byte) error {

	// So the issue is that the key is a Material, we need to convert string to it

	var rawMap map[string]BlockPaletteData
	if err := json.Unmarshal(data, &rawMap); err != nil {
		return errorx.IllegalState.Wrap(err, "failed to unmarshal raw map")
	}

	var outputMap = make(BlockPaletteJSON, len(rawMap))

	for key, value := range rawMap {

		formattedKey := strings.ToUpper(strings.TrimPrefix(key, "minecraft:"))

		material, exists := Material_value[formattedKey]
		if !exists {
			return errorx.IllegalState.New("failed to find material for %q", key)
		}
		outputMap[Material(material)] = value
	}

	*p = outputMap

	return nil
}

func (p BlockPaletteJSON) MarshalJSON() ([]byte, error) {

	var rawMap = make(map[string]BlockPaletteData, len(p))

	for key, value := range p {

		materialName, exists := Material_name[int32(key)]
		if !exists {
			return nil, errorx.IllegalState.New("failed to find material name for %q", key)
		}

		formattedMaterialName := fmt.Sprintf("minecraft:%s", strings.ToLower(materialName))
		rawMap[formattedMaterialName] = value
	}

	return json.Marshal(rawMap)
}

func (s BlockPaletteState) MarshalJSON() ([]byte, error) {

	// Makes id not omitEmpty
	correctRepresentation := struct {
		Default    *bool                 `json:"default,omitempty"`
		Id         uint32                `json:"id"`
		Properties *BlockStateProperties `json:"properties,omitempty"`
	}{
		Default:    s.Default,
		Id:         s.Id,
		Properties: s.Properties,
	}

	return json.Marshal(correctRepresentation)
}

func (p *BlockProperties) UnmarshalJSON(data []byte) error {

	var rawMap map[string][]json.RawMessage
	if err := json.Unmarshal(data, &rawMap); err != nil {
		return errorx.IllegalState.Wrap(err, "failed to unmarshal raw map")
	}

	v := reflect.ValueOf(p).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {

		field := t.Field(i)
		jsonKey := field.Tag.Get("json")

		commaIndex := strings.Index(jsonKey, ",")
		if commaIndex != -1 {
			jsonKey = jsonKey[:commaIndex]
		}

		rawValues, exists := rawMap[jsonKey]
		if !exists || !v.Field(i).CanSet() {
			continue
		}

		fieldValue := v.Field(i)

		switch fieldValue.Type().Elem().Kind() {
		case reflect.Uint32:

			var values []uint32
			for _, rawValue := range rawValues {

				intValue, err := strconv.ParseUint(strings.Trim(string(rawValue), "\""), 10, 32)
				if err != nil {
					return errorx.IllegalState.Wrap(err, "failed to parse uint32 field %v", field.Name)
				}

				values = append(values, uint32(intValue))
			}

			fieldValue.Set(reflect.ValueOf(values))

		case reflect.Bool:

			var values []bool
			for _, rawValue := range rawValues {

				boolValue, err := strconv.ParseBool(strings.Trim(string(rawValue), "\""))
				if err != nil {
					return errorx.IllegalState.Wrap(err, "failed to parse bool field %v", field.Name)
				}

				values = append(values, boolValue)
			}

			fieldValue.Set(reflect.ValueOf(values))

		case reflect.Int32:

			wasEnum := false

			var values []int32
			for _, rawValue := range rawValues {

				// Try to parse enum
				foundValue, err := getEnumIntValue(jsonKey, rawValue)
				if err != nil {
					return errorx.IllegalState.Wrap(err, "failed to parse enum field %v", field.Name)
				}
				if foundValue != -1 {
					values = append(values, foundValue)
					wasEnum = true
					continue
				}

				value, err := strconv.ParseInt(strings.Trim(string(rawValue), "\""), 10, 32)
				if err != nil {
					return errorx.IllegalState.Wrap(err, "failed to parse int32 field %v", field.Name)
				}

				values = append(values, int32(value))
			}

			if wasEnum {
				marshal, err := json.Marshal(values)
				if err != nil {
					return errorx.IllegalState.Wrap(err, "failed to marshal enum field %v", field.Name)
				}

				if err := json.Unmarshal(marshal, fieldValue.Addr().Interface()); err != nil {
					return errorx.IllegalState.Wrap(err, "failed to unmarshal enum field %v", field.Name)
				}
			} else {
				fieldValue.Set(reflect.ValueOf(values))
			}

		default:

			marshal, err := json.Marshal(rawValues)
			if err != nil {
				return errorx.IllegalState.Wrap(err, "failed to marshal field %v", field.Name)
			}

			if err := json.Unmarshal(marshal, fieldValue.Addr().Interface()); err != nil {
				return errorx.IllegalState.Wrap(err, "failed to unmarshal field %v", field.Name)
			}
		}
	}

	return nil
}

func (p BlockProperties) MarshalJSON() ([]byte, error) {

	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	rawMap := make(map[string][]string)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonKey := field.Tag.Get("json")
		fieldValue := v.Field(i)

		omitEmpty := strings.Contains(jsonKey, ",omitempty")

		if omitEmpty && fieldValue.IsZero() && jsonKey != "id" {
			continue
		}

		commaIndex := strings.Index(jsonKey, ",")
		if commaIndex != -1 {
			jsonKey = jsonKey[:commaIndex]
		}

		if jsonKey == "" {
			continue
		}

		switch fieldValue.Kind() {
		case reflect.Int32, reflect.Struct:
			continue

		case reflect.Slice:
			var values []string

			for j := 0; j < fieldValue.Len(); j++ {
				elem := fieldValue.Index(j)
				if elem.Kind() == reflect.String {
					values = append(values, elem.String())
					continue
				}

				valueFormatted := strings.ToLower(fixEnumValueName(jsonKey, fmt.Sprintf("%v", elem.Interface())))
				values = append(values, valueFormatted)
			}

			rawMap[jsonKey] = values

		default:
			return nil, errorx.IllegalState.New("unsupported field type %v", fieldValue.Kind())
		}
	}

	return json.Marshal(rawMap)
}

func (p BlockStateProperties) MarshalJSON() ([]byte, error) {

	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	rawMap := make(map[string]string)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonKey := field.Tag.Get("json")
		fieldValue := v.Field(i)

		omitEmpty := strings.Contains(jsonKey, ",omitempty")

		if omitEmpty && fieldValue.IsZero() {
			continue
		}

		commaIndex := strings.Index(jsonKey, ",")
		if commaIndex != -1 {
			jsonKey = jsonKey[:commaIndex]
		}

		if jsonKey == "" {
			continue
		}

		switch fieldValue.Elem().Kind() {

		case reflect.Bool:
			rawMap[jsonKey] = strconv.FormatBool(fieldValue.Elem().Bool())
		case reflect.Int32:

			intValue := int32(fieldValue.Elem().Int())

			enumName, err := getEnumStringValue(jsonKey, intValue)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to parse enum field %v", field.Name)
			}

			if enumName == nil {
				rawMap[jsonKey] = strconv.Itoa(int(intValue))
				continue
			}

			rawMap[jsonKey] = strings.ToLower(*enumName)

		case reflect.Uint32:

			intValue := uint32(fieldValue.Elem().Uint())

			enumName, err := getEnumStringValue(jsonKey, int32(intValue))
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to parse enum field %v", field.Name)
			}

			if enumName == nil {
				rawMap[jsonKey] = strconv.Itoa(int(intValue))
				continue
			}

			rawMap[jsonKey] = strings.ToLower(*enumName)

		default:
			rawMap[jsonKey] = fieldValue.Elem().String()
		}
	}

	return json.Marshal(rawMap)
}

func (p *BlockStateProperties) UnmarshalJSON(data []byte) error {

	var rawMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMap); err != nil {
		return errorx.IllegalState.Wrap(err, "failed to unmarshal raw map")
	}

	v := reflect.ValueOf(p).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {

		field := t.Field(i)
		jsonKey := field.Tag.Get("json")

		commaIndex := strings.Index(jsonKey, ",")
		if commaIndex != -1 {
			jsonKey = jsonKey[:commaIndex]
		}

		rawValue, exists := rawMap[jsonKey]
		if !exists || !v.Field(i).CanSet() {
			continue
		}

		fieldValue := v.Field(i)

		foundValue, err := getEnumIntValue(jsonKey, rawValue)
		if err != nil {
			return errorx.IllegalState.Wrap(err, "failed to parse enum field %v", field.Name)
		}

		if foundValue != -1 {
			rawValue, err = json.Marshal(foundValue)
			if err != nil {
				return errorx.IllegalState.Wrap(err, "failed to marshal enum field %v", field.Name)
			}
		}

		if strings.TrimPrefix(fieldValue.Type().String(), "*") == "uint32" {
			intValue, err := strconv.ParseUint(strings.Trim(string(rawValue), "\""), 10, 32)
			if err != nil {
				return errorx.IllegalState.Wrap(err, "failed to parse uint32 field %v", field.Name)
			}

			marshalled, err := json.Marshal(uint32(intValue))
			if err != nil {
				return errorx.IllegalState.Wrap(err, "failed to marshal uint32 field %v", field.Name)
			}
			err = json.Unmarshal(marshalled, fieldValue.Addr().Interface())
			if err != nil {
				return errorx.IllegalState.Wrap(err, "failed to unmarshal uint32 field %v", field.Name)
			}
			continue
		}

		if strings.TrimPrefix(fieldValue.Type().String(), "*") == "bool" {
			boolValue, err := strconv.ParseBool(strings.Trim(string(rawValue), "\""))
			if err != nil {
				return errorx.IllegalState.Wrap(err, "failed to parse bool field %v", field.Name)
			}

			fieldValue.Set(reflect.ValueOf(&boolValue))
			continue
		}

		if err := json.Unmarshal(rawValue, fieldValue.Addr().Interface()); err != nil {
			return errorx.IllegalState.Wrap(err, "failed to unmarshal field %v", field.Name)
		}

	}

	return nil
}

func getEnumStringValue(jsonKey string, input int32) (foundValue *string, err error) {

	var (
		actualValue string
		exists      bool
	)

	switch jsonKey {

	case "attachment":
		actualValue, exists = Attachment_name[input]
	case "axis":
		actualValue, exists = Axis_name[input]
	case "east":
		actualValue, exists = East_name[input]
	case "face":
		actualValue, exists = Face_name[input]
	case "facing":
		actualValue, exists = Facing_name[input]
	case "half":
		actualValue, exists = Half_name[input]
	case "hinge":
		actualValue, exists = Hinge_name[input]
	case "instrument":
		actualValue, exists = Instrument_name[input]
	case "leaves":
		actualValue, exists = Leaves_name[input]
	case "mode":
		actualValue, exists = Mode_name[input]
	case "north":
		actualValue, exists = North_name[input]
	case "orientation":
		actualValue, exists = Orientation_name[input]
	case "part":
		actualValue, exists = Part_name[input]
	case "sculk_sensor_phase":
		actualValue, exists = SculkSensorPhase_name[input]
	case "shape":
		actualValue, exists = Shape_name[input]
	case "south":
		actualValue, exists = South_name[input]
	case "thickness":
		actualValue, exists = Thickness_name[input]
	case "tilt":
		actualValue, exists = Tilt_name[input]
	case "trial_spawner_state":
		actualValue, exists = TrialSpawnerState_name[input]
	case "type":
		actualValue, exists = Type_name[input]
	case "vertical_direction":
		actualValue, exists = VerticalDirection_name[input]
	case "west":
		actualValue, exists = West_name[input]

	default:
		return nil, nil
	}

	if !exists {
		return nil, errorx.IllegalState.New("failed to find enum value for %q=%v", jsonKey, input)
	}

	actualValue = fixEnumValueName(jsonKey, actualValue)

	return &actualValue, nil
}

func fixEnumValueName(jsonKey string, value string) string {

	value = strings.TrimPrefix(value, strings.ToUpper(jsonKey)+"_")

	if strings.HasPrefix(value, "SPAWNER_STATE_") {
		value = strings.TrimPrefix(value, "SPAWNER_STATE_")
	}
	if strings.HasPrefix(value, "SCULK_SENSOR_PHASE_") {
		value = strings.TrimPrefix(value, "SCULK_SENSOR_PHASE_")
	}

	return value
}

func getEnumIntValue(jsonKey string, value json.RawMessage) (foundValue int32, err error) {

	valueAsName := strings.ToUpper(fmt.Sprintf("%s_%s", jsonKey, strings.Trim(string(value), "\"")))

	var (
		actualValue int32
		exists      bool
	)

	switch jsonKey {

	case "attachment":
		actualValue, exists = Attachment_value[valueAsName]
	case "axis":
		actualValue, exists = Axis_value[valueAsName]
	case "east":
		actualValue, exists = East_value[valueAsName]
	case "face":
		actualValue, exists = Face_value[valueAsName]
	case "facing":
		actualValue, exists = Facing_value[valueAsName]
	case "half":
		actualValue, exists = Half_value[valueAsName]
	case "hinge":
		actualValue, exists = Hinge_value[valueAsName]
	case "instrument":
		actualValue, exists = Instrument_value[valueAsName]
	case "leaves":
		actualValue, exists = Leaves_value[valueAsName]
	case "mode":
		actualValue, exists = Mode_value[valueAsName]
	case "north":
		actualValue, exists = North_value[valueAsName]
	case "orientation":
		actualValue, exists = Orientation_value[valueAsName]
	case "part":
		actualValue, exists = Part_value[valueAsName]
	case "sculk_sensor_phase":
		actualValue, exists = SculkSensorPhase_value[valueAsName]
	case "shape":
		actualValue, exists = Shape_value[valueAsName]
	case "south":
		actualValue, exists = South_value[valueAsName]
	case "thickness":
		actualValue, exists = Thickness_value[valueAsName]
	case "tilt":
		actualValue, exists = Tilt_value[valueAsName]
	case "trial_spawner_state":
		actualValue, exists = TrialSpawnerState_value[valueAsName]
	case "type":
		actualValue, exists = Type_value[valueAsName]
	case "vertical_direction":
		actualValue, exists = VerticalDirection_value[valueAsName]
	case "west":
		actualValue, exists = West_value[valueAsName]

	default:
		return -1, nil
	}

	if !exists {
		return -1, errorx.IllegalState.New("failed to find enum value for %q=%q", jsonKey, valueAsName)
	}

	return actualValue, nil
}
