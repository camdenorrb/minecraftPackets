package palette

type Attachment string

const (
	AttachmentFloor      Attachment = "floor"
	AttachmentCeiling    Attachment = "ceiling"
	AttachmentSingleWall Attachment = "single_wall"
	AttachmentDoubleWall Attachment = "double_wall"
)

type Axis string

const (
	AxisX Axis = "x"
	AxisY Axis = "y"
	AxisZ Axis = "z"
)

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

type Face string

const (
	FaceFloor   Face = "floor"
	FaceWall    Face = "wall"
	FaceCeiling Face = "ceiling"
)

type Facing string

const (
	FacingNorth Facing = "north"
	FacingSouth Facing = "south"
	FacingWest  Facing = "west"
	FacingEast  Facing = "east"
	FacingUp    Facing = "up"
	FacingDown  Facing = "down"
)

type Half string

const (
	HalfTop    Half = "top"
	HalfBottom Half = "bottom"
	HalfUpper  Half = "upper"
	HalfLower  Half = "lower"
)

type Hinge string

const (
	HingeLeft  Hinge = "left"
	HingeRight Hinge = "right"
)

type Instrument string

const (
	InstrumentHarp           Instrument = "harp"
	InstrumentBasedrum       Instrument = "basedrum"
	InstrumentSnare          Instrument = "snare"
	InstrumentHat            Instrument = "hat"
	InstrumentBass           Instrument = "bass"
	InstrumentFlute          Instrument = "flute"
	InstrumentBell           Instrument = "bell"
	InstrumentGuitar         Instrument = "guitar"
	InstrumentChime          Instrument = "chime"
	InstrumentXylophone      Instrument = "xylophone"
	InstrumentIronXylophone  Instrument = "iron_xylophone"
	InstrumentCowBell        Instrument = "cow_bell"
	InstrumentDidgeridoo     Instrument = "didgeridoo"
	InstrumentBit            Instrument = "bit"
	InstrumentBanjo          Instrument = "banjo"
	InstrumentPling          Instrument = "pling"
	InstrumentZombie         Instrument = "zombie"
	InstrumentSkeleton       Instrument = "skeleton"
	InstrumentCreeper        Instrument = "creeper"
	InstrumentDragon         Instrument = "dragon"
	InstrumentWitherSkeleton Instrument = "wither_skeleton"
	InstrumentPiglin         Instrument = "piglin"
	InstrumentCustomHead     Instrument = "custom_head"
)

type Leaves string

const (
	LeavesNone  Leaves = "none"
	LeavesSmall Leaves = "small"
	LeavesLarge Leaves = "large"
)

type Mode string

const (
	ModeSave     Mode = "save"
	ModeLoad     Mode = "load"
	ModeCorner   Mode = "corner"
	ModeData     Mode = "data"
	ModeCompare  Mode = "compare"
	ModeSubtract Mode = "subtract"
)

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

type Part string

const (
	PartHead Part = "head"
	PartFoot Part = "foot"
)

type SculkSensorPhase string

const (
	SculkSensorPhaseInactive SculkSensorPhase = "inactive"
	SculkSensorPhaseActive   SculkSensorPhase = "active"
	SculkSensorPhaseCooldown SculkSensorPhase = "cooldown"
)

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

type Thickness string

const (
	ThicknessTipMerge Thickness = "tip_merge"
	ThicknessTip      Thickness = "tip"
	ThicknessFrustum  Thickness = "frustum"
	ThicknessMiddle   Thickness = "middle"
	ThicknessBase     Thickness = "base"
)

type Tilt string

const (
	TiltNone     Tilt = "none"
	TiltUnstable Tilt = "unstable"
	TiltPartial  Tilt = "partial"
	TiltFull     Tilt = "full"
)

type TrialSpawnerState string

const (
	TrialSpawnerStateInActive                 TrialSpawnerState = "inactive"
	TrialSpawnerStateWaitingForPlayers        TrialSpawnerState = "waiting_for_players"
	TrialSpawnerStateActive                   TrialSpawnerState = "active"
	TrialSpawnerStateWaitingForRewardEjection TrialSpawnerState = "waiting_for_reward_ejection"
	TrialSpawnerStateEjectingReward           TrialSpawnerState = "ejecting_reward"
	TrialSpawnerStateCooldown                 TrialSpawnerState = "cooldown"
)

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

type VerticalDirection string

const (
	VerticalDirectionUp   VerticalDirection = "up"
	VerticalDirectionDown VerticalDirection = "down"
)

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
