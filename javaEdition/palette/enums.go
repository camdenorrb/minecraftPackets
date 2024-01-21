package palette

import (
	"github.com/MisterKaiou/go-functional/result"
	"github.com/camdenorrb/minecraftPackets/javaEdition/palette/proto"
	"github.com/joomcode/errorx"
)

type Attachment string

const (
	AttachmentFloor      Attachment = "floor"
	AttachmentCeiling    Attachment = "ceiling"
	AttachmentSingleWall Attachment = "single_wall"
	AttachmentDoubleWall Attachment = "double_wall"
)

func (a Attachment) AsProto() result.Of[proto.Attachment] {
	switch a {
	case AttachmentFloor:
		return result.Ok(proto.Attachment_ATTACHMENT_FLOOR)
	case AttachmentCeiling:
		return result.Ok(proto.Attachment_ATTACHMENT_CEILING)
	case AttachmentSingleWall:
		return result.Ok(proto.Attachment_ATTACHMENT_SINGLE_WALL)
	case AttachmentDoubleWall:
		return result.Ok(proto.Attachment_ATTACHMENT_DOUBLE_WALL)
	default:
		return result.Error[proto.Attachment](errorx.IllegalState.New("unknown attachment: %s", a))
	}
}

type Axis string

const (
	AxisX Axis = "x"
	AxisY Axis = "y"
	AxisZ Axis = "z"
)

func (a Axis) AsProto() result.Of[proto.Axis] {
	switch a {
	case AxisX:
		return result.Ok(proto.Axis_AXIS_X)
	case AxisY:
		return result.Ok(proto.Axis_AXIS_Y)
	case AxisZ:
		return result.Ok(proto.Axis_AXIS_Z)
	default:
		return result.Error[proto.Axis](errorx.IllegalState.New("unknown axis: %s", a))
	}
}

type East string

const (
	EastNone  East = "none"
	EastLow   East = "low"
	EastTall  East = "tall"
	EastTrue  East = "true"
	EastFalse East = "false"
	EastUp    East = "up"
	EastSide  East = "side"
)

func (e East) AsProto() result.Of[proto.East] {
	switch e {
	case EastNone:
		return result.Ok(proto.East_EAST_NONE)
	case EastLow:
		return result.Ok(proto.East_EAST_LOW)
	case EastTall:
		return result.Ok(proto.East_EAST_TALL)
	case EastTrue:
		return result.Ok(proto.East_EAST_TRUE)
	case EastFalse:
		return result.Ok(proto.East_EAST_FALSE)
	case EastUp:
		return result.Ok(proto.East_EAST_UP)
	case EastSide:
		return result.Ok(proto.East_EAST_SIDE)
	default:
		return result.Error[proto.East](errorx.IllegalState.New("unknown east: %s", e))
	}
}

type Face string

const (
	FaceFloor   Face = "floor"
	FaceWall    Face = "wall"
	FaceCeiling Face = "ceiling"
)

func (f Face) AsProto() result.Of[proto.Face] {
	switch f {
	case FaceFloor:
		return result.Ok(proto.Face_FACE_FLOOR)
	case FaceWall:
		return result.Ok(proto.Face_FACE_WALL)
	case FaceCeiling:
		return result.Ok(proto.Face_FACE_CEILING)
	default:
		return result.Error[proto.Face](errorx.IllegalState.New("unknown face: %s", f))
	}
}

type Facing string

const (
	FacingNorth Facing = "north"
	FacingSouth Facing = "south"
	FacingWest  Facing = "west"
	FacingEast  Facing = "east"
	FacingUp    Facing = "up"
	FacingDown  Facing = "down"
)

func (f Facing) AsProto() result.Of[proto.Facing] {
	switch f {
	case FacingNorth:
		return result.Ok(proto.Facing_FACING_NORTH)
	case FacingSouth:
		return result.Ok(proto.Facing_FACING_SOUTH)
	case FacingWest:
		return result.Ok(proto.Facing_FACING_WEST)
	case FacingEast:
		return result.Ok(proto.Facing_FACING_EAST)
	case FacingUp:
		return result.Ok(proto.Facing_FACING_UP)
	case FacingDown:
		return result.Ok(proto.Facing_FACING_DOWN)
	default:
		return result.Error[proto.Facing](errorx.IllegalState.New("unknown facing: %s", f))
	}
}

type Half string

