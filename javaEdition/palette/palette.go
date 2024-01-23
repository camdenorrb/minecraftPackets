package palette

import (
	"encoding/json"
	"fmt"
	"github.com/joomcode/errorx"
	"reflect"
	"strconv"
	"strings"
)

// TODO: Read from https://wiki.vg/Data_Generators

// TODO:
// 	generated blocks.json -> blocks.proto
// 	Have one struct representing original data
// 	Make map from struct to do state as string -> ID
// 	Store map in registry on registry load

// TODO:
// 	 For older versions which don't have blocks.json
//	 Have a way to go from map back to struct, to generate a ~"blocks.json"

// TODO: Move proto files to protocol folder based on version

type BlockPaletteJSON map[string]BlockPaletteData

type BlockStateKey struct {
	Name       string
	Properties *BlockPaletteStateProperties
}

type StateMap struct {
	StateToID map[BlockStateKey]uint
	IDToState map[uint]BlockStateKey
}

func (p BlockPaletteJSON) AsStateMap() StateMap {

	stateToID := make(map[BlockStateKey]uint)
	idToState := make(map[uint]BlockStateKey)

	for name, data := range p {
		for _, state := range data.States {

			key := BlockStateKey{
				Name:       name,
				Properties: state.Properties,
			}

			stateToID[key] = state.ID
			idToState[state.ID] = key
		}
	}

	return StateMap{
		StateToID: stateToID,
		IDToState: idToState,
	}
}

type BlockPaletteData struct {
	Properties *BlockPaletteProperties `json:"properties,omitempty"`
	States     []BlockPaletteState     `json:"states,omitempty"`
}

type BlockPaletteState struct {
	Default    *bool                        `json:"default,omitempty"`
	ID         uint                         `json:"id"`
	Properties *BlockPaletteStateProperties `json:"properties,omitempty"`
}

