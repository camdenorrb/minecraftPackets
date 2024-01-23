package proto

import (
	"fmt"
	"github.com/camdenorrb/minecraftPackets/javaEdition/palette"
	"github.com/joomcode/errorx"
	"strings"
)

func FromBlockPalette(blockPalette palette.BlockPaletteJSON) (*BlockPalette, error) {

	protoPaletteEntries := make([]*BlockPaletteEntry, 0, len(blockPalette))

	for material, data := range blockPalette {

		key := strings.ToUpper(strings.TrimPrefix(material, "minecraft:"))
		materialID := Material_value[key]

		dataAsProto, err := FromBlockPaletteData(data)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert data to proto")
		}

		protoPaletteEntries = append(protoPaletteEntries, &BlockPaletteEntry{
			Material: Material(materialID),
			Data:     dataAsProto,
		})
	}

	return &BlockPalette{
		Entries: protoPaletteEntries,
	}, nil
}

func FromBlockPaletteData(data palette.BlockPaletteData) (*BlockPaletteData, error) {

	protoStates := make([]*BlockPaletteState, len(data.States))

	for index, state := range data.States {

		stateAsProto, err := FromBlockPaletteState(state)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert state to proto")
		}

		protoStates[index] = stateAsProto
	}

	var propertiesAsProto *BlockProperties

	if data.Properties != nil {
		var err error
		propertiesAsProto, err = PalettePropertiesAsProto(*data.Properties)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert properties to proto")
		}
	}

	return &BlockPaletteData{
		Properties: propertiesAsProto,
		States:     protoStates,
	}, nil
}

func PalettePropertiesAsProto(properties palette.BlockPaletteProperties) (*BlockProperties, error) {

	var attachment []Attachment
	if properties.Attachment != nil {

		attachment = make([]Attachment, len(properties.Attachment))

		for index, value := range properties.Attachment {

			attachmentValue, err := FromAttachment(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert attachment to proto")
			}

			attachment[index] = *attachmentValue
		}
	}

	var axis []Axis
	if properties.Axis != nil {

		axis = make([]Axis, len(properties.Axis))

		for index, value := range properties.Axis {

			axisValue, err := FromAxis(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert axis to proto")
			}

			axis[index] = *axisValue
		}
	}

	var east []East
	if properties.East != nil {

		east = make([]East, len(properties.East))

		for index, value := range properties.East {

			eastValue, err := FromEast(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert east to proto")
			}

			east[index] = *eastValue
		}
	}

	var face []Face
	if properties.Face != nil {

		face = make([]Face, len(properties.Face))

		for index, value := range properties.Face {

			faceValue, err := FromFace(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert face to proto")
			}

			face[index] = *faceValue
		}
	}

	var facing []Facing
	if properties.Facing != nil {

		facing = make([]Facing, len(properties.Facing))

		for index, value := range properties.Facing {

			facingValue, err := FromFacing(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert facing to proto")
			}

			facing[index] = *facingValue
		}
	}

	var half []Half
	if properties.Half != nil {

		half = make([]Half, len(properties.Half))

		for index, value := range properties.Half {

			halfValue, err := FromHalf(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert half to proto")
			}

			half[index] = *halfValue
		}
	}

	var hinge []Hinge
	if properties.Hinge != nil {

		hinge = make([]Hinge, len(properties.Hinge))

		for index, value := range properties.Hinge {

			hingeValue, err := FromHinge(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert hinge to proto")
			}

			hinge[index] = *hingeValue
		}
	}

	var instrument []Instrument
	if properties.Instrument != nil {

		instrument = make([]Instrument, len(properties.Instrument))

		for index, value := range properties.Instrument {

			instrumentValue, err := FromInstrument(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert instrument to proto")
			}

			instrument[index] = *instrumentValue
		}
	}

	var leaves []Leaves
	if properties.Leaves != nil {

		leaves = make([]Leaves, len(properties.Leaves))

		for index, value := range properties.Leaves {

			leavesValue, err := FromLeaves(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert leaves to proto")
			}

			leaves[index] = *leavesValue
		}
	}

	var mode []Mode
	if properties.Mode != nil {

		mode = make([]Mode, len(properties.Mode))

		for index, value := range properties.Mode {

			modeValue, err := FromMode(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert mode to proto")
			}

			mode[index] = *modeValue
		}
	}

	var north []North
	if properties.North != nil {

		north = make([]North, len(properties.North))

		for index, value := range properties.North {

			northValue, err := FromNorth(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert north to proto")
			}

			north[index] = *northValue
		}
	}

	var orientation []Orientation
	if properties.Orientation != nil {

		orientation = make([]Orientation, len(properties.Orientation))

		for index, value := range properties.Orientation {

			orientationValue, err := FromOrientation(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert orientation to proto")
			}

			orientation[index] = *orientationValue
		}
	}

	var part []Part
	if properties.Part != nil {

		part = make([]Part, len(properties.Part))

		for index, value := range properties.Part {

			partValue, err := FromPart(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert part to proto")
			}

			part[index] = *partValue
		}
	}

	var sculkSensorPhase []SculkSensorPhase
	if properties.SculkSensorPhase != nil {

		sculkSensorPhase = make([]SculkSensorPhase, len(properties.SculkSensorPhase))

		for index, value := range properties.SculkSensorPhase {

			sculkSensorPhaseValue, err := FromSculkSensorPhase(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert sculk sensor phase to proto")
			}

			sculkSensorPhase[index] = *sculkSensorPhaseValue
		}
	}

	var shape []Shape
	if properties.Shape != nil {

		shape = make([]Shape, len(properties.Shape))

		for index, value := range properties.Shape {

			shapeValue, err := FromShape(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert shape to proto")
			}

			shape[index] = *shapeValue
		}
	}

	var south []South
	if properties.South != nil {

		south = make([]South, len(properties.South))

		for index, value := range properties.South {

			southValue, err := FromSouth(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert south to proto")
			}

			south[index] = *southValue
		}
	}

	var thickness []Thickness
	if properties.Thickness != nil {

		thickness = make([]Thickness, len(properties.Thickness))

		for index, value := range properties.Thickness {

			thicknessValue, err := FromThickness(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert thickness to proto")
			}

			thickness[index] = *thicknessValue
		}
	}

	var tilt []Tilt
	if properties.Tilt != nil {

		tilt = make([]Tilt, len(properties.Tilt))

		for index, value := range properties.Tilt {

			tiltValue, err := FromTilt(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert tilt to proto")
			}

			tilt[index] = *tiltValue
		}
	}

	var trialSpawnerState []TrialSpawnerState
	if properties.TrialSpawnerState != nil {

		trialSpawnerState = make([]TrialSpawnerState, len(properties.TrialSpawnerState))

		for index, value := range properties.TrialSpawnerState {

			trialSpawnerStateValue, err := FromTrialSpawnerState(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert trial spawner state to proto")
			}

			trialSpawnerState[index] = *trialSpawnerStateValue
		}
	}

	var typeProperty []Type
	if properties.Type != nil {

		typeProperty = make([]Type, len(properties.Type))

		for index, value := range properties.Type {

			typePropertyValue, err := FromType(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert type to proto")
			}

			typeProperty[index] = *typePropertyValue
		}
	}

	var verticalDirection []VerticalDirection
	if properties.VerticalDirection != nil {

		verticalDirection = make([]VerticalDirection, len(properties.VerticalDirection))

		for index, value := range properties.VerticalDirection {

			verticalDirectionValue, err := FromVerticalDirection(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert vertical direction to proto")
			}

			verticalDirection[index] = *verticalDirectionValue
		}
	}

	var west []West
	if properties.West != nil {

		west = make([]West, len(properties.West))

		for index, value := range properties.West {

			westValue, err := FromWest(value)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert west to proto")
			}

			west[index] = *westValue
		}
	}

	return &BlockProperties{
		Age:               properties.Age,
		Attached:          properties.Attached,
		Attachment:        attachment,
		Axis:              axis,
		Berries:           properties.Berries,
		Bites:             properties.Bites,
		Bottom:            properties.Bottom,
		Bloom:             properties.Bloom,
		CanSummon:         properties.CanSummon,
		Candles:           properties.Candles,
		Conditional:       properties.Conditional,
		Cracked:           properties.Cracked,
		Crafting:          properties.Crafting,
		Charges:           properties.Charges,
		Delay:             properties.Delay,
		Disarmed:          properties.Disarmed,
		Distance:          properties.Distance,
		Down:              properties.Down,
		Drag:              properties.Drag,
		Dusted:            properties.Dusted,
		East:              east,
		Eggs:              properties.Eggs,
		Eye:               properties.Eye,
		Enabled:           properties.Enabled,
		Extended:          properties.Extended,
		Face:              face,
		Facing:            facing,
		FlowerAmount:      properties.FlowerAmount,
		Half:              half,
		Hanging:           properties.Hanging,
		HasBook:           properties.HasBook,
		HasBottle_0:       properties.HasBottle0,
		HasBottle_1:       properties.HasBottle1,
		HasBottle_2:       properties.HasBottle2,
		HasRecord:         properties.HasRecord,
		Hatch:             properties.Hatch,
		Hinge:             hinge,
		HoneyLevel:        properties.HoneyLevel,
		Instrument:        instrument,
		Inverted:          properties.Inverted,
		Layers:            properties.Layers,
		Level:             properties.Level,
		Leaves:            leaves,
		Lit:               properties.Lit,
		Locked:            properties.Locked,
		Mode:              mode,
		Moisture:          properties.Moisture,
		North:             north,
		Orientation:       orientation,
		Note:              properties.Note,
		Occupied:          properties.Occupied,
		InWall:            properties.InWall,
		Open:              properties.Open,
		Persistent:        properties.Persistent,
		Pickles:           properties.Pickles,
		Power:             properties.Power,
		Powered:           properties.Powered,
		Rotation:          properties.Rotation,
		Part:              part,
		SculkSensorPhase:  sculkSensorPhase,
		Shape:             shape,
		Short:             properties.Short,
		Shrieking:         properties.Shrieking,
		SignalFire:        properties.SignalFire,
		Slot_0Occupied:    properties.Slot0Occupied,
		Slot_1Occupied:    properties.Slot1Occupied,
		Slot_2Occupied:    properties.Slot2Occupied,
		Slot_3Occupied:    properties.Slot3Occupied,
		Slot_4Occupied:    properties.Slot4Occupied,
		Slot_5Occupied:    properties.Slot5Occupied,
		Snowy:             properties.Snowy,
		South:             south,
		Stage:             properties.Stage,
		Thickness:         thickness,
		Tilt:              tilt,
		Triggered:         properties.Triggered,
		TrialSpawnerState: trialSpawnerState,
		Type:              typeProperty,
		Unstable:          properties.Unstable,
		Up:                properties.Up,
		VerticalDirection: verticalDirection,
		Waterlogged:       properties.Waterlogged,
		West:              west,
	}, nil
}