const (
	HalfTop    Half = "top"
	HalfBottom Half = "bottom"
	HalfUpper  Half = "upper"
	HalfLower  Half = "lower"
)

func (h Half) AsProto() result.Of[proto.Half] {
	switch h {
	case HalfTop:
		return result.Ok(proto.Half_HALF_TOP)
	case HalfBottom:
		return result.Ok(proto.Half_HALF_BOTTOM)
	case HalfUpper:
		return result.Ok(proto.Half_HALF_UPPER)
	case HalfLower:
		return result.Ok(proto.Half_HALF_LOWER)
	default:
		return result.Error[proto.Half](errorx.IllegalState.New("unknown half: %s", h))
	}
}

type Hinge string

const (
	HingeLeft  Hinge = "left"
	HingeRight Hinge = "right"
)

func (h Hinge) AsProto() result.Of[proto.Hinge] {
	switch h {
	case HingeLeft:
		return result.Ok(proto.Hinge_HINGE_LEFT)
	case HingeRight:
		return result.Ok(proto.Hinge_HINGE_RIGHT)
	default:
		return result.Error[proto.Hinge](errorx.IllegalState.New("unknown hinge: %s", h))
	}
}

type Instrument string

const (
	InstrumentHarp           Instrument = "harp"
	InstrumentBasedrum       Instrument = "basedrum"
	InstrumentSnare          Instrument = "snare"
	InstrumentHat            Instrument = "hat"
	InstrumentBass           Instrument = "bass"
	IntrumentFlute           Instrument = "flute"
	IntrumentBell            Instrument = "bell"
	IntrumentGuitar          Instrument = "guitar"
	IntrumentChime           Instrument = "chime"
	IntrumentXylophone       Instrument = "xylophone"
	InstrumentIronXylophone  Instrument = "iron_xylophone"
	IntrumentCowBell         Instrument = "cow_bell"
	IntrumentDidgeridoo      Instrument = "didgeridoo"
	IntrumentBit             Instrument = "bit"
	IntrumentBanjo           Instrument = "banjo"
	instrumentPling          Instrument = "pling"
	InstrumentZombie         Instrument = "zombie"
	InstrumentSkeleton       Instrument = "skeleton"
	InstrumentCreeper        Instrument = "creeper"
	InstrumentDragon         Instrument = "dragon"
	InstrumentWitherSkeleton Instrument = "wither_skeleton"
	InstrumentPiglin         Instrument = "piglin"
	InstrumentCustomHead     Instrument = "custom_head"
)

func (i Instrument) AsProto() result.Of[proto.Instrument] {
	switch i {
	case InstrumentHarp:
		return result.Ok(proto.Instrument_INSTRUMENT_HARP)
	case InstrumentBasedrum:
		return result.Ok(proto.Instrument_INSTRUMENT_BASEDRUM)
	case InstrumentSnare:
		return result.Ok(proto.Instrument_INSTRUMENT_SNARE)
	case InstrumentHat:
		return result.Ok(proto.Instrument_INSTRUMENT_HAT)
	case InstrumentBass:
		return result.Ok(proto.Instrument_INSTRUMENT_BASS)
	case IntrumentFlute:
		return result.Ok(proto.Instrument_INSTRUMENT_FLUTE)
	case IntrumentBell:
		return result.Ok(proto.Instrument_INSTRUMENT_BELL)
	case IntrumentGuitar:
		return result.Ok(proto.Instrument_INSTRUMENT_GUITAR)
	case IntrumentChime:
		return result.Ok(proto.Instrument_INSTRUMENT_CHIME)
	case IntrumentXylophone:
		return result.Ok(proto.Instrument_INSTRUMENT_XYLOPHONE)
	case InstrumentIronXylophone:
		return result.Ok(proto.Instrument_INSTRUMENT_IRON_XYLOPHONE)
	case IntrumentCowBell:
		return result.Ok(proto.Instrument_INSTRUMENT_COW_BELL)
	case IntrumentDidgeridoo:
		return result.Ok(proto.Instrument_INSTRUMENT_DIDGERIDOO)
	case IntrumentBit:
		return result.Ok(proto.Instrument_INSTRUMENT_BIT)
	case IntrumentBanjo:
		return result.Ok(proto.Instrument_INSTRUMENT_BANJO)
	case instrumentPling:
		return result.Ok(proto.Instrument_INSTRUMENT_PLING)
	case InstrumentZombie:
		return result.Ok(proto.Instrument_INSTRUMENT_ZOMBIE)
	case InstrumentSkeleton:
		return result.Ok(proto.Instrument_INSTRUMENT_SKELETON)
	case InstrumentCreeper:
		return result.Ok(proto.Instrument_INSTRUMENT_CREEPER)
	case InstrumentDragon:
		return result.Ok(proto.Instrument_INSTRUMENT_DRAGON)
	case InstrumentWitherSkeleton:
		return result.Ok(proto.Instrument_INSTRUMENT_WITHER_SKELETON)
	case InstrumentPiglin:
		return result.Ok(proto.Instrument_INSTRUMENT_PIGLIN)
	case InstrumentCustomHead:
		return result.Ok(proto.Instrument_INSTRUMENT_CUSTOM_HEAD)
	default:
		return result.Error[proto.Instrument](errorx.IllegalState.New("unknown instrument: %s", i))
	}
}

