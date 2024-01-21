package palette

import (
	"encoding/json"
	"github.com/MisterKaiou/go-functional/result"
	"github.com/camdenorrb/minecraftPackets/javaEdition/palette/proto"
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
	Properties BlockPaletteStateProperties
}

type StateMap struct {
	StateToID map[BlockStateKey]uint
	IDToState map[uint]BlockStateKey
}

func (p BlockPaletteJSON) AsProto() result.Of[proto.BlockPalette] {

	protoPaletteEntries := make([]*proto.BlockPaletteEntry, 0, len(p))

	for material, data := range p {

		key := strings.ToUpper(strings.TrimPrefix(material, "minecraft:"))
		materialID := proto.Material_value[key]

		dataAsProto := data.AsProto()
		if dataAsProto.IsError() {
			return result.Error[proto.BlockPalette](
				errorx.IllegalState.Wrap(dataAsProto.UnwrapError(), "failed to convert data to proto"),
			)
		}

		dataAsProtoValue := dataAsProto.Unwrap()

		protoPaletteEntries = append(protoPaletteEntries, &proto.BlockPaletteEntry{
			Material: proto.Material(materialID),
			Data:     &dataAsProtoValue,
		})
	}

	return result.Ok(proto.BlockPalette{
		Entries: protoPaletteEntries,
	})
}

/*
 TODO: Do this on protobuf representation
// Name + state string -> ID
func (p *BlockPaletteJSON) AsStateMap() StateMap {

	stateToID := make(map[BlockStateKey]uint)
	idToState := make(map[uint]BlockStateKey)

	for name, data := range *p {
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
*/

type BlockPaletteData struct {
	Properties BlockPaletteProperties `json:"properties"`
	States     []BlockPaletteState    `json:"states"`
}

func (d BlockPaletteData) AsProto() result.Of[proto.BlockPaletteData] {

	protoStates := make([]*proto.BlockPaletteState, 0, len(d.States))

	for _, state := range d.States {

		stateAsProto := state.AsProto()
		if stateAsProto.IsError() {
			return result.Error[proto.BlockPaletteData](
				errorx.IllegalState.Wrap(stateAsProto.UnwrapError(), "failed to convert state to proto"),
			)
		}

		stateAsProtoValue := stateAsProto.Unwrap()
		protoStates = append(protoStates, &stateAsProtoValue)
	}

	propertiesAsProto := d.Properties.AsProto()
	if propertiesAsProto.IsError() {
		return result.Error[proto.BlockPaletteData](
			errorx.IllegalState.Wrap(propertiesAsProto.UnwrapError(), "failed to convert properties to proto"),
		)
	}

	propertiesAsProtoValue := propertiesAsProto.Unwrap()

	return result.Ok(proto.BlockPaletteData{
		Properties: &propertiesAsProtoValue,
		States:     protoStates,
	})
}

type BlockPaletteState struct {
	ID         uint                        `json:"id"`
	Properties BlockPaletteStateProperties `json:"properties"`
}

func (s BlockPaletteState) AsProto() result.Of[proto.BlockPaletteState] {

	propertiesAsProto := s.Properties.AsProto()
	if propertiesAsProto.IsError() {
		return result.Error[proto.BlockPaletteState](
			errorx.IllegalState.Wrap(propertiesAsProto.UnwrapError(), "failed to convert properties to proto"),
		)
	}

	propertiesAsProtoValue := propertiesAsProto.Unwrap()

	return result.Ok(proto.BlockPaletteState{
		Id:         uint32(s.ID),
		Properties: &propertiesAsProtoValue,
	})
}