func FromBlockPaletteState(state palette.BlockPaletteState) (*BlockPaletteState, error) {

	var propertiesAsProto *BlockStateProperties

	if state.Properties != nil {
		var err error
		propertiesAsProto, err = FromBlockPaletteStateProperties(*state.Properties)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert properties to proto")
		}
	}

	return &BlockPaletteState{
		Default:    state.Default,
		Id:         uint32(state.ID),
		Properties: propertiesAsProto,
	}, nil
}

func FromBlockPaletteStateProperties(properties palette.BlockPaletteStateProperties) (*BlockStateProperties, error) {

	var attachment *Attachment
	if properties.Attachment != nil {

		asProto, err := FromAttachment(*properties.Attachment)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert attachment to proto")
		}

		attachment = asProto
	}

	var axis *Axis
	if properties.Axis != nil {

		asProto, err := FromAxis(*properties.Axis)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert axis to proto")
		}

		axis = asProto
	}

	var east *East
	if properties.East != nil {

		asProto, err := FromEast(*properties.East)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert east to proto")
		}

		east = asProto
	}

	var face *Face
	if properties.Face != nil {

		asProto, err := FromFace(*properties.Face)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert face to proto")
		}

		face = asProto
	}

	var facing *Facing
	if properties.Facing != nil {

		asProto, err := FromFacing(*properties.Facing)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert facing to proto")
		}

		facing = asProto
	}

	var half *Half
	if properties.Half != nil {

		asProto, err := FromHalf(*properties.Half)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert half to proto")
		}

		half = asProto
	}

	var hinge *Hinge
	if properties.Hinge != nil {

		asProto, err := FromHinge(*properties.Hinge)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert hinge to proto")
		}

		hinge = asProto
	}

	var leaves *Leaves
	if properties.Leaves != nil {

		asProto, err := FromLeaves(*properties.Leaves)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert leaves to proto")
		}

		leaves = asProto
	}

	var instrument *Instrument
	if properties.Instrument != nil {

		asProto, err := FromInstrument(*properties.Instrument)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert instrument to proto")
		}

		instrument = asProto
	}

	var mode *Mode
	if properties.Mode != nil {

		asProto, err := FromMode(*properties.Mode)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert mode to proto")
		}

		mode = asProto
	}

	var north *North
	if properties.North != nil {

		asProto, err := FromNorth(*properties.North)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert north to proto")
		}

		north = asProto
	}

	var orientation *Orientation
	if properties.Orientation != nil {

		asProto, err := FromOrientation(*properties.Orientation)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert orientation to proto")
		}

		orientation = asProto
	}

	var part *Part
	if properties.Part != nil {

		asProto, err := FromPart(*properties.Part)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert part to proto")
		}

		part = asProto
	}

	var sculkSensorPhase *SculkSensorPhase
	if properties.SculkSensorPhase != nil {

		asProto, err := FromSculkSensorPhase(*properties.SculkSensorPhase)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert sculk sensor phase to proto")
		}

		sculkSensorPhase = asProto
	}

	var shape *Shape
	if properties.Shape != nil {

		asProto, err := FromShape(*properties.Shape)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert shape to proto")
		}

		shape = asProto
	}

	var south *South
	if properties.South != nil {

		asProto, err := FromSouth(*properties.South)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert south to proto")
		}

		south = asProto
	}

	var thickness *Thickness
	if properties.Thickness != nil {

		asProto, err := FromThickness(*properties.Thickness)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert thickness to proto")
		}

		thickness = asProto
	}

	var tilt *Tilt
	if properties.Tilt != nil {

		asProto, err := FromTilt(*properties.Tilt)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert tilt to proto")
		}

		tilt = asProto
	}

	var trialSpawnerState *TrialSpawnerState
	if properties.TrialSpawnerState != nil {

		asProto, err := FromTrialSpawnerState(*properties.TrialSpawnerState)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert trial spawner state to proto")
		}

		trialSpawnerState = asProto
	}

	var typeProperty *Type
	if properties.Type != nil {

		asProto, err := FromType(*properties.Type)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert type to proto")
		}

		typeProperty = asProto
	}

	var verticalDirection *VerticalDirection
	if properties.VerticalDirection != nil {

		asProto, err := FromVerticalDirection(*properties.VerticalDirection)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert vertical direction to proto")
		}

		verticalDirection = asProto
	}

	var west *West
	if properties.West != nil {

		asProto, err := FromWest(*properties.West)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert west to proto")
		}

		west = asProto
	}

	return &BlockStateProperties{
		Age:               properties.Age,
		Attached:          properties.Attached,
		Attachment:        attachment,
		Axis:              axis,
		Berries:           properties.Berries,
		Bites:             properties.Bites,
		Bottom:            properties.Bottom,
		Bloom:             properties.Bloom,
		CanSummon:         properties.CanSummon,
		Candles:           properties.Candles,
		Conditional:       properties.Conditional,
		Cracked:           properties.Cracked,
		Crafting:          properties.Crafting,
		Charges:           properties.Charges,
		Delay:             properties.Delay,
		Disarmed:          properties.Disarmed,
		Distance:          properties.Distance,
		Down:              properties.Down,
		Drag:              properties.Drag,
		Dusted:            properties.Dusted,
		East:              east,
		Eggs:              properties.Eggs,
		Eye:               properties.Eye,
		Enabled:           properties.Enabled,
		Extended:          properties.Extended,
		Face:              face,
		Facing:            facing,
		FlowerAmount:      properties.FlowerAmount,
		Half:              half,
		Hanging:           properties.Hanging,
		HasBook:           properties.HasBook,
		HasBottle_0:       properties.HasBottle0,
		HasBottle_1:       properties.HasBottle1,
		HasBottle_2:       properties.HasBottle2,
		HasRecord:         properties.HasRecord,
		Hatch:             properties.Hatch,
		Hinge:             hinge,
		HoneyLevel:        properties.HoneyLevel,
		Instrument:        instrument,
		Inverted:          properties.Inverted,
		Layers:            properties.Layers,
		Level:             properties.Level,
		Leaves:            leaves,
		Lit:               properties.Lit,
		Locked:            properties.Locked,
		Mode:              mode,
		Moisture:          properties.Moisture,
		North:             north,
		Orientation:       orientation,
		Part:              part,
		Note:              properties.Note,
		Occupied:          properties.Occupied,
		InWall:            properties.InWall,
		Open:              properties.Open,
		Persistent:        properties.Persistent,
		Pickles:           properties.Pickles,
		Power:             properties.Power,
		Powered:           properties.Powered,
		Rotation:          properties.Rotation,
		SculkSensorPhase:  sculkSensorPhase,
		Shape:             shape,
		Short:             properties.Short,
		Shrieking:         properties.Shrieking,
		SignalFire:        properties.SignalFire,
		Slot_0Occupied:    properties.Slot0Occupied,
		Slot_1Occupied:    properties.Slot1Occupied,
		Slot_2Occupied:    properties.Slot2Occupied,
		Slot_3Occupied:    properties.Slot3Occupied,
		Slot_4Occupied:    properties.Slot4Occupied,
		Slot_5Occupied:    properties.Slot5Occupied,
		Snowy:             properties.Snowy,
		South:             south,
		Stage:             properties.Stage,
		Thickness:         thickness,
		Tilt:              tilt,
		TrialSpawnerState: trialSpawnerState,
		Triggered:         properties.Triggered,
		Type:              typeProperty,
		Unstable:          properties.Unstable,
		Up:                properties.Up,
		VerticalDirection: verticalDirection,
		Waterlogged:       properties.Waterlogged,
		West:              west,
	}, nil
}