type Leaves string

const (
	LeavesNone  Leaves = "none"
	LeavesSmall Leaves = "small"
	LeavesLarge Leaves = "large"
)

func (l Leaves) AsProto() result.Of[proto.Leaves] {
	switch l {
	case LeavesNone:
		return result.Ok(proto.Leaves_LEAVES_NONE)
	case LeavesSmall:
		return result.Ok(proto.Leaves_LEAVES_SMALL)
	case LeavesLarge:
		return result.Ok(proto.Leaves_LEAVES_LARGE)
	default:
		return result.Error[proto.Leaves](errorx.IllegalState.New("unknown leaves: %s", l))
	}
}

type Mode string

const (
	ModeSave     Mode = "save"
	ModeLoad     Mode = "load"
	ModeCorner   Mode = "corner"
	ModeData     Mode = "data"
	ModeCompare  Mode = "compare"
	ModeSubtract Mode = "subtract"
)

func (m Mode) AsProto() result.Of[proto.Mode] {
	switch m {
	case ModeSave:
		return result.Ok(proto.Mode_MODE_SAVE)
	case ModeLoad:
		return result.Ok(proto.Mode_MODE_LOAD)
	case ModeCorner:
		return result.Ok(proto.Mode_MODE_CORNER)
	case ModeData:
		return result.Ok(proto.Mode_MODE_DATA)
	case ModeCompare:
		return result.Ok(proto.Mode_MODE_COMPARE)
	case ModeSubtract:
		return result.Ok(proto.Mode_MODE_SUBTRACT)

	default:
		return result.Error[proto.Mode](errorx.IllegalState.New("unknown mode: %s", m))
	}
}

type North string

const (
	NorthNone  North = "none"
	NorthLow   North = "low"
	NorthTall  North = "tall"
	NorthTrue  North = "true"
	NorthFalse North = "false"
	NorthUp    North = "up"
	NorthSide  North = "side"
)

func (n North) AsProto() result.Of[proto.North] {
	switch n {
	case NorthNone:
		return result.Ok(proto.North_NORTH_NONE)
	case NorthLow:
		return result.Ok(proto.North_NORTH_LOW)
	case NorthTall:
		return result.Ok(proto.North_NORTH_TALL)
	case NorthTrue:
		return result.Ok(proto.North_NORTH_TRUE)
	case NorthFalse:
		return result.Ok(proto.North_NORTH_FALSE)
	case NorthUp:
		return result.Ok(proto.North_NORTH_UP)
	case NorthSide:
		return result.Ok(proto.North_NORTH_SIDE)
	default:
		return result.Error[proto.North](errorx.IllegalState.New("unknown north: %s", n))
	}
}

type Orientation string

const (
	OrientationDownEast  Orientation = "down_east"
	OrientationDownNorth Orientation = "down_north"
	OrientationDownSouth Orientation = "down_south"
	OrientationDownWest  Orientation = "down_west"
	OrientationUpEast    Orientation = "up_east"
	OrientationUpNorth   Orientation = "up_north"
	OrientationUpSouth   Orientation = "up_south"
	OrientationUpWest    Orientation = "up_west"
	OrientationWestUp    Orientation = "west_up"
	OrientationEastUp    Orientation = "east_up"
	OrientationNorthUp   Orientation = "north_up"
	OrientationSouthUp   Orientation = "south_up"
)