// TODO: Make primitive arrays that can be parsed from strings
type BlockPaletteProperties struct {
	Age               []uint32            `json:"age,omitempty"`
	Attached          []bool              `json:"attached,omitempty"`
	Attachment        []Attachment        `json:"attachment,omitempty"`
	Axis              []Axis              `json:"axis,omitempty"`
	Berries           []bool              `json:"berries,omitempty"`
	Bites             []uint32            `json:"bites,omitempty"`
	Bottom            []bool              `json:"bottom,omitempty"`
	Bloom             []bool              `json:"bloom,omitempty"`
	CanSummon         []bool              `json:"can_summon,omitempty"`
	Candles           []uint32            `json:"candles,omitempty"`
	Conditional       []bool              `json:"conditional,omitempty"`
	Cracked           []bool              `json:"cracked,omitempty"`
	Crafting          []bool              `json:"crafting,omitempty"`
	Charges           []uint32            `json:"charges,omitempty"`
	Delay             []uint32            `json:"delay,omitempty"`
	Disarmed          []bool              `json:"disarmed,omitempty"`
	Distance          []uint32            `json:"distance,omitempty"`
	Down              []bool              `json:"down,omitempty"`
	Drag              []bool              `json:"drag,omitempty"`
	Dusted            []uint32            `json:"dusted,omitempty"`
	East              []East              `json:"east,omitempty"`
	Eggs              []uint32            `json:"eggs,omitempty"`
	Eye               []bool              `json:"eye,omitempty"`
	Enabled           []bool              `json:"enabled,omitempty"`
	Extended          []bool              `json:"extended,omitempty"`
	Face              []Face              `json:"face,omitempty"`
	Facing            []Facing            `json:"facing,omitempty"`
	FlowerAmount      []uint32            `json:"flower_amount,omitempty"`
	Half              []Half              `json:"half,omitempty"`
	Hanging           []bool              `json:"hanging,omitempty"`
	HasBook           []bool              `json:"has_book,omitempty"`
	HasBottle0        []bool              `json:"has_bottle_0,omitempty"`
	HasBottle1        []bool              `json:"has_bottle_1,omitempty"`
	HasBottle2        []bool              `json:"has_bottle_2,omitempty"`
	HasRecord         []bool              `json:"has_record,omitempty"`
	Hatch             []uint32            `json:"hatch,omitempty"`
	Hinge             []Hinge             `json:"hinge,omitempty"`
	HoneyLevel        []uint32            `json:"honey_level,omitempty"`
	Instrument        []Instrument        `json:"instrument,omitempty"`
	Inverted          []bool              `json:"inverted,omitempty"`
	Layers            []uint32            `json:"layers,omitempty"`
	Leaves            []Leaves            `json:"leaves,omitempty"`
	Level             []uint32            `json:"level,omitempty"`
	Lit               []bool              `json:"lit,omitempty"`
	Locked            []bool              `json:"locked,omitempty"`
	Mode              []Mode              `json:"mode,omitempty"`
	Moisture          []uint32            `json:"moisture,omitempty"`
	North             []North             `json:"north,omitempty"`
	Note              []uint32            `json:"note,omitempty"`
	Occupied          []bool              `json:"occupied,omitempty"`
	InWall            []bool              `json:"in_wall,omitempty"`
	Open              []bool              `json:"open,omitempty"`
	Orientation       []Orientation       `json:"orientation,omitempty"`
	Part              []Part              `json:"part,omitempty"`
	Persistent        []bool              `json:"persistent,omitempty"`
	Pickles           []uint32            `json:"pickles,omitempty"`
	Power             []uint32            `json:"power,omitempty"`
	Powered           []bool              `json:"powered,omitempty"`
	Rotation          []uint32            `json:"rotation,omitempty"`
	SculkSensorPhase  []SculkSensorPhase  `json:"sculk_sensor_phase,omitempty"`
	Shape             []Shape             `json:"shape,omitempty"`
	Short             []bool              `json:"short,omitempty"`
	Shrieking         []bool              `json:"shrieking,omitempty"`
	SignalFire        []bool              `json:"signal_fire,omitempty"`
	Slot0Occupied     []bool              `json:"slot_0_occupied,omitempty"`
	Slot1Occupied     []bool              `json:"slot_1_occupied,omitempty"`
	Slot2Occupied     []bool              `json:"slot_2_occupied,omitempty"`
	Slot3Occupied     []bool              `json:"slot_3_occupied,omitempty"`
	Slot4Occupied     []bool              `json:"slot_4_occupied,omitempty"`
	Slot5Occupied     []bool              `json:"slot_5_occupied,omitempty"`
	Snowy             []bool              `json:"snowy,omitempty"`
	South             []South             `json:"south,omitempty"`
	Stage             []uint32            `json:"stage,omitempty"`
	Thickness         []Thickness         `json:"thickness,omitempty"`
	Tilt              []Tilt              `json:"tilt,omitempty"`
	TrialSpawnerState []TrialSpawnerState `json:"trial_spawner_state,omitempty"`
	Triggered         []bool              `json:"triggered,omitempty"`
	Type              []Type              `json:"type,omitempty"`
	Unstable          []bool              `json:"unstable,omitempty"`
	Up                []bool              `json:"up,omitempty"`
	VerticalDirection []VerticalDirection `json:"vertical_direction,omitempty"`
	Waterlogged       []bool              `json:"waterlogged,omitempty"`
	West              []West              `json:"west,omitempty"`
}