func (p *BlockPalette) AsJSONStruct() (palette.BlockPaletteJSON, error) {

	paletteResult := make(palette.BlockPaletteJSON, len(p.Entries))

	for _, entry := range p.Entries {

		materialName := fmt.Sprintf("minecraft:%s", strings.ToLower(entry.Material.String()))

		states := make([]palette.BlockPaletteState, len(entry.Data.States))
		for i, state := range entry.Data.States {
			stateAsJSONStruct, err := state.AsJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert state to json struct")
			}
			states[i] = *stateAsJSONStruct
		}

		properties, err := entry.Data.Properties.AsJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert properties to json struct")
		}

		paletteResult[materialName] = palette.BlockPaletteData{
			Properties: properties,
			States:     states,
		}
	}

	return paletteResult, nil
}

func (p *BlockProperties) AsJSONStruct() (*palette.BlockPaletteProperties, error) {

	if p == nil {
		return nil, nil
	}

	var attachment []palette.Attachment
	if p.Attachment != nil {
		attachment = make([]palette.Attachment, len(p.Attachment))
		for i, attachmentValue := range p.Attachment {
			attachmentResult, err := attachmentValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert attachment to json struct")
			}
			attachment[i] = *attachmentResult
		}
	}

	var axis []palette.Axis
	if p.Axis != nil {
		axis = make([]palette.Axis, len(p.Axis))
		for i, axisValue := range p.Axis {
			axisResult, err := axisValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert axis to json struct")
			}
			axis[i] = *axisResult
		}
	}

	var east []palette.East
	if p.East != nil {
		east = make([]palette.East, len(p.East))
		for i, eastValue := range p.East {
			eastResult, err := eastValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert east to json struct")
			}
			east[i] = *eastResult
		}
	}

	var face []palette.Face
	if p.Face != nil {
		face = make([]palette.Face, len(p.Face))
		for i, faceValue := range p.Face {
			faceResult, err := faceValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert face to json struct")
			}
			face[i] = *faceResult
		}
	}

	var facing []palette.Facing
	if p.Facing != nil {
		facing = make([]palette.Facing, len(p.Facing))
		for i, facingValue := range p.Facing {
			facingResult, err := facingValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert facing to json struct")
			}
			facing[i] = *facingResult
		}
	}

	var half []palette.Half
	if p.Half != nil {
		half = make([]palette.Half, len(p.Half))
		for i, halfValue := range p.Half {
			halfResult, err := halfValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert half to json struct")
			}
			half[i] = *halfResult
		}
	}

	var hinge []palette.Hinge
	if p.Hinge != nil {
		hinge = make([]palette.Hinge, len(p.Hinge))
		for i, hingeValue := range p.Hinge {
			hingeResult, err := hingeValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert hinge to json struct")
			}
			hinge[i] = *hingeResult
		}
	}

	var instrument []palette.Instrument
	if p.Instrument != nil {
		instrument = make([]palette.Instrument, len(p.Instrument))
		for i, instrumentValue := range p.Instrument {
			instrumentResult, err := instrumentValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert instrument to json struct")
			}
			instrument[i] = *instrumentResult
		}
	}

	var leaves []palette.Leaves
	if p.Leaves != nil {
		leaves = make([]palette.Leaves, len(p.Leaves))
		for i, leavesValue := range p.Leaves {
			leavesResult, err := leavesValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert leaves to json struct")
			}
			leaves[i] = *leavesResult
		}
	}

	var mode []palette.Mode
	if p.Mode != nil {
		mode = make([]palette.Mode, len(p.Mode))
		for i, modeValue := range p.Mode {
			modeResult, err := modeValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert mode to json struct")
			}
			mode[i] = *modeResult
		}
	}

	var north []palette.North
	if p.North != nil {
		north = make([]palette.North, len(p.North))
		for i, northValue := range p.North {
			northResult, err := northValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert north to json struct")
			}
			north[i] = *northResult
		}
	}

	var orientation []palette.Orientation
	if p.Orientation != nil {
		orientation = make([]palette.Orientation, len(p.Orientation))
		for i, orientationValue := range p.Orientation {
			orientationResult, err := orientationValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert orientation to json struct")
			}
			orientation[i] = *orientationResult
		}
	}

	var part []palette.Part
	if p.Part != nil {
		part = make([]palette.Part, len(p.Part))
		for i, partValue := range p.Part {
			partResult, err := partValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert part to json struct")
			}
			part[i] = *partResult
		}
	}

	var sculkSensorPhase []palette.SculkSensorPhase
	if p.SculkSensorPhase != nil {
		sculkSensorPhase = make([]palette.SculkSensorPhase, len(p.SculkSensorPhase))
		for i, sculkSensorPhaseValue := range p.SculkSensorPhase {
			sculkSensorPhaseResult, err := sculkSensorPhaseValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert sculkSensorPhase to json struct")
			}
			sculkSensorPhase[i] = *sculkSensorPhaseResult
		}
	}

	var shape []palette.Shape
	if p.Shape != nil {
		shape = make([]palette.Shape, len(p.Shape))
		for i, shapeValue := range p.Shape {
			shapeResult, err := shapeValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert shape to json struct")
			}
			shape[i] = *shapeResult
		}
	}

	var south []palette.South
	if p.South != nil {
		south = make([]palette.South, len(p.South))
		for i, southValue := range p.South {
			southResult, err := southValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert south to json struct")
			}
			south[i] = *southResult
		}
	}

	var thickness []palette.Thickness
	if p.Thickness != nil {
		thickness = make([]palette.Thickness, len(p.Thickness))
		for i, thicknessValue := range p.Thickness {
			thicknessResult, err := thicknessValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert thickness to json struct")
			}
			thickness[i] = *thicknessResult
		}
	}

	var tilt []palette.Tilt
	if p.Tilt != nil {
		tilt = make([]palette.Tilt, len(p.Tilt))
		for i, tiltValue := range p.Tilt {
			tiltResult, err := tiltValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert tilt to json struct")
			}
			tilt[i] = *tiltResult
		}
	}

	var trialSpawnerState []palette.TrialSpawnerState
	if p.TrialSpawnerState != nil {
		trialSpawnerState = make([]palette.TrialSpawnerState, len(p.TrialSpawnerState))
		for i, TrialSpawnerStateValue := range p.TrialSpawnerState {
			TrialSpawnerStateResult, err := TrialSpawnerStateValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert TrialSpawnerState to json struct")
			}
			trialSpawnerState[i] = *TrialSpawnerStateResult
		}
	}

	var typeProperty []palette.Type
	if p.Type != nil {
		typeProperty = make([]palette.Type, len(p.Type))
		for i, typePropertyValue := range p.Type {
			typePropertyResult, err := typePropertyValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert type to json struct")
			}
			typeProperty[i] = *typePropertyResult
		}
	}

	var verticalDirection []palette.VerticalDirection
	if p.VerticalDirection != nil {
		verticalDirection = make([]palette.VerticalDirection, len(p.VerticalDirection))
		for i, verticalDirectionValue := range p.VerticalDirection {
			verticalDirectionResult, err := verticalDirectionValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert verticalDirection to json struct")
			}
			verticalDirection[i] = *verticalDirectionResult
		}
	}

	var west []palette.West
	if p.West != nil {
		west = make([]palette.West, len(p.West))
		for i, westValue := range p.West {
			westResult, err := westValue.asJSONStruct()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to convert west to json struct")
			}
			west[i] = *westResult
		}
	}

	return &palette.BlockPaletteProperties{
		Age:               p.Age,
		Attached:          p.Attached,
		Attachment:        attachment,
		Axis:              axis,
		Berries:           p.Berries,
		Bites:             p.Bites,
		Bottom:            p.Bottom,
		Bloom:             p.Bloom,
		CanSummon:         p.CanSummon,
		Candles:           p.Candles,
		Conditional:       p.Conditional,
		Cracked:           p.Cracked,
		Crafting:          p.Crafting,
		Charges:           p.Charges,
		Delay:             p.Delay,
		Disarmed:          p.Disarmed,
		Distance:          p.Distance,
		Down:              p.Down,
		Drag:              p.Drag,
		Dusted:            p.Dusted,
		East:              east,
		Eggs:              p.Eggs,
		Eye:               p.Eye,
		Enabled:           p.Enabled,
		Extended:          p.Extended,
		Face:              face,
		Facing:            facing,
		FlowerAmount:      p.FlowerAmount,
		Half:              half,
		Hanging:           p.Hanging,
		HasBook:           p.HasBook,
		HasBottle0:        p.HasBottle_0,
		HasBottle1:        p.HasBottle_1,
		HasBottle2:        p.HasBottle_2,
		HasRecord:         p.HasRecord,
		Hatch:             p.Hatch,
		Hinge:             hinge,
		HoneyLevel:        p.HoneyLevel,
		Instrument:        instrument,
		Inverted:          p.Inverted,
		Layers:            p.Layers,
		Level:             p.Level,
		Leaves:            leaves,
		Lit:               p.Lit,
		Locked:            p.Locked,
		Mode:              mode,
		Moisture:          p.Moisture,
		North:             north,
		Note:              p.Note,
		Occupied:          p.Occupied,
		InWall:            p.InWall,
		Open:              p.Open,
		Orientation:       orientation,
		Part:              part,
		Persistent:        p.Persistent,
		Pickles:           p.Pickles,
		Power:             p.Power,
		Powered:           p.Powered,
		Rotation:          p.Rotation,
		SculkSensorPhase:  sculkSensorPhase,
		Shape:             shape,
		Short:             p.Short,
		Shrieking:         p.Shrieking,
		SignalFire:        p.SignalFire,
		Slot0Occupied:     p.Slot_0Occupied,
		Slot1Occupied:     p.Slot_1Occupied,
		Slot2Occupied:     p.Slot_2Occupied,
		Slot3Occupied:     p.Slot_3Occupied,
		Slot4Occupied:     p.Slot_4Occupied,
		Slot5Occupied:     p.Slot_5Occupied,
		Snowy:             p.Snowy,
		South:             south,
		Stage:             p.Stage,
		Thickness:         thickness,
		Tilt:              tilt,
		TrialSpawnerState: trialSpawnerState,
		Triggered:         p.Triggered,
		Type:              typeProperty,
		Unstable:          p.Unstable,
		Up:                p.Up,
		VerticalDirection: verticalDirection,
		Waterlogged:       p.Waterlogged,
		West:              west,
	}, nil
}