func (o Orientation) AsProto() result.Of[proto.Orientation] {
	switch o {
	case OrientationDownEast:
		return result.Ok(proto.Orientation_ORIENTATION_DOWN_EAST)
	case OrientationDownNorth:
		return result.Ok(proto.Orientation_ORIENTATION_DOWN_NORTH)
	case OrientationDownSouth:
		return result.Ok(proto.Orientation_ORIENTATION_DOWN_SOUTH)
	case OrientationDownWest:
		return result.Ok(proto.Orientation_ORIENTATION_DOWN_WEST)
	case OrientationUpEast:
		return result.Ok(proto.Orientation_ORIENTATION_UP_EAST)
	case OrientationUpNorth:
		return result.Ok(proto.Orientation_ORIENTATION_UP_NORTH)
	case OrientationUpSouth:
		return result.Ok(proto.Orientation_ORIENTATION_UP_SOUTH)
	case OrientationUpWest:
		return result.Ok(proto.Orientation_ORIENTATION_UP_WEST)
	case OrientationWestUp:
		return result.Ok(proto.Orientation_ORIENTATION_WEST_UP)
	case OrientationEastUp:
		return result.Ok(proto.Orientation_ORIENTATION_EAST_UP)
	case OrientationNorthUp:
		return result.Ok(proto.Orientation_ORIENTATION_NORTH_UP)
	case OrientationSouthUp:
		return result.Ok(proto.Orientation_ORIENTATION_SOUTH_UP)
	default:
		return result.Error[proto.Orientation](errorx.IllegalState.New("unknown orientation: %s", o))
	}
}

type Part string

const (
	PartHead Part = "head"
	PartFoot Part = "foot"
)

func (p Part) AsProto() result.Of[proto.Part] {
	switch p {
	case PartHead:
		return result.Ok(proto.Part_PART_HEAD)
	case PartFoot:
		return result.Ok(proto.Part_PART_FOOT)
	default:
		return result.Error[proto.Part](errorx.IllegalState.New("unknown part: %s", p))
	}
}

type SculkSensorPhase string

const (
	SculkSensorPhaseInactive SculkSensorPhase = "inactive"
	SculkSensorPhaseActive   SculkSensorPhase = "active"
	SculkSensorPhaseCooldown SculkSensorPhase = "cooldown"
)

func (s SculkSensorPhase) AsProto() result.Of[proto.SculkSensorPhase] {
	switch s {
	case SculkSensorPhaseInactive:
		return result.Ok(proto.SculkSensorPhase_SCULK_SENSOR_PHASE_INACTIVE)
	case SculkSensorPhaseActive:
		return result.Ok(proto.SculkSensorPhase_SCULK_SENSOR_PHASE_ACTIVE)
	case SculkSensorPhaseCooldown:
		return result.Ok(proto.SculkSensorPhase_SCULK_SENSOR_PHASE_COOLDOWN)
	default:
		return result.Error[proto.SculkSensorPhase](errorx.IllegalState.New("unknown sculk sensor phase: %s", s))
	}
}

type Shape string

const (
	ShapeNorthSouth     Shape = "north_south"
	ShapeEastWest       Shape = "east_west"
	ShapeAscendingEast  Shape = "ascending_east"
	ShapeAscendingWest  Shape = "ascending_west"
	ShapeAscendingNorth Shape = "ascending_north"
	ShapeAscendingSouth Shape = "ascending_south"
	ShapeSouthEast      Shape = "south_east"
	ShapeSouthWest      Shape = "south_west"
	ShapeNorthWest      Shape = "north_west"
	ShapeNorthEast      Shape = "north_east"
	ShapeStraight       Shape = "straight"
	ShapeInnerLeft      Shape = "inner_left"
	ShapeInnerRight     Shape = "inner_right"
	ShapeOuterLeft      Shape = "outer_left"
	ShapeOuterRight     Shape = "outer_right"
)