// UnmarshalJSON custom unmarshaler for BlockPaletteProperties
// Values in the json are wrapped with quotes, which is why we need this to go to primitive values
// NOTE: Needs to be a pointer
func (p *BlockPaletteProperties) UnmarshalJSON(data []byte) error {

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

		switch v.Field(i).Type().Elem().Kind() {
		case reflect.Uint32:

			var values []uint32
			for _, rawValue := range rawValues {

				intValue, err := strconv.ParseUint(strings.Trim(string(rawValue), "\""), 10, 32)
				if err != nil {
					return errorx.IllegalState.Wrap(err, "failed to parse uint32 field %v", field.Name)
				}

				values = append(values, uint32(intValue))
			}

			v.Field(i).Set(reflect.ValueOf(values))

		case reflect.Bool:

			var values []bool
			for _, rawValue := range rawValues {

				boolValue, err := strconv.ParseBool(strings.Trim(string(rawValue), "\""))
				if err != nil {
					return errorx.IllegalState.Wrap(err, "failed to parse bool field %v", field.Name)
				}

				values = append(values, boolValue)
			}

			v.Field(i).Set(reflect.ValueOf(values))

		default:

			fieldValue := v.Field(i)

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

func (p BlockPaletteProperties) MarshalJSON() ([]byte, error) {
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)
	rawMap := make(map[string][]string)

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

		switch fieldValue.Kind() {
		case reflect.Slice:
			var values []string

			for j := 0; j < fieldValue.Len(); j++ {
				elem := fieldValue.Index(j)
				if elem.Kind() == reflect.String {
					values = append(values, elem.String())
				} else {
					values = append(values, fmt.Sprintf("%v", elem.Interface()))
				}
			}

			rawMap[jsonKey] = values

		default:
			return nil, errorx.IllegalState.New("unsupported field type %v", fieldValue.Kind())
		}
	}

	return json.Marshal(rawMap)
}