func (s *BlockPaletteState) AsJSONStruct() (*palette.BlockPaletteState, error) {

	properties, err := s.Properties.asJSONStruct()
	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "failed to convert properties to json struct")
	}

	return &palette.BlockPaletteState{
		Default:    s.Default,
		ID:         uint(s.Id),
		Properties: properties,
	}, nil
}

func (s *BlockStateProperties) asJSONStruct() (*palette.BlockPaletteStateProperties, error) {

	if s == nil {
		return nil, nil
	}

	var attachment *palette.Attachment
	if s.Attachment != nil {
		attachmentResult, err := s.Attachment.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert attachment to json struct")
		}
		attachment = attachmentResult
	}

	var axis *palette.Axis
	if s.Axis != nil {
		axisResult, err := s.Axis.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert axis to json struct")
		}
		axis = axisResult
	}

	var east *palette.East
	if s.East != nil {
		eastResult, err := s.East.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert east to json struct")
		}
		east = eastResult
	}

	var face *palette.Face
	if s.Face != nil {
		faceResult, err := s.Face.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert face to json struct")
		}
		face = faceResult
	}

	var facing *palette.Facing
	if s.Facing != nil {
		facingResult, err := s.Facing.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert facing to json struct")
		}
		facing = facingResult
	}

	var half *palette.Half
	if s.Half != nil {
		halfResult, err := s.Half.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert half to json struct")
		}
		half = halfResult
	}

	var hinge *palette.Hinge
	if s.Hinge != nil {
		hingeResult, err := s.Hinge.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert hinge to json struct")
		}
		hinge = hingeResult
	}

	var instrument *palette.Instrument
	if s.Instrument != nil {
		instrumentResult, err := s.Instrument.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert instrument to json struct")
		}
		instrument = instrumentResult

	}

	var leaves *palette.Leaves
	if s.Leaves != nil {
		leavesResult, err := s.Leaves.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert leaves to json struct")
		}
		leaves = leavesResult
	}

	var mode *palette.Mode
	if s.Mode != nil {
		modeResult, err := s.Mode.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert mode to json struct")
		}
		mode = modeResult
	}

	var north *palette.North
	if s.North != nil {
		northResult, err := s.North.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert north to json struct")
		}
		north = northResult
	}

	var orientation *palette.Orientation
	if s.Orientation != nil {
		orientationResult, err := s.Orientation.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert orientation to json struct")
		}
		orientation = orientationResult
	}

	var part *palette.Part
	if s.Part != nil {
		partResult, err := s.Part.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert part to json struct")
		}
		part = partResult
	}

	var sculkSensorPhase *palette.SculkSensorPhase
	if s.SculkSensorPhase != nil {
		sculkSensorPhaseResult, err := s.SculkSensorPhase.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert sculkSensorPhase to json struct")
		}
		sculkSensorPhase = sculkSensorPhaseResult
	}

	var shape *palette.Shape
	if s.Shape != nil {
		shapeResult, err := s.Shape.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert shape to json struct")
		}
		shape = shapeResult
	}

	var south *palette.South
	if s.South != nil {
		southResult, err := s.South.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert south to json struct")
		}
		south = southResult
	}

	var thickness *palette.Thickness
	if s.Thickness != nil {
		thicknessResult, err := s.Thickness.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert thickness to json struct")
		}
		thickness = thicknessResult
	}

	var tilt *palette.Tilt
	if s.Tilt != nil {
		tiltResult, err := s.Tilt.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert tilt to json struct")
		}
		tilt = tiltResult
	}

	var trialSpawnerState *palette.TrialSpawnerState
	if s.TrialSpawnerState != nil {
		trialSpawnerStateResult, err := s.TrialSpawnerState.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert trialSpawnerState to json struct")
		}
		trialSpawnerState = trialSpawnerStateResult
	}

	var typeProperty *palette.Type
	if s.Type != nil {
		typePropertyResult, err := s.Type.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert type to json struct")
		}
		typeProperty = typePropertyResult
	}

	var verticalDirection *palette.VerticalDirection
	if s.VerticalDirection != nil {
		verticalDirectionResult, err := s.VerticalDirection.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert verticalDirection to json struct")
		}
		verticalDirection = verticalDirectionResult
	}

	var west *palette.West
	if s.West != nil {
		westResult, err := s.West.asJSONStruct()
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "failed to convert west to json struct")
		}
		west = westResult
	}

	return &palette.BlockPaletteStateProperties{
		Age:               s.Age,
		Attached:          s.Attached,
		Attachment:        attachment,
		Axis:              axis,
		Berries:           s.Berries,
		Bites:             s.Bites,
		Bottom:            s.Bottom,
		Bloom:             s.Bloom,
		CanSummon:         s.CanSummon,
		Candles:           s.Candles,
		Conditional:       s.Conditional,
		Cracked:           s.Cracked,
		Crafting:          s.Crafting,
		Charges:           s.Charges,
		Delay:             s.Delay,
		Disarmed:          s.Disarmed,
		Distance:          s.Distance,
		Down:              s.Down,
		Drag:              s.Drag,
		Dusted:            s.Dusted,
		East:              east,
		Eggs:              s.Eggs,
		Eye:               s.Eye,
		Enabled:           s.Enabled,
		Extended:          s.Extended,
		Face:              face,
		Facing:            facing,
		FlowerAmount:      s.FlowerAmount,
		Half:              half,
		Hanging:           s.Hanging,
		HasBook:           s.HasBook,
		HasBottle0:        s.HasBottle_0,
		HasBottle1:        s.HasBottle_1,
		HasBottle2:        s.HasBottle_2,
		HasRecord:         s.HasRecord,
		Hatch:             s.Hatch,
		Hinge:             hinge,
		HoneyLevel:        s.HoneyLevel,
		Instrument:        instrument,
		Inverted:          s.Inverted,
		Layers:            s.Layers,
		Level:             s.Level,
		Leaves:            leaves,
		Lit:               s.Lit,
		Locked:            s.Locked,
		Mode:              mode,
		Moisture:          s.Moisture,
		North:             north,
		Note:              s.Note,
		Occupied:          s.Occupied,
		InWall:            s.InWall,
		Open:              s.Open,
		Orientation:       orientation,
		Part:              part,
		Persistent:        s.Persistent,
		Pickles:           s.Pickles,
		Power:             s.Power,
		Powered:           s.Powered,
		Rotation:          s.Rotation,
		SculkSensorPhase:  sculkSensorPhase,
		Shape:             shape,
		Short:             s.Short,
		Shrieking:         s.Shrieking,
		SignalFire:        s.SignalFire,
		Slot0Occupied:     s.Slot_0Occupied,
		Slot1Occupied:     s.Slot_1Occupied,
		Slot2Occupied:     s.Slot_2Occupied,
		Slot3Occupied:     s.Slot_3Occupied,
		Slot4Occupied:     s.Slot_4Occupied,
		Slot5Occupied:     s.Slot_5Occupied,
		Snowy:             s.Snowy,
		South:             south,
		Stage:             s.Stage,
		Thickness:         thickness,
		Tilt:              tilt,
		TrialSpawnerState: trialSpawnerState,
		Triggered:         s.Triggered,
		Type:              typeProperty,
		Unstable:          s.Unstable,
		Up:                s.Up,
		VerticalDirection: verticalDirection,
		Waterlogged:       s.Waterlogged,
		West:              west,
	}, nil
}