func (s Shape) AsProto() result.Of[proto.Shape] {
	switch s {
	case ShapeNorthSouth:
		return result.Ok(proto.Shape_SHAPE_NORTH_SOUTH)
	case ShapeEastWest:
		return result.Ok(proto.Shape_SHAPE_EAST_WEST)
	case ShapeAscendingEast:
		return result.Ok(proto.Shape_SHAPE_ASCENDING_EAST)
	case ShapeAscendingWest:
		return result.Ok(proto.Shape_SHAPE_ASCENDING_WEST)
	case ShapeAscendingNorth:
		return result.Ok(proto.Shape_SHAPE_ASCENDING_NORTH)
	case ShapeAscendingSouth:
		return result.Ok(proto.Shape_SHAPE_ASCENDING_SOUTH)
	case ShapeSouthEast:
		return result.Ok(proto.Shape_SHAPE_SOUTH_EAST)
	case ShapeSouthWest:
		return result.Ok(proto.Shape_SHAPE_SOUTH_WEST)
	case ShapeNorthWest:
		return result.Ok(proto.Shape_SHAPE_NORTH_WEST)
	case ShapeNorthEast:
		return result.Ok(proto.Shape_SHAPE_NORTH_EAST)
	case ShapeStraight:
		return result.Ok(proto.Shape_SHAPE_STRAIGHT)
	case ShapeInnerLeft:
		return result.Ok(proto.Shape_SHAPE_INNER_LEFT)
	case ShapeInnerRight:
		return result.Ok(proto.Shape_SHAPE_INNER_RIGHT)
	case ShapeOuterLeft:
		return result.Ok(proto.Shape_SHAPE_OUTER_LEFT)
	case ShapeOuterRight:
		return result.Ok(proto.Shape_SHAPE_OUTER_RIGHT)
	default:
		return result.Error[proto.Shape](errorx.IllegalState.New("unknown shape: %s", s))
	}
}

type South string

const (
	SouthNone  South = "none"
	SouthLow   South = "low"
	SouthTall  South = "tall"
	SouthTrue  South = "true"
	SouthFalse South = "false"
	SouthUp    South = "up"
	SouthSide  South = "side"
)

func (s South) AsProto() result.Of[proto.South] {
	switch s {
	case SouthNone:
		return result.Ok(proto.South_SOUTH_NONE)
	case SouthLow:
		return result.Ok(proto.South_SOUTH_LOW)
	case SouthTall:
		return result.Ok(proto.South_SOUTH_TALL)
	case SouthTrue:
		return result.Ok(proto.South_SOUTH_TRUE)
	case SouthFalse:
		return result.Ok(proto.South_SOUTH_FALSE)
	case SouthUp:
		return result.Ok(proto.South_SOUTH_UP)
	case SouthSide:
		return result.Ok(proto.South_SOUTH_SIDE)
	default:
		return result.Error[proto.South](errorx.IllegalState.New("unknown south: %s", s))
	}
}

type Thickness string

const (
	ThicknessTipMerge Thickness = "tip_merge"
	ThicknessTip      Thickness = "tip"
	ThicknessFrustum  Thickness = "frustum"
	ThicknessMiddle   Thickness = "middle"
	ThicknessBase     Thickness = "base"
)

func (t Thickness) AsProto() result.Of[proto.Thickness] {
	switch t {
	case ThicknessTipMerge:
		return result.Ok(proto.Thickness_THICKNESS_TIP_MERGE)
	case ThicknessTip:
		return result.Ok(proto.Thickness_THICKNESS_TIP)
	case ThicknessFrustum:
		return result.Ok(proto.Thickness_THICKNESS_FRUSTUM)
	case ThicknessMiddle:
		return result.Ok(proto.Thickness_THICKNESS_MIDDLE)
	case ThicknessBase:
		return result.Ok(proto.Thickness_THICKNESS_BASE)
	default:
		return result.Error[proto.Thickness](errorx.IllegalState.New("unknown thickness: %s", t))
	}
}

type Tilt string

const (
	TiltNone     Tilt = "none"
	TiltUnstable Tilt = "unstable"
	TiltPartial  Tilt = "partial"
	TiltFull     Tilt = "full"
)

func (t Tilt) AsProto() result.Of[proto.Tilt] {
	switch t {
	case TiltNone:
		return result.Ok(proto.Tilt_TILT_NONE)
	case TiltUnstable:
		return result.Ok(proto.Tilt_TILT_UNSTABLE)
	case TiltPartial:
		return result.Ok(proto.Tilt_TILT_PARTIAL)
	case TiltFull:
		return result.Ok(proto.Tilt_TILT_FULL)
	default:
		return result.Error[proto.Tilt](errorx.IllegalState.New("unknown tilt: %s", t))
	}
}

type TrialSpawnerState string

const (
	TrialSpawnerStateInActive                 TrialSpawnerState = "inactive"
	TrialSpawnerStateWaitingForPlayers        TrialSpawnerState = "waiting_for_players"
	TrialSpawnerStateActive                   TrialSpawnerState = "active"
	TrialSpawnerStateWaitingForRewardEjection TrialSpawnerState = "waiting_for_reward_ejection"
	TrialSpawnerStateEjectingReward           TrialSpawnerState = "ejecting_reward"
	TrialSpawnerStateCooldown                 TrialSpawnerState = "cooldown"
)