// TODO: Make primitive arrays that can be parsed from strings
type BlockPaletteProperties struct {
	Age               []uint32            `json:"age"`
	Attached          []bool              `json:"attached"`
	Attachment        []Attachment        `json:"attachment"`
	Axis              []Axis              `json:"axis"`
	Berries           []bool              `json:"berries"`
	Bottom            []bool              `json:"bottom"`
	CanSummon         []bool              `json:"can_summon"`
	Candles           []uint32            `json:"candles"`
	Conditional       []bool              `json:"conditional"`
	Cracked           []bool              `json:"cracked"`
	Crafting          []bool              `json:"crafting"`
	Delay             []uint32            `json:"delay"`
	Disarmed          []bool              `json:"disarmed"`
	Distance          []uint32            `json:"distance"`
	Down              []bool              `json:"down"`
	Drag              []bool              `json:"drag"`
	Dusted            []uint32            `json:"dusted"`
	East              []East              `json:"east"`
	Eggs              []uint32            `json:"eggs"`
	Enabled           []bool              `json:"enabled"`
	Extended          []bool              `json:"extended"`
	Face              []Face              `json:"face"`
	Facing            []Facing            `json:"facing"`
	FlowerAmount      []uint32            `json:"flower_amount"`
	Half              []Half              `json:"half"`
	Hanging           []bool              `json:"hanging"`
	HasBook           []bool              `json:"has_book"`
	HasBottle0        []bool              `json:"has_bottle_0"`
	HasBottle1        []bool              `json:"has_bottle_1"`
	HasBottle2        []bool              `json:"has_bottle_2"`
	HasRecord         []bool              `json:"has_record"`
	Hinge             []Hinge             `json:"hinge"`
	Inverted          []bool              `json:"inverted"`
	Layers            []uint32            `json:"layers"`
	Leaves            []Leaves            `json:"leaves"`
	Lit               []bool              `json:"lit"`
	Locked            []bool              `json:"locked"`
	Mode              []Mode              `json:"mode"`
	Moisture          []uint32            `json:"moisture"`
	North             []North             `json:"north"`
	Note              []uint32            `json:"note"`
	Occupied          []bool              `json:"occupied"`
	Open              []bool              `json:"open"`
	Persistent        []bool              `json:"persistent"`
	Pickles           []uint32            `json:"pickles"`
	Powered           []bool              `json:"powered"`
	Rotation          []uint32            `json:"rotation"`
	SculkSensorPhase  []SculkSensorPhase  `json:"sculk_sensor_phase"`
	Shape             []Shape             `json:"shape"`
	Short             []bool              `json:"short"`
	Shrieking         []bool              `json:"shrieking"`
	SignalFire        []bool              `json:"signal_fire"`
	Slot0Occupied     []bool              `json:"slot_0_occupied"`
	Slot1Occupied     []bool              `json:"slot_1_occupied"`
	Slot2Occupied     []bool              `json:"slot_2_occupied"`
	Slot3Occupied     []bool              `json:"slot_3_occupied"`
	Slot4Occupied     []bool              `json:"slot_4_occupied"`
	Slot5Occupied     []bool              `json:"slot_5_occupied"`
	Snowy             []bool              `json:"snowy"`
	South             []South             `json:"south"`
	Stage             []uint32            `json:"stage"`
	Thickness         []Thickness         `json:"thickness"`
	Tilt              []Tilt              `json:"tilt"`
	Triggered         []bool              `json:"triggered"`
	Type              []Type              `json:"type"`
	Unstable          []bool              `json:"unstable"`
	Up                []bool              `json:"up"`
	VerticalDirection []VerticalDirection `json:"vertical_direction"`
	Waterlogged       []bool              `json:"waterlogged"`
	West              []West              `json:"west"`
}

// UnmarshalJSON custom unmarshaler for BlockPaletteProperties
// Values in the json are wrapped with quotes, which is why we need this to go to primitive values
func (p BlockPaletteProperties) UnmarshalJSON(data []byte) error {

	var rawMap map[string][]json.RawMessage
	if err := json.Unmarshal(data, &rawMap); err != nil {
		return errorx.IllegalState.Wrap(err, "failed to unmarshal raw map")
	}

	v := reflect.ValueOf(&p).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {

		field := t.Field(i)
		jsonKey := field.Tag.Get("json")

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
				return err
			}

			if err := json.Unmarshal(marshal, fieldValue.Addr().Interface()); err != nil {
				return errorx.IllegalState.Wrap(err, "failed to unmarshal field %v", field.Name)
			}
		}
	}

	return nil
}