func FromAttachment(attachment palette.Attachment) (*Attachment, error) {

	var output Attachment

	switch attachment {
	case palette.AttachmentFloor:
		output = Attachment_ATTACHMENT_FLOOR
	case palette.AttachmentCeiling:
		output = Attachment_ATTACHMENT_CEILING
	case palette.AttachmentSingleWall:
		output = Attachment_ATTACHMENT_SINGLE_WALL
	case palette.AttachmentDoubleWall:
		output = Attachment_ATTACHMENT_DOUBLE_WALL
	default:
		return nil, errorx.IllegalState.New("unknown attachment: %s", attachment)
	}

	return &output, nil
}

func FromAxis(axis palette.Axis) (*Axis, error) {

	var output Axis

	switch axis {
	case palette.AxisX:
		output = Axis_AXIS_X
	case palette.AxisY:
		output = Axis_AXIS_Y
	case palette.AxisZ:
		output = Axis_AXIS_Z
	default:
		return nil, errorx.IllegalState.New("unknown axis: %s", axis)
	}

	return &output, nil
}

func FromEast(east palette.East) (*East, error) {

	var output East

	switch east {
	case palette.EastFalse:
		output = East_EAST_FALSE
	case palette.EastTrue:
		output = East_EAST_TRUE
	case palette.EastNone:
		output = East_EAST_NONE
	case palette.EastLow:
		output = East_EAST_LOW
	case palette.EastTall:
		output = East_EAST_TALL
	case palette.EastUp:
		output = East_EAST_UP
	case palette.EastSide:
		output = East_EAST_SIDE
	default:
		return nil, errorx.IllegalState.New("unknown east: %s", east)
	}

	return &output, nil
}

func FromFace(face palette.Face) (*Face, error) {

	var output Face

	switch face {
	case palette.FaceFloor:
		output = Face_FACE_FLOOR
	case palette.FaceCeiling:
		output = Face_FACE_CEILING
	case palette.FaceWall:
		output = Face_FACE_WALL
	default:
		return nil, errorx.IllegalState.New("unknown face: %s", face)
	}

	return &output, nil
}

func FromFacing(facing palette.Facing) (*Facing, error) {

	var output Facing

	switch facing {
	case palette.FacingNorth:
		output = Facing_FACING_NORTH
	case palette.FacingSouth:
		output = Facing_FACING_SOUTH
	case palette.FacingWest:
		output = Facing_FACING_WEST
	case palette.FacingEast:
		output = Facing_FACING_EAST
	case palette.FacingDown:
		output = Facing_FACING_DOWN
	case palette.FacingUp:
		output = Facing_FACING_UP
	default:
		return nil, errorx.IllegalState.New("unknown facing: %s", facing)
	}

	return &output, nil
}

func FromHalf(half palette.Half) (*Half, error) {

	var output Half

	switch half {
	case palette.HalfTop:
		output = Half_HALF_TOP
	case palette.HalfBottom:
		output = Half_HALF_BOTTOM
	case palette.HalfUpper:
		output = Half_HALF_UPPER
	case palette.HalfLower:
		output = Half_HALF_LOWER
	default:
		return nil, errorx.IllegalState.New("unknown half: %s", half)
	}

	return &output, nil
}

func FromHinge(hinge palette.Hinge) (*Hinge, error) {

	var output Hinge

	switch hinge {
	case palette.HingeLeft:
		output = Hinge_HINGE_LEFT
	case palette.HingeRight:
		output = Hinge_HINGE_RIGHT
	default:
		return nil, errorx.IllegalState.New("unknown hinge: %s", hinge)
	}

	return &output, nil
}

func FromInstrument(instrument palette.Instrument) (*Instrument, error) {

	var output Instrument

	switch instrument {
	case palette.InstrumentHarp:
		output = Instrument_INSTRUMENT_HARP
	case palette.InstrumentBasedrum:
		output = Instrument_INSTRUMENT_BASEDRUM
	case palette.InstrumentSnare:
		output = Instrument_INSTRUMENT_SNARE
	case palette.InstrumentHat:
		output = Instrument_INSTRUMENT_HAT
	case palette.InstrumentBass:
		output = Instrument_INSTRUMENT_BASS
	case palette.InstrumentFlute:
		output = Instrument_INSTRUMENT_FLUTE
	case palette.InstrumentBell:
		output = Instrument_INSTRUMENT_BELL
	case palette.InstrumentGuitar:
		output = Instrument_INSTRUMENT_GUITAR
	case palette.InstrumentChime:
		output = Instrument_INSTRUMENT_CHIME
	case palette.InstrumentXylophone:
		output = Instrument_INSTRUMENT_XYLOPHONE
	case palette.InstrumentIronXylophone:
		output = Instrument_INSTRUMENT_IRON_XYLOPHONE
	case palette.InstrumentCowBell:
		output = Instrument_INSTRUMENT_COW_BELL
	case palette.InstrumentDidgeridoo:
		output = Instrument_INSTRUMENT_DIDGERIDOO
	case palette.InstrumentBit:
		output = Instrument_INSTRUMENT_BIT
	case palette.InstrumentBanjo:
		output = Instrument_INSTRUMENT_BANJO
	case palette.InstrumentPling:
		output = Instrument_INSTRUMENT_PLING
	case palette.InstrumentZombie:
		output = Instrument_INSTRUMENT_ZOMBIE
	case palette.InstrumentSkeleton:
		output = Instrument_INSTRUMENT_SKELETON
	case palette.InstrumentCreeper:
		output = Instrument_INSTRUMENT_CREEPER
	case palette.InstrumentDragon:
		output = Instrument_INSTRUMENT_DRAGON
	case palette.InstrumentWitherSkeleton:
		output = Instrument_INSTRUMENT_WITHER_SKELETON
	case palette.InstrumentPiglin:
		output = Instrument_INSTRUMENT_PIGLIN
	case palette.InstrumentCustomHead:
		output = Instrument_INSTRUMENT_CUSTOM_HEAD
	default:
		return nil, errorx.IllegalState.New("unknown instrument: %s", instrument)
	}

	return &output, nil
}

func FromLeaves(leaves palette.Leaves) (*Leaves, error) {

	var output Leaves

	switch leaves {
	case palette.LeavesNone:
		output = Leaves_LEAVES_NONE
	case palette.LeavesSmall:
		output = Leaves_LEAVES_SMALL
	case palette.LeavesLarge:
		output = Leaves_LEAVES_LARGE
	default:
		return nil, errorx.IllegalState.New("unknown leaves: %s", leaves)
	}

	return &output, nil
}

func FromMode(mode palette.Mode) (*Mode, error) {

	var output Mode

	switch mode {
	case palette.ModeSave:
		output = Mode_MODE_SAVE
	case palette.ModeLoad:
		output = Mode_MODE_LOAD
	case palette.ModeCorner:
		output = Mode_MODE_CORNER
	case palette.ModeData:
		output = Mode_MODE_DATA
	case palette.ModeCompare:
		output = Mode_MODE_COMPARE
	case palette.ModeSubtract:
		output = Mode_MODE_SUBTRACT
	default:
		return nil, errorx.IllegalState.New("unknown mode: %s", mode)
	}

	return &output, nil
}

func FromNorth(north palette.North) (*North, error) {

	var output North

	switch north {
	case palette.NorthFalse:
		output = North_NORTH_FALSE
	case palette.NorthTrue:
		output = North_NORTH_TRUE
	case palette.NorthNone:
		output = North_NORTH_NONE
	case palette.NorthLow:
		output = North_NORTH_LOW
	case palette.NorthTall:
		output = North_NORTH_TALL
	case palette.NorthUp:
		output = North_NORTH_UP
	case palette.NorthSide:
		output = North_NORTH_SIDE
	default:
		return nil, errorx.IllegalState.New("unknown north: %s", north)
	}

	return &output, nil
}