func (t TrialSpawnerState) AsProto() result.Of[proto.TrialSpawnerState] {
	switch t {
	case TrialSpawnerStateInActive:
		return result.Ok(proto.TrialSpawnerState_TRIAL_SPAWNER_STATE_INACTIVE)
	case TrialSpawnerStateWaitingForPlayers:
		return result.Ok(proto.TrialSpawnerState_TRIAL_SPAWNER_STATE_WAITING_FOR_PLAYERS)
	case TrialSpawnerStateActive:
		return result.Ok(proto.TrialSpawnerState_TRIAL_SPAWNER_STATE_ACTIVE)
	case TrialSpawnerStateWaitingForRewardEjection:
		return result.Ok(proto.TrialSpawnerState_TRIAL_SPAWNER_STATE_WAITING_FOR_REWARD_EJECTION)
	case TrialSpawnerStateEjectingReward:
		return result.Ok(proto.TrialSpawnerState_TRIAL_SPAWNER_STATE_EJECTING_REWARD)
	case TrialSpawnerStateCooldown:
		return result.Ok(proto.TrialSpawnerState_TRIAL_SPAWNER_STATE_COOLDOWN)
	default:
		return result.Error[proto.TrialSpawnerState](errorx.IllegalState.New("unknown trial spawner state: %s", t))
	}
}

type Type string

const (
	TypeTop    Type = "top"
	TypeBottom Type = "bottom"
	TypeDouble Type = "double"
	TypeNormal Type = "normal"
	TypeSingle Type = "single"
	TypeSticky Type = "sticky"
	TypeLeft   Type = "left"
	TypeRight  Type = "right"
)

func (t Type) AsProto() result.Of[proto.Type] {
	switch t {
	case TypeTop:
		return result.Ok(proto.Type_TYPE_TOP)
	case TypeBottom:
		return result.Ok(proto.Type_TYPE_BOTTOM)
	case TypeDouble:
		return result.Ok(proto.Type_TYPE_DOUBLE)
	case TypeNormal:
		return result.Ok(proto.Type_TYPE_NORMAL)
	case TypeSingle:
		return result.Ok(proto.Type_TYPE_SINGLE)
	case TypeSticky:
		return result.Ok(proto.Type_TYPE_STICKY)
	case TypeLeft:
		return result.Ok(proto.Type_TYPE_LEFT)
	case TypeRight:
		return result.Ok(proto.Type_TYPE_RIGHT)
	default:
		return result.Error[proto.Type](errorx.IllegalState.New("unknown type: %s", t))
	}
}

type VerticalDirection string

const (
	VerticalDirectionUp   VerticalDirection = "up"
	VerticalDirectionDown VerticalDirection = "down"
)

func (v VerticalDirection) AsProto() result.Of[proto.VerticalDirection] {
	switch v {
	case VerticalDirectionUp:
		return result.Ok(proto.VerticalDirection_VERTICAL_DIRECTION_UP)
	case VerticalDirectionDown:
		return result.Ok(proto.VerticalDirection_VERTICAL_DIRECTION_DOWN)
	default:
		return result.Error[proto.VerticalDirection](errorx.IllegalState.New("unknown vertical direction: %s", v))
	}
}

type West string

const (
	WestNone  West = "none"
	WestLow   West = "low"
	WestTall  West = "tall"
	WestTrue  West = "true"
	WestFalse West = "false"
	WestUp    West = "up"
	WestSide  West = "side"
)

func (w West) AsProto() result.Of[proto.West] {
	switch w {
	case WestNone:
		return result.Ok(proto.West_WEST_NONE)
	case WestLow:
		return result.Ok(proto.West_WEST_LOW)
	case WestTall:
		return result.Ok(proto.West_WEST_TALL)
	case WestTrue:
		return result.Ok(proto.West_WEST_TRUE)
	case WestFalse:
		return result.Ok(proto.West_WEST_FALSE)
	case WestUp:
		return result.Ok(proto.West_WEST_UP)
	case WestSide:
		return result.Ok(proto.West_WEST_SIDE)
	default:
		return result.Error[proto.West](errorx.IllegalState.New("unknown west: %s", w))
	}
}