func (p BlockPaletteProperties) AsProto() result.Of[proto.BlockProperties] {

	var attachment []proto.Attachment
	if p.Attachment != nil {

		attachment = make([]proto.Attachment, 0, len(p.Attachment))

		for _, value := range p.Attachment {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert attachment to proto"),
				)
			}

			attachmentValue := asProto.Unwrap()
			attachment = append(attachment, attachmentValue)
		}
	}

	var axis []proto.Axis
	if p.Axis != nil {

		axis = make([]proto.Axis, 0, len(p.Axis))

		for _, value := range p.Axis {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert axis to proto"),
				)
			}

			axisValue := asProto.Unwrap()
			axis = append(axis, axisValue)
		}
	}

	var east []proto.East
	if p.East != nil {

		east = make([]proto.East, 0, len(p.East))

		for _, value := range p.East {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert east to proto"),
				)
			}

			eastValue := asProto.Unwrap()
			east = append(east, eastValue)
		}
	}

	var face []proto.Face
	if p.Face != nil {

		face = make([]proto.Face, 0, len(p.Face))

		for _, value := range p.Face {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert face to proto"),
				)
			}

			faceValue := asProto.Unwrap()
			face = append(face, faceValue)
		}
	}

	var facing []proto.Facing
	if p.Facing != nil {

		facing = make([]proto.Facing, 0, len(p.Facing))

		for _, value := range p.Facing {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert facing to proto"),
				)
			}

			facingValue := asProto.Unwrap()
			facing = append(facing, facingValue)
		}
	}

	var half []proto.Half
	if p.Half != nil {

		half = make([]proto.Half, 0, len(p.Half))

		for _, value := range p.Half {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert half to proto"),
				)
			}

			halfValue := asProto.Unwrap()
			half = append(half, halfValue)
		}
	}

	var hinge []proto.Hinge
	if p.Hinge != nil {

		hinge = make([]proto.Hinge, 0, len(p.Hinge))

		for _, value := range p.Hinge {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert hinge to proto"),
				)
			}

			hingeValue := asProto.Unwrap()
			hinge = append(hinge, hingeValue)
		}
	}

	var leaves []proto.Leaves
	if p.Leaves != nil {

		leaves = make([]proto.Leaves, 0, len(p.Leaves))

		for _, value := range p.Leaves {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert leaves to proto"),
				)
			}

			leavesValue := asProto.Unwrap()
			leaves = append(leaves, leavesValue)
		}
	}

	var mode []proto.Mode
	if p.Mode != nil {

		mode = make([]proto.Mode, 0, len(p.Mode))

		for _, value := range p.Mode {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert mode to proto"),
				)
			}

			modeValue := asProto.Unwrap()
			mode = append(mode, modeValue)
		}

	}

	var north []proto.North
	if p.North != nil {

		north = make([]proto.North, 0, len(p.North))

		for _, value := range p.North {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), ""),
				)
			}

			northValue := asProto.Unwrap()
			north = append(north, northValue)
		}
	}

	var sculkSensorPhase []proto.SculkSensorPhase
	if p.SculkSensorPhase != nil {

		sculkSensorPhase = make([]proto.SculkSensorPhase, 0, len(p.SculkSensorPhase))

		for _, value := range p.SculkSensorPhase {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), ""),
				)
			}

			sculkSensorPhaseValue := asProto.Unwrap()
			sculkSensorPhase = append(sculkSensorPhase, sculkSensorPhaseValue)
		}
	}

	var shape []proto.Shape
	if p.Shape != nil {

		shape = make([]proto.Shape, 0, len(p.Shape))

		for _, value := range p.Shape {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), ""),
				)
			}

			shapeValue := asProto.Unwrap()
			shape = append(shape, shapeValue)
		}
	}

	var south []proto.South
	if p.South != nil {

		south = make([]proto.South, 0, len(p.South))

		for _, value := range p.South {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), ""),
				)
			}

			southValue := asProto.Unwrap()
			south = append(south, southValue)
		}

	}

	var thickness []proto.Thickness
	if p.Thickness != nil {

		thickness = make([]proto.Thickness, 0, len(p.Thickness))

		for _, value := range p.Thickness {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), ""),
				)
			}

			thicknessValue := asProto.Unwrap()
			thickness = append(thickness, thicknessValue)
		}
	}

	var tilt []proto.Tilt
	if p.Tilt != nil {

		tilt = make([]proto.Tilt, 0, len(p.Tilt))

		for _, value := range p.Tilt {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), ""),
				)
			}

			tiltValue := asProto.Unwrap()
			tilt = append(tilt, tiltValue)
		}
	}

	var typeProperty []proto.Type
	if p.Type != nil {

		typeProperty = make([]proto.Type, 0, len(p.Type))

		for _, value := range p.Type {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), ""),
				)
			}

			typePropertyValue := asProto.Unwrap()
			typeProperty = append(typeProperty, typePropertyValue)
		}
	}

	var verticalDirection []proto.VerticalDirection
	if p.VerticalDirection != nil {

		verticalDirection = make([]proto.VerticalDirection, 0, len(p.VerticalDirection))

		for _, value := range p.VerticalDirection {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), ""),
				)
			}

			verticalDirectionValue := asProto.Unwrap()
			verticalDirection = append(verticalDirection, verticalDirectionValue)
		}
	}

	var west []proto.West
	if p.West != nil {

		west = make([]proto.West, 0, len(p.West))

		for _, value := range p.West {

			asProto := value.AsProto()
			if asProto.IsError() {
				return result.Error[proto.BlockProperties](
					errorx.IllegalState.Wrap(asProto.UnwrapError(), ""),
				)
			}

			westValue := asProto.Unwrap()
			west = append(west, westValue)
		}
	}

	return result.Ok(proto.BlockProperties{
		Age:               p.Age,
		Attached:          p.Attached,
		Attachment:        attachment,
		Axis:              axis,
		Berries:           p.Berries,
		Bottom:            p.Bottom,
		CanSummon:         p.CanSummon,
		Candles:           p.Candles,
		Conditional:       p.Conditional,
		Cracked:           p.Cracked,
		Crafting:          p.Crafting,
		Delay:             p.Delay,
		Disarmed:          p.Disarmed,
		Distance:          p.Distance,
		Down:              p.Down,
		Drag:              p.Drag,
		Dusted:            p.Dusted,
		East:              east,
		Eggs:              p.Eggs,
		Enabled:           p.Enabled,
		Extended:          p.Extended,
		Face:              face,
		Facing:            facing,
		FlowerAmount:      p.FlowerAmount,
		Half:              half,
		Hanging:           p.Hanging,
		HasBook:           p.HasBook,
		HasBottle_0:       p.HasBottle0,
		HasBottle_1:       p.HasBottle1,
		HasBottle_2:       p.HasBottle2,
		HasRecord:         p.HasRecord,
		Hinge:             hinge,
		Inverted:          p.Inverted,
		Layers:            p.Layers,
		Leaves:            leaves,
		Lit:               p.Lit,
		Locked:            p.Locked,
		Mode:              mode,
		Moisture:          p.Moisture,
		North:             north,
		Note:              p.Note,
		Occupied:          p.Occupied,
		Open:              p.Open,
		Persistent:        p.Persistent,
		Pickles:           p.Pickles,
		Powered:           p.Powered,
		Rotation:          p.Rotation,
		SculkSensorPhase:  sculkSensorPhase,
		Shape:             shape,
		Short:             p.Short,
		Shrieking:         p.Shrieking,
		SignalFire:        p.SignalFire,
		Slot_0Occupied:    p.Slot0Occupied,
		Slot_1Occupied:    p.Slot1Occupied,
		Slot_2Occupied:    p.Slot2Occupied,
		Slot_3Occupied:    p.Slot3Occupied,
		Slot_4Occupied:    p.Slot4Occupied,
		Slot_5Occupied:    p.Slot5Occupied,
		Snowy:             p.Snowy,
		South:             south,
		Stage:             p.Stage,
		Thickness:         thickness,
		Tilt:              tilt,
		Triggered:         p.Triggered,
		Type:              typeProperty,
		Unstable:          p.Unstable,
		Up:                p.Up,
		VerticalDirection: verticalDirection,
		Waterlogged:       p.Waterlogged,
		West:              west,
	})
}