func FromOrientation(orientation palette.Orientation) (*Orientation, error) {

	var output Orientation

	switch orientation {
	case palette.OrientationDownEast:
		output = Orientation_ORIENTATION_DOWN_EAST
	case palette.OrientationDownNorth:
		output = Orientation_ORIENTATION_DOWN_NORTH
	case palette.OrientationDownSouth:
		output = Orientation_ORIENTATION_DOWN_SOUTH
	case palette.OrientationDownWest:
		output = Orientation_ORIENTATION_DOWN_WEST
	case palette.OrientationEastUp:
		output = Orientation_ORIENTATION_EAST_UP
	case palette.OrientationNorthUp:
		output = Orientation_ORIENTATION_NORTH_UP
	case palette.OrientationSouthUp:
		output = Orientation_ORIENTATION_SOUTH_UP
	case palette.OrientationWestUp:
		output = Orientation_ORIENTATION_WEST_UP
	case palette.OrientationUpEast:
		output = Orientation_ORIENTATION_UP_EAST
	case palette.OrientationUpNorth:
		output = Orientation_ORIENTATION_UP_NORTH
	case palette.OrientationUpSouth:
		output = Orientation_ORIENTATION_UP_SOUTH
	case palette.OrientationUpWest:
		output = Orientation_ORIENTATION_UP_WEST
	default:
		return nil, errorx.IllegalState.New("unknown orientation: %s", orientation)
	}

	return &output, nil
}

func FromPart(part palette.Part) (*Part, error) {

	var output Part

	switch part {
	case palette.PartHead:
		output = Part_PART_HEAD
	case palette.PartFoot:
		output = Part_PART_FOOT
	default:
		return nil, errorx.IllegalState.New("unknown part: %s", part)
	}

	return &output, nil
}

func FromSculkSensorPhase(sculkSensorPhase palette.SculkSensorPhase) (*SculkSensorPhase, error) {

	var output SculkSensorPhase

	switch sculkSensorPhase {
	case palette.SculkSensorPhaseInactive:
		output = SculkSensorPhase_SCULK_SENSOR_PHASE_INACTIVE
	case palette.SculkSensorPhaseActive:
		output = SculkSensorPhase_SCULK_SENSOR_PHASE_ACTIVE
	case palette.SculkSensorPhaseCooldown:
		output = SculkSensorPhase_SCULK_SENSOR_PHASE_COOLDOWN
	default:
		return nil, errorx.IllegalState.New("unknown sculkSensorPhase: %s", sculkSensorPhase)
	}

	return &output, nil
}

func FromShape(shape palette.Shape) (*Shape, error) {

	var output Shape

	switch shape {
	case palette.ShapeStraight:
		output = Shape_SHAPE_STRAIGHT
	case palette.ShapeInnerLeft:
		output = Shape_SHAPE_INNER_LEFT
	case palette.ShapeInnerRight:
		output = Shape_SHAPE_INNER_RIGHT
	case palette.ShapeOuterLeft:
		output = Shape_SHAPE_OUTER_LEFT
	case palette.ShapeOuterRight:
		output = Shape_SHAPE_OUTER_RIGHT
	case palette.ShapeNorthWest:
		output = Shape_SHAPE_NORTH_WEST
	case palette.ShapeNorthEast:
		output = Shape_SHAPE_NORTH_EAST
	case palette.ShapeSouthEast:
		output = Shape_SHAPE_SOUTH_EAST
	case palette.ShapeSouthWest:
		output = Shape_SHAPE_SOUTH_WEST
	case palette.ShapeNorthSouth:
		output = Shape_SHAPE_NORTH_SOUTH
	case palette.ShapeEastWest:
		output = Shape_SHAPE_EAST_WEST
	case palette.ShapeAscendingNorth:
		output = Shape_SHAPE_ASCENDING_NORTH
	case palette.ShapeAscendingEast:
		output = Shape_SHAPE_ASCENDING_EAST
	case palette.ShapeAscendingSouth:
		output = Shape_SHAPE_ASCENDING_SOUTH
	case palette.ShapeAscendingWest:
		output = Shape_SHAPE_ASCENDING_WEST
	default:
		return nil, errorx.IllegalState.New("unknown shape: %s", shape)
	}

	return &output, nil
}

func FromSouth(south palette.South) (*South, error) {

	var output South

	switch south {
	case palette.SouthFalse:
		output = South_SOUTH_FALSE
	case palette.SouthTrue:
		output = South_SOUTH_TRUE
	case palette.SouthNone:
		output = South_SOUTH_NONE
	case palette.SouthLow:
		output = South_SOUTH_LOW
	case palette.SouthTall:
		output = South_SOUTH_TALL
	case palette.SouthUp:
		output = South_SOUTH_UP
	case palette.SouthSide:
		output = South_SOUTH_SIDE
	default:
		return nil, errorx.IllegalState.New("unknown south: %s", south)
	}

	return &output, nil
}

func FromThickness(thickness palette.Thickness) (*Thickness, error) {

	var output Thickness

	switch thickness {
	case palette.ThicknessTipMerge:
		output = Thickness_THICKNESS_TIP_MERGE
	case palette.ThicknessTip:
		output = Thickness_THICKNESS_TIP
	case palette.ThicknessFrustum:
		output = Thickness_THICKNESS_FRUSTUM
	case palette.ThicknessMiddle:
		output = Thickness_THICKNESS_MIDDLE
	case palette.ThicknessBase:
		output = Thickness_THICKNESS_BASE

	default:
		return nil, errorx.IllegalState.New("unknown thickness: %s", thickness)
	}

	return &output, nil
}

func FromTilt(tilt palette.Tilt) (*Tilt, error) {

	var output Tilt

	switch tilt {
	case palette.TiltNone:
		output = Tilt_TILT_NONE
	case palette.TiltUnstable:
		output = Tilt_TILT_UNSTABLE
	case palette.TiltFull:
		output = Tilt_TILT_FULL
	case palette.TiltPartial:
		output = Tilt_TILT_PARTIAL
	default:
		return nil, errorx.IllegalState.New("unknown tilt: %s", tilt)
	}

	return &output, nil
}

func FromTrialSpawnerState(state palette.TrialSpawnerState) (*TrialSpawnerState, error) {

	var output TrialSpawnerState

	switch state {
	case palette.TrialSpawnerStateInActive:
		output = TrialSpawnerState_TRIAL_SPAWNER_STATE_INACTIVE
	case palette.TrialSpawnerStateWaitingForPlayers:
		output = TrialSpawnerState_TRIAL_SPAWNER_STATE_WAITING_FOR_PLAYERS
	case palette.TrialSpawnerStateActive:
		output = TrialSpawnerState_TRIAL_SPAWNER_STATE_ACTIVE
	case palette.TrialSpawnerStateWaitingForRewardEjection:
		output = TrialSpawnerState_TRIAL_SPAWNER_STATE_WAITING_FOR_REWARD_EJECTION
	case palette.TrialSpawnerStateEjectingReward:
		output = TrialSpawnerState_TRIAL_SPAWNER_STATE_EJECTING_REWARD
	case palette.TrialSpawnerStateCooldown:
		output = TrialSpawnerState_TRIAL_SPAWNER_STATE_COOLDOWN
	default:
		return nil, errorx.IllegalState.New("unknown trial spawner state: %s", state)
	}

	return &output, nil
}

func FromType(typeProperty palette.Type) (*Type, error) {

	var output Type

	switch typeProperty {
	case palette.TypeSingle:
		output = Type_TYPE_SINGLE
	case palette.TypeLeft:
		output = Type_TYPE_LEFT
	case palette.TypeRight:
		output = Type_TYPE_RIGHT
	case palette.TypeTop:
		output = Type_TYPE_TOP
	case palette.TypeBottom:
		output = Type_TYPE_BOTTOM
	case palette.TypeDouble:
		output = Type_TYPE_DOUBLE
	case palette.TypeSticky:
		output = Type_TYPE_STICKY
	case palette.TypeNormal:
		output = Type_TYPE_NORMAL
	default:
		return nil, errorx.IllegalState.New("unknown type: %s", typeProperty)
	}

	return &output, nil
}

func FromVerticalDirection(verticalDirection palette.VerticalDirection) (*VerticalDirection, error) {

	var output VerticalDirection

	switch verticalDirection {
	case palette.VerticalDirectionUp:
		output = VerticalDirection_VERTICAL_DIRECTION_UP
	case palette.VerticalDirectionDown:
		output = VerticalDirection_VERTICAL_DIRECTION_DOWN
	default:
		return nil, errorx.IllegalState.New("unknown vertical direction: %s", verticalDirection)
	}

	return &output, nil
}

func FromWest(west palette.West) (*West, error) {

	var output West

	switch west {
	case palette.WestFalse:
		output = West_WEST_FALSE
	case palette.WestTrue:
		output = West_WEST_TRUE
	case palette.WestNone:
		output = West_WEST_NONE
	case palette.WestLow:
		output = West_WEST_LOW
	case palette.WestTall:
		output = West_WEST_TALL
	case palette.WestUp:
		output = West_WEST_UP
	case palette.WestSide:
		output = West_WEST_SIDE
	default:
		return nil, errorx.IllegalState.New("unknown west: %s", west)
	}

	return &output, nil
}