// Singular values
type BlockPaletteStateProperties struct {
	Age               *uint32            `json:"age,string,omitempty"`
	Attached          *bool              `json:"attached,string,omitempty"`
	Attachment        *Attachment        `json:"attachment,omitempty"`
	Axis              *Axis              `json:"axis,omitempty"`
	Berries           *bool              `json:"berries,string,omitempty"`
	Bites             *uint32            `json:"bites,string,omitempty"`
	Bottom            *bool              `json:"bottom,string,omitempty"`
	Bloom             *bool              `json:"bloom,string,omitempty"`
	CanSummon         *bool              `json:"can_summon,string,omitempty"`
	Candles           *uint32            `json:"candles,string,omitempty"`
	Conditional       *bool              `json:"conditional,string,omitempty"`
	Cracked           *bool              `json:"cracked,string,omitempty"`
	Crafting          *bool              `json:"crafting,string,omitempty"`
	Charges           *uint32            `json:"charges,string,omitempty"`
	Delay             *uint32            `json:"delay,string,omitempty"`
	Disarmed          *bool              `json:"disarmed,string,omitempty"`
	Distance          *uint32            `json:"distance,string,omitempty"`
	Down              *bool              `json:"down,string,omitempty"`
	Drag              *bool              `json:"drag,string,omitempty"`
	Dusted            *uint32            `json:"dusted,string,omitempty"`
	East              *East              `json:"east,omitempty"`
	Eggs              *uint32            `json:"eggs,string,omitempty"`
	Eye               *bool              `json:"eye,string,omitempty"`
	Enabled           *bool              `json:"enabled,string,omitempty"`
	Extended          *bool              `json:"extended,string,omitempty"`
	Face              *Face              `json:"face,omitempty"`
	Facing            *Facing            `json:"facing,omitempty"`
	FlowerAmount      *uint32            `json:"flower_amount,string,omitempty"`
	Half              *Half              `json:"half,omitempty"`
	Hanging           *bool              `json:"hanging,string,omitempty"`
	HasBook           *bool              `json:"has_book,string,omitempty"`
	HasBottle0        *bool              `json:"has_bottle_0,string,omitempty"`
	HasBottle1        *bool              `json:"has_bottle_1,string,omitempty"`
	HasBottle2        *bool              `json:"has_bottle_2,string,omitempty"`
	HasRecord         *bool              `json:"has_record,string,omitempty"`
	Hatch             *uint32            `json:"hatch,string,omitempty"`
	Hinge             *Hinge             `json:"hinge,omitempty"`
	HoneyLevel        *uint32            `json:"honey_level,string,omitempty"`
	Instrument        *Instrument        `json:"instrument,omitempty"`
	Inverted          *bool              `json:"inverted,string,omitempty"`
	Layers            *uint32            `json:"layers,string,omitempty"`
	Leaves            *Leaves            `json:"leaves,omitempty"`
	Level             *uint32            `json:"level,string,omitempty"`
	Lit               *bool              `json:"lit,string,omitempty"`
	Locked            *bool              `json:"locked,string,omitempty"`
	Mode              *Mode              `json:"mode,omitempty"`
	Moisture          *uint32            `json:"moisture,string,omitempty"`
	North             *North             `json:"north,omitempty"`
	Note              *uint32            `json:"note,string,omitempty"`
	Occupied          *bool              `json:"occupied,string,omitempty"`
	Open              *bool              `json:"open,string,omitempty"`
	InWall            *bool              `json:"in_wall,string,omitempty"`
	Orientation       *Orientation       `json:"orientation,omitempty"`
	Part              *Part              `json:"part,omitempty"`
	Persistent        *bool              `json:"persistent,string,omitempty"`
	Pickles           *uint32            `json:"pickles,string,omitempty"`
	Power             *uint32            `json:"power,string,omitempty"`
	Powered           *bool              `json:"powered,string,omitempty"`
	Rotation          *uint32            `json:"rotation,string,omitempty"`
	SculkSensorPhase  *SculkSensorPhase  `json:"sculk_sensor_phase,omitempty"`
	Shape             *Shape             `json:"shape,omitempty"`
	Short             *bool              `json:"short,string,omitempty"`
	Shrieking         *bool              `json:"shrieking,string,omitempty"`
	SignalFire        *bool              `json:"signal_fire,string,omitempty"`
	Slot0Occupied     *bool              `json:"slot_0_occupied,string,omitempty"`
	Slot1Occupied     *bool              `json:"slot_1_occupied,string,omitempty"`
	Slot2Occupied     *bool              `json:"slot_2_occupied,string,omitempty"`
	Slot3Occupied     *bool              `json:"slot_3_occupied,string,omitempty"`
	Slot4Occupied     *bool              `json:"slot_4_occupied,string,omitempty"`
	Slot5Occupied     *bool              `json:"slot_5_occupied,string,omitempty"`
	Snowy             *bool              `json:"snowy,string,omitempty"`
	South             *South             `json:"south,omitempty"`
	Stage             *uint32            `json:"stage,string,omitempty"`
	Thickness         *Thickness         `json:"thickness,omitempty"`
	Tilt              *Tilt              `json:"tilt,omitempty"`
	TrialSpawnerState *TrialSpawnerState `json:"trial_spawner_state,omitempty"`
	Triggered         *bool              `json:"triggered,string,omitempty"`
	Type              *Type              `json:"type,omitempty"`
	Unstable          *bool              `json:"unstable,string,omitempty"`
	Up                *bool              `json:"up,string,omitempty"`
	VerticalDirection *VerticalDirection `json:"vertical_direction,omitempty"`
	Waterlogged       *bool              `json:"waterlogged,string,omitempty"`
	West              *West              `json:"west,omitempty"`
}

/*
func (p *BlockPaletteStateProperties) AsString() string {

	// Join all not null values with a comma
	element := reflect.ValueOf(p).Elem()
	elementType := element.Type()

	values := make([]string, 0, element.NumField())
	for i := 0; i < element.NumField(); i++ {

		field := element.Field(i)

		if field.IsNil() {
			continue
		}

		if field.Elem().Kind() == reflect.Bool {
			values = append(values, elementType.Field(i).Name+"="+strconv.FormatBool(field.Elem().Bool()))
		} else {
			values = append(values, elementType.Field(i).Name+"="+field.Elem().String())
		}
	}

	// Sort values
	sort.Strings(values)

	return "{" + strings.Join(values, ",") + "}"
}
*/