// Singular values
type BlockPaletteStateProperties struct {
	Age               *uint32            `json:"age,string"`
	Attached          *bool              `json:"attached,string"`
	Attachment        *Attachment        `json:"attachment"`
	Axis              *Axis              `json:"axis"`
	Berries           *bool              `json:"berries,string"`
	Bottom            *bool              `json:"bottom,string"`
	CanSummon         *bool              `json:"can_summon,string"`
	Candles           *uint32            `json:"candles,string"`
	Conditional       *bool              `json:"conditional,string"`
	Cracked           *bool              `json:"cracked,string"`
	Crafting          *bool              `json:"crafting,string"`
	Delay             *uint32            `json:"delay,string"`
	Disarmed          *bool              `json:"disarmed,string"`
	Distance          *uint32            `json:"distance,string"`
	Down              *bool              `json:"down,string"`
	Drag              *bool              `json:"drag,string"`
	Dusted            *uint32            `json:"dusted,string"`
	East              *East              `json:"east"`
	Eggs              *uint32            `json:"eggs,string"`
	Enabled           *bool              `json:"enabled,string"`
	Extended          *bool              `json:"extended,string"`
	Face              *Face              `json:"face"`
	Facing            *Facing            `json:"facing"`
	FlowerAmount      *uint32            `json:"flower_amount,string"`
	Half              *Half              `json:"half"`
	Hanging           *bool              `json:"hanging,string"`
	HasBook           *bool              `json:"has_book,string"`
	HasBottle0        *bool              `json:"has_bottle_0,string"`
	HasBottle1        *bool              `json:"has_bottle_1,string"`
	HasBottle2        *bool              `json:"has_bottle_2,string"`
	HasRecord         *bool              `json:"has_record,string"`
	Hinge             *Hinge             `json:"hinge"`
	Inverted          *bool              `json:"inverted,string"`
	Layers            *uint32            `json:"layers,string"`
	Leaves            *Leaves            `json:"leaves"`
	Lit               *bool              `json:"lit,string"`
	Locked            *bool              `json:"locked,string"`
	Mode              *Mode              `json:"mode"`
	Moisture          *uint32            `json:"moisture,string"`
	North             *North             `json:"north"`
	Note              *uint32            `json:"note,string"`
	Occupied          *bool              `json:"occupied,string"`
	Open              *bool              `json:"open,string"`
	Persistent        *bool              `json:"persistent,string"`
	Pickles           *uint32            `json:"pickles,string"`
	Powered           *bool              `json:"powered,string"`
	Rotation          *uint32            `json:"rotation,string"`
	SculkSensorPhase  *SculkSensorPhase  `json:"sculk_sensor_phase"`
	Shape             *Shape             `json:"shape"`
	Short             *bool              `json:"short,string"`
	Shrieking         *bool              `json:"shrieking,string"`
	SignalFire        *bool              `json:"signal_fire,string"`
	Slot0Occupied     *bool              `json:"slot_0_occupied,string"`
	Slot1Occupied     *bool              `json:"slot_1_occupied,string"`
	Slot2Occupied     *bool              `json:"slot_2_occupied,string"`
	Slot3Occupied     *bool              `json:"slot_3_occupied,string"`
	Slot4Occupied     *bool              `json:"slot_4_occupied,string"`
	Slot5Occupied     *bool              `json:"slot_5_occupied,string"`
	Snowy             *bool              `json:"snowy,string"`
	South             *South             `json:"south"`
	Stage             *uint32            `json:"stage,string"`
	Thickness         *Thickness         `json:"thickness"`
	Tilt              *Tilt              `json:"tilt"`
	Triggered         *bool              `json:"triggered,string"`
	Type              *Type              `json:"type"`
	Unstable          *bool              `json:"unstable,string"`
	Up                *bool              `json:"up,string"`
	VerticalDirection *VerticalDirection `json:"vertical_direction"`
	Waterlogged       *bool              `json:"waterlogged,string"`
	West              *West              `json:"west"`
}