func (a Attachment) asJSONStruct() (*palette.Attachment, error) {

	var attachment palette.Attachment

	switch a {
	case Attachment_ATTACHMENT_FLOOR:
		attachment = palette.AttachmentFloor
	case Attachment_ATTACHMENT_CEILING:
		attachment = palette.AttachmentCeiling
	case Attachment_ATTACHMENT_SINGLE_WALL:
		attachment = palette.AttachmentSingleWall
	case Attachment_ATTACHMENT_DOUBLE_WALL:
		attachment = palette.AttachmentDoubleWall
	default:
		return nil, fmt.Errorf("unknown attachment: %v", a)
	}

	return &attachment, nil
}

func (a Axis) asJSONStruct() (*palette.Axis, error) {

	var axis palette.Axis

	switch a {
	case Axis_AXIS_X:
		axis = palette.AxisX
	case Axis_AXIS_Y:
		axis = palette.AxisY
	case Axis_AXIS_Z:
		axis = palette.AxisZ
	default:
		return nil, fmt.Errorf("unknown axis: %v", a)
	}

	return &axis, nil
}

func (e East) asJSONStruct() (*palette.East, error) {

	var east palette.East

	switch e {
	case East_EAST_FALSE:
		east = palette.EastFalse
	case East_EAST_TRUE:
		east = palette.EastTrue
	case East_EAST_NONE:
		east = palette.EastNone
	case East_EAST_LOW:
		east = palette.EastLow
	case East_EAST_TALL:
		east = palette.EastTall
	case East_EAST_UP:
		east = palette.EastUp
	case East_EAST_SIDE:
		east = palette.EastSide
	default:
		return nil, fmt.Errorf("unknown east: %v", e)
	}

	return &east, nil
}

func (f Face) asJSONStruct() (*palette.Face, error) {

	var face palette.Face

	switch f {
	case Face_FACE_FLOOR:
		face = palette.FaceFloor
	case Face_FACE_CEILING:
		face = palette.FaceCeiling
	case Face_FACE_WALL:
		face = palette.FaceWall
	default:
		return nil, fmt.Errorf("unknown face: %v", f)
	}

	return &face, nil
}

func (f Facing) asJSONStruct() (*palette.Facing, error) {

	var facing palette.Facing

	switch f {
	case Facing_FACING_NORTH:
		facing = palette.FacingNorth
	case Facing_FACING_SOUTH:
		facing = palette.FacingSouth
	case Facing_FACING_WEST:
		facing = palette.FacingWest
	case Facing_FACING_EAST:
		facing = palette.FacingEast
	case Facing_FACING_DOWN:
		facing = palette.FacingDown
	case Facing_FACING_UP:
		facing = palette.FacingUp
	default:
		return nil, fmt.Errorf("unknown facing: %v", f)
	}

	return &facing, nil
}

func (h Half) asJSONStruct() (*palette.Half, error) {

	var half palette.Half

	switch h {
	case Half_HALF_TOP:
		half = palette.HalfTop
	case Half_HALF_BOTTOM:
		half = palette.HalfBottom
	case Half_HALF_UPPER:
		half = palette.HalfUpper
	case Half_HALF_LOWER:
		half = palette.HalfLower
	default:
		return nil, fmt.Errorf("unknown half: %v", h)
	}

	return &half, nil
}

func (h Hinge) asJSONStruct() (*palette.Hinge, error) {

	var hinge palette.Hinge

	switch h {
	case Hinge_HINGE_LEFT:
		hinge = palette.HingeLeft
	case Hinge_HINGE_RIGHT:
		hinge = palette.HingeRight
	default:
		return nil, fmt.Errorf("unknown hinge: %v", h)
	}

	return &hinge, nil
}

func (i Instrument) asJSONStruct() (*palette.Instrument, error) {

	var instrument palette.Instrument

	switch i {
	case Instrument_INSTRUMENT_HARP:
		instrument = palette.InstrumentHarp
	case Instrument_INSTRUMENT_BASEDRUM:
		instrument = palette.InstrumentBasedrum
	case Instrument_INSTRUMENT_SNARE:
		instrument = palette.InstrumentSnare
	case Instrument_INSTRUMENT_HAT:
		instrument = palette.InstrumentHat
	case Instrument_INSTRUMENT_BASS:
		instrument = palette.InstrumentBass
	case Instrument_INSTRUMENT_FLUTE:
		instrument = palette.InstrumentFlute
	case Instrument_INSTRUMENT_BELL:
		instrument = palette.InstrumentBell
	case Instrument_INSTRUMENT_GUITAR:
		instrument = palette.InstrumentGuitar
	case Instrument_INSTRUMENT_CHIME:
		instrument = palette.InstrumentChime
	case Instrument_INSTRUMENT_XYLOPHONE:
		instrument = palette.InstrumentXylophone
	case Instrument_INSTRUMENT_IRON_XYLOPHONE:
		instrument = palette.InstrumentIronXylophone
	case Instrument_INSTRUMENT_COW_BELL:
		instrument = palette.InstrumentCowBell
	case Instrument_INSTRUMENT_DIDGERIDOO:
		instrument = palette.InstrumentDidgeridoo
	case Instrument_INSTRUMENT_BIT:
		instrument = palette.InstrumentBit
	case Instrument_INSTRUMENT_BANJO:
		instrument = palette.InstrumentBanjo
	case Instrument_INSTRUMENT_PLING:
		instrument = palette.InstrumentPling
	case Instrument_INSTRUMENT_ZOMBIE:
		instrument = palette.InstrumentZombie
	case Instrument_INSTRUMENT_SKELETON:
		instrument = palette.InstrumentSkeleton
	case Instrument_INSTRUMENT_CREEPER:
		instrument = palette.InstrumentCreeper
	case Instrument_INSTRUMENT_DRAGON:
		instrument = palette.InstrumentDragon
	case Instrument_INSTRUMENT_WITHER_SKELETON:
		instrument = palette.InstrumentWitherSkeleton
	case Instrument_INSTRUMENT_PIGLIN:
		instrument = palette.InstrumentPiglin
	case Instrument_INSTRUMENT_CUSTOM_HEAD:
		instrument = palette.InstrumentCustomHead
	default:
		return nil, fmt.Errorf("unknown instrument: %v", i)
	}

	return &instrument, nil
}

func (l Leaves) asJSONStruct() (*palette.Leaves, error) {

	var leaves palette.Leaves

	switch l {
	case Leaves_LEAVES_NONE:
		leaves = palette.LeavesNone
	case Leaves_LEAVES_SMALL:
		leaves = palette.LeavesSmall
	case Leaves_LEAVES_LARGE:
		leaves = palette.LeavesLarge
	default:
		return nil, fmt.Errorf("unknown leaves: %v", l)
	}

	return &leaves, nil
}

func (m Mode) asJSONStruct() (*palette.Mode, error) {

	var mode palette.Mode

	switch m {
	case Mode_MODE_SAVE:
		mode = palette.ModeSave
	case Mode_MODE_LOAD:
		mode = palette.ModeLoad
	case Mode_MODE_CORNER:
		mode = palette.ModeCorner
	case Mode_MODE_DATA:
		mode = palette.ModeData
	case Mode_MODE_COMPARE:
		mode = palette.ModeCompare
	case Mode_MODE_SUBTRACT:
		mode = palette.ModeSubtract
	default:
		return nil, fmt.Errorf("unknown mode: %v", m)
	}

	return &mode, nil
}

func (n North) asJSONStruct() (*palette.North, error) {

	var north palette.North

	switch n {
	case North_NORTH_FALSE:
		north = palette.NorthFalse
	case North_NORTH_TRUE:
		north = palette.NorthTrue
	case North_NORTH_NONE:
		north = palette.NorthNone
	case North_NORTH_LOW:
		north = palette.NorthLow
	case North_NORTH_TALL:
		north = palette.NorthTall
	case North_NORTH_UP:
		north = palette.NorthUp
	case North_NORTH_SIDE:
		north = palette.NorthSide
	default:
		return nil, fmt.Errorf("unknown north: %v", n)
	}

	return &north, nil
}

func (o Orientation) asJSONStruct() (*palette.Orientation, error) {

	var orientation palette.Orientation

	switch o {
	case Orientation_ORIENTATION_DOWN_EAST:
		orientation = palette.OrientationDownEast
	case Orientation_ORIENTATION_DOWN_NORTH:
		orientation = palette.OrientationDownNorth
	case Orientation_ORIENTATION_DOWN_SOUTH:
		orientation = palette.OrientationDownSouth
	case Orientation_ORIENTATION_DOWN_WEST:
		orientation = palette.OrientationDownWest
	case Orientation_ORIENTATION_EAST_UP:
		orientation = palette.OrientationEastUp
	case Orientation_ORIENTATION_NORTH_UP:
		orientation = palette.OrientationNorthUp
	case Orientation_ORIENTATION_SOUTH_UP:
		orientation = palette.OrientationSouthUp
	case Orientation_ORIENTATION_UP_EAST:
		orientation = palette.OrientationUpEast
	case Orientation_ORIENTATION_UP_NORTH:
		orientation = palette.OrientationUpNorth
	case Orientation_ORIENTATION_UP_SOUTH:
		orientation = palette.OrientationUpSouth
	case Orientation_ORIENTATION_UP_WEST:
		orientation = palette.OrientationUpWest
	case Orientation_ORIENTATION_WEST_UP:
		orientation = palette.OrientationWestUp
	default:
		return nil, fmt.Errorf("unknown orientation: %v", o)
	}

	return &orientation, nil
}

func (s SculkSensorPhase) asJSONStruct() (*palette.SculkSensorPhase, error) {

	var sculkSensorPhase palette.SculkSensorPhase

	switch s {
	case SculkSensorPhase_SCULK_SENSOR_PHASE_INACTIVE:
		sculkSensorPhase = palette.SculkSensorPhaseInactive
	case SculkSensorPhase_SCULK_SENSOR_PHASE_ACTIVE:
		sculkSensorPhase = palette.SculkSensorPhaseActive
	case SculkSensorPhase_SCULK_SENSOR_PHASE_COOLDOWN:
		sculkSensorPhase = palette.SculkSensorPhaseCooldown
	default:
		return nil, fmt.Errorf("unknown sculkSensorPhase: %v", s)
	}

	return &sculkSensorPhase, nil
}

func (p Part) asJSONStruct() (*palette.Part, error) {

	var part palette.Part

	switch p {
	case Part_PART_HEAD:
		part = palette.PartHead
	case Part_PART_FOOT:
		part = palette.PartFoot
	default:
		return nil, fmt.Errorf("unknown part: %v", p)
	}

	return &part, nil
}

func (s Shape) asJSONStruct() (*palette.Shape, error) {

	var shape palette.Shape

	switch s {
	case Shape_SHAPE_STRAIGHT:
		shape = palette.ShapeStraight
	case Shape_SHAPE_INNER_LEFT:
		shape = palette.ShapeInnerLeft
	case Shape_SHAPE_INNER_RIGHT:
		shape = palette.ShapeInnerRight
	case Shape_SHAPE_OUTER_LEFT:
		shape = palette.ShapeOuterLeft
	case Shape_SHAPE_OUTER_RIGHT:
		shape = palette.ShapeOuterRight
	case Shape_SHAPE_NORTH_SOUTH:
		shape = palette.ShapeNorthSouth
	case Shape_SHAPE_EAST_WEST:
		shape = palette.ShapeEastWest
	case Shape_SHAPE_ASCENDING_EAST:
		shape = palette.ShapeAscendingEast
	case Shape_SHAPE_ASCENDING_WEST:
		shape = palette.ShapeAscendingWest
	case Shape_SHAPE_ASCENDING_NORTH:
		shape = palette.ShapeAscendingNorth
	case Shape_SHAPE_ASCENDING_SOUTH:
		shape = palette.ShapeAscendingSouth
	case Shape_SHAPE_SOUTH_EAST:
		shape = palette.ShapeSouthEast
	case Shape_SHAPE_SOUTH_WEST:
		shape = palette.ShapeSouthWest
	case Shape_SHAPE_NORTH_WEST:
		shape = palette.ShapeNorthWest
	case Shape_SHAPE_NORTH_EAST:
		shape = palette.ShapeNorthEast
	default:
		return nil, fmt.Errorf("unknown shape: %v", s)
	}

	return &shape, nil
}

func (s South) asJSONStruct() (*palette.South, error) {

	var south palette.South

	switch s {
	case South_SOUTH_FALSE:
		south = palette.SouthFalse
	case South_SOUTH_TRUE:
		south = palette.SouthTrue
	case South_SOUTH_NONE:
		south = palette.SouthNone
	case South_SOUTH_LOW:
		south = palette.SouthLow
	case South_SOUTH_TALL:
		south = palette.SouthTall
	case South_SOUTH_UP:
		south = palette.SouthUp
	case South_SOUTH_SIDE:
		south = palette.SouthSide
	default:
		return nil, fmt.Errorf("unknown south: %v", s)
	}

	return &south, nil
}

func (t Thickness) asJSONStruct() (*palette.Thickness, error) {

	var thickness palette.Thickness

	switch t {
	case Thickness_THICKNESS_TIP_MERGE:
		thickness = palette.ThicknessTipMerge
	case Thickness_THICKNESS_TIP:
		thickness = palette.ThicknessTip
	case Thickness_THICKNESS_FRUSTUM:
		thickness = palette.ThicknessFrustum
	case Thickness_THICKNESS_MIDDLE:
		thickness = palette.ThicknessMiddle
	case Thickness_THICKNESS_BASE:
		thickness = palette.ThicknessBase
	default:
		return nil, fmt.Errorf("unknown thickness: %v", t)
	}

	return &thickness, nil
}

func (t Tilt) asJSONStruct() (*palette.Tilt, error) {

	var tilt palette.Tilt

	switch t {
	case Tilt_TILT_NONE:
		tilt = palette.TiltNone
	case Tilt_TILT_UNSTABLE:
		tilt = palette.TiltUnstable
	case Tilt_TILT_PARTIAL:
		tilt = palette.TiltPartial
	case Tilt_TILT_FULL:
		tilt = palette.TiltFull
	default:
		return nil, fmt.Errorf("unknown tilt: %v", t)
	}

	return &tilt, nil
}

func (t TrialSpawnerState) asJSONStruct() (*palette.TrialSpawnerState, error) {

	var trialSpawnerState palette.TrialSpawnerState

	switch t {
	case TrialSpawnerState_TRIAL_SPAWNER_STATE_INACTIVE:
		trialSpawnerState = palette.TrialSpawnerStateInActive
	case TrialSpawnerState_TRIAL_SPAWNER_STATE_WAITING_FOR_PLAYERS:
		trialSpawnerState = palette.TrialSpawnerStateWaitingForPlayers
	case TrialSpawnerState_TRIAL_SPAWNER_STATE_ACTIVE:
		trialSpawnerState = palette.TrialSpawnerStateActive
	case TrialSpawnerState_TRIAL_SPAWNER_STATE_WAITING_FOR_REWARD_EJECTION:
		trialSpawnerState = palette.TrialSpawnerStateWaitingForRewardEjection
	case TrialSpawnerState_TRIAL_SPAWNER_STATE_EJECTING_REWARD:
		trialSpawnerState = palette.TrialSpawnerStateEjectingReward
	case TrialSpawnerState_TRIAL_SPAWNER_STATE_COOLDOWN:
		trialSpawnerState = palette.TrialSpawnerStateCooldown
	default:
		return nil, fmt.Errorf("unknown trialSpawnerState: %v", t)
	}

	return &trialSpawnerState, nil
}

func (t Type) asJSONStruct() (*palette.Type, error) {

	var type_ palette.Type

	switch t {
	case Type_TYPE_SINGLE:
		type_ = palette.TypeSingle
	case Type_TYPE_LEFT:
		type_ = palette.TypeLeft
	case Type_TYPE_RIGHT:
		type_ = palette.TypeRight
	case Type_TYPE_TOP:
		type_ = palette.TypeTop
	case Type_TYPE_BOTTOM:
		type_ = palette.TypeBottom
	case Type_TYPE_DOUBLE:
		type_ = palette.TypeDouble
	case Type_TYPE_STICKY:
		type_ = palette.TypeSticky
	case Type_TYPE_NORMAL:
		type_ = palette.TypeNormal
	default:
		return nil, fmt.Errorf("unknown type: %v", t)
	}

	return &type_, nil
}

func (v VerticalDirection) asJSONStruct() (*palette.VerticalDirection, error) {

	var verticalDirection palette.VerticalDirection

	switch v {
	case VerticalDirection_VERTICAL_DIRECTION_UP:
		verticalDirection = palette.VerticalDirectionUp
	case VerticalDirection_VERTICAL_DIRECTION_DOWN:
		verticalDirection = palette.VerticalDirectionDown
	default:
		return nil, fmt.Errorf("unknown verticalDirection: %v", v)
	}

	return &verticalDirection, nil
}

func (w West) asJSONStruct() (*palette.West, error) {

	var west palette.West

	switch w {
	case West_WEST_FALSE:
		west = palette.WestFalse
	case West_WEST_TRUE:
		west = palette.WestTrue
	case West_WEST_NONE:
		west = palette.WestNone
	case West_WEST_LOW:
		west = palette.WestLow
	case West_WEST_TALL:
		west = palette.WestTall
	case West_WEST_UP:
		west = palette.WestUp
	case West_WEST_SIDE:
		west = palette.WestSide

	default:
		return nil, fmt.Errorf("unknown west: %v", w)
	}

	return &west, nil
}

// TODO: As json representation (Will help with testing)