func (p BlockPaletteStateProperties) AsProto() result.Of[proto.BlockStateProperties] {

	var attachment *proto.Attachment
	if p.Attachment != nil {

		asProto := p.Attachment.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert attachment to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		attachment = &unwrapped
	}

	var axis *proto.Axis
	if p.Axis != nil {

		asProto := p.Axis.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert axis to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		axis = &unwrapped
	}

	var east *proto.East
	if p.East != nil {

		asProto := p.East.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert east to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		east = &unwrapped
	}

	var face *proto.Face
	if p.Face != nil {

		asProto := p.Face.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert face to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		face = &unwrapped
	}

	var facing *proto.Facing
	if p.Facing != nil {

		asProto := p.Facing.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert facing to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		facing = &unwrapped
	}

	var half *proto.Half
	if p.Half != nil {

		asProto := p.Half.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert half to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		half = &unwrapped
	}

	var hinge *proto.Hinge
	if p.Hinge != nil {

		asProto := p.Hinge.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert hinge to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		hinge = &unwrapped
	}

	var leaves *proto.Leaves
	if p.Leaves != nil {

		asProto := p.Leaves.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert leaves to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		leaves = &unwrapped
	}

	var mode *proto.Mode
	if p.Mode != nil {

		asProto := p.Mode.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert mode to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		mode = &unwrapped
	}

	var north *proto.North
	if p.North != nil {

		asProto := p.North.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert north to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		north = &unwrapped
	}

	var sculkSensorPhase *proto.SculkSensorPhase
	if p.SculkSensorPhase != nil {

		asProto := p.SculkSensorPhase.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert sculk sensor phase to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		sculkSensorPhase = &unwrapped
	}

	var shape *proto.Shape
	if p.Shape != nil {

		asProto := p.Shape.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert shape to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		shape = &unwrapped
	}

	var south *proto.South
	if p.South != nil {

		asProto := p.South.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed convert south to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		south = &unwrapped
	}

	var thickness *proto.Thickness
	if p.Thickness != nil {

		asProto := p.Thickness.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to convert thickness to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		thickness = &unwrapped
	}

	var tilt *proto.Tilt
	if p.Tilt != nil {

		asProto := p.Tilt.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed to tilt to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		tilt = &unwrapped
	}

	var typeProperty *proto.Type
	if p.Type != nil {

		asProto := p.Type.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), ""),
			)
		}

		unwrapped := asProto.Unwrap()
		typeProperty = &unwrapped
	}

	var verticalDirection *proto.VerticalDirection
	if p.VerticalDirection != nil {

		asProto := p.VerticalDirection.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed convert vertical direction to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		verticalDirection = &unwrapped
	}

	var west *proto.West
	if p.West != nil {

		asProto := p.West.AsProto()
		if asProto.IsError() {
			return result.Error[proto.BlockStateProperties](
				errorx.IllegalState.Wrap(asProto.UnwrapError(), "failed convert west to proto"),
			)
		}

		unwrapped := asProto.Unwrap()
		west = &unwrapped
	}

	return result.Ok(proto.BlockStateProperties{
		Age:               p.Age,
		Attached:          p.Attached,
		Attachment:        attachment,
		Axis:              axis,
		Berries:           p.Berries,
		Bottom:            p.Bottom,
		CanSummon:         p.CanSummon,
		Candles:           p.Candles,
		Conditional:       p.Conditional,
		Cracked:           p.Cracked,
		Crafting:          p.Crafting,
		Delay:             p.Delay,
		Disarmed:          p.Disarmed,
		Distance:          p.Distance,
		Down:              p.Down,
		Drag:              p.Drag,
		Dusted:            p.Dusted,
		East:              east,
		Eggs:              p.Eggs,
		Enabled:           p.Enabled,
		Extended:          p.Extended,
		Face:              face,
		Facing:            facing,
		FlowerAmount:      p.FlowerAmount,
		Half:              half,
		Hanging:           p.Hanging,
		HasBook:           p.HasBook,
		HasBottle_0:       p.HasBottle0,
		HasBottle_1:       p.HasBottle1,
		HasBottle_2:       p.HasBottle2,
		HasRecord:         p.HasRecord,
		Hinge:             hinge,
		Inverted:          p.Inverted,
		Layers:            p.Layers,
		Leaves:            leaves,
		Lit:               p.Lit,
		Locked:            p.Locked,
		Mode:              mode,
		Moisture:          p.Moisture,
		North:             north,
		Note:              p.Note,
		Occupied:          p.Occupied,
		Open:              p.Open,
		Persistent:        p.Persistent,
		Pickles:           p.Pickles,
		Powered:           p.Powered,
		Rotation:          p.Rotation,
		SculkSensorPhase:  sculkSensorPhase,
		Shape:             shape,
		Short:             p.Short,
		Shrieking:         p.Shrieking,
		SignalFire:        p.SignalFire,
		Slot_0Occupied:    p.Slot0Occupied,
		Slot_1Occupied:    p.Slot1Occupied,
		Slot_2Occupied:    p.Slot2Occupied,
		Slot_3Occupied:    p.Slot3Occupied,
		Slot_4Occupied:    p.Slot4Occupied,
		Slot_5Occupied:    p.Slot5Occupied,
		Snowy:             p.Snowy,
		South:             south,
		Stage:             p.Stage,
		Thickness:         thickness,
		Tilt:              tilt,
		Triggered:         p.Triggered,
		Type:              typeProperty,
		Unstable:          p.Unstable,
		Up:                p.Up,
		VerticalDirection: verticalDirection,
		Waterlogged:       p.Waterlogged,
		West:              west,
	})
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
