package play

import (
	"github.com/google/uuid"
	"minecraftPackets/javaEdition/types"
)

type SpawnEntity struct {
	EntityID   types.VarInt
	ObjectUUID string
	Type       types.VarInt
	X          float64
	Y          float64
	Z          float64
	Pitch      types.Angle
	Yaw        types.Angle
	HeadYaw    types.Angle
	Data       types.VarInt
	VelocityX  int16
	VelocityY  int16
	VelocityZ  int16
}

type SpawnExperienceOrb struct {
	EntityID types.VarInt
	X        float64
	Y        float64
	Z        float64
	Count    int16
}

type SpawnPlayer struct {
	EntityID   types.VarInt
	PlayerUUID uuid.UUID
	X          float64
	Y          float64
	Z          float64
	Yaw        types.Angle
	Pitch      types.Angle
}

type EntityAnimation struct {
	EntityID  types.VarInt
	Animation uint8 // TODO: Enum
}

type AwardStatistic struct {
	CategoryID  types.VarInt
	StatisticID types.VarInt
	Value       types.VarInt
}

type AwardStatistics []AwardStatistic

type AcknowledgeBlockChange struct {
	SequenceID types.VarInt
}

type SetBlockDestroyStage struct {
	EntityID types.VarInt
	Location types.BlockPosition
	Stage    byte // TODO: Enum
}

type BlockEntityData struct {
	Location types.BlockPosition
	Type     types.VarInt
	NBTData  NBTTag // TODO: NBT
}

type BlockAction struct {
	Location    types.BlockPosition
	ActionID    uint8
	ActionParam uint8
	BlockType   types.VarInt
}

type BlockUpdate struct {
	Location types.BlockPosition
	BlockID  types.VarInt
}

type BossBarActionAdd struct {
	Title    types.Chat
	Health   float32
	Color    types.VarInt // TODO: Enum
	Division types.VarInt // TODO: Enum
	Flags    uint8        // TODO: Enum
}

type BossBarActionRemove struct{}

type BossBarActionUpdateHealth struct {
	Health float32
}

type BossBarActionUpdateTitle struct {
	Title types.Chat
}

type BossBarActionUpdateStyle struct {
	Color    types.VarInt // TODO: Enum
	Division types.VarInt // TODO: Enum
}

type BossBarActionUpdateFlags struct {
	Flags uint8 // TODO: Enum
}

type BossBar struct {
	UUID         uuid.UUID
	Action       types.VarInt // TODO: Enum
	Add          *BossBarActionAdd
	Remove       *BossBarActionRemove
	UpdateHealth *BossBarActionUpdateHealth
	UpdateTitle  *BossBarActionUpdateTitle
	UpdateStyle  *BossBarActionUpdateStyle
	UpdateFlags  *BossBarActionUpdateFlags
}

type ChangeDifficulty struct {
	Difficulty         uint8 // TODO: Enum
	IsDifficultyLocked bool
}

type ClearTitles struct {
	Reset bool
}

type CommandSuggestionsMatch struct {
	Match      string
	HasToolTip bool
	Tooltip    *types.Chat
}

type CommandSuggestionsResponse struct {
	ID      types.VarInt
	Start   types.VarInt
	Length  types.VarInt
	Count   types.VarInt
	Matches []CommandSuggestionsMatch
}

type Commands struct {
	Count     types.VarInt
	Nodes     []CommandNode // TODO: Command Node
	RootIndex types.VarInt
}

type CloseContainer struct {
	WindowID uint8
}

type Slot struct {
	IsPresent bool
	ItemID    types.VarInt
	ItemCount byte
	NBTData   NBT // TODO: NBT
}

type SetContainerContent struct {
	WindowID    uint8
	StateID     types.VarInt
	Count       types.VarInt
	SlotData    []Slot
	CarriedItem Slot
}

type SetContainerProperty struct {
	WindowID uint8
	Property int16
	Value    int16
}

type SetContainerSlot struct {
	WindowID uint8
	StateID  types.VarInt
	Slot     int16
	SlotData Slot
}

type SetCooldown struct {
	ItemID types.VarInt
	Ticks  types.VarInt
}

type ChatSuggestions struct {
	Action  types.VarInt // TODO: Enum
	Count   types.VarInt
	Entries []string
}

type PluginMessage struct {
	Channel types.Identifier
	Data    []byte
}

type DeleteMessage struct {
	SignatureLength types.VarInt
	Signature       []byte
}

type Disconnect struct {
	Reason types.Chat
}

type DisguisedChatMessage struct {
	Message       types.Chat
	ChatType      types.VarInt // TODO: Enum
	ChatTypeName  types.Chat
	HasTargetName bool
	TargetName    types.Chat
}

type EntityEvent struct {
	EntityID     int32
	EntityStatus uint8 // TODO: Enum
}

type Explosion struct {
	X             float64
	Y             float64
	Z             float64
	Strength      float32
	RecordCount   types.VarInt
	Records       [][]byte
	PlayerMotionX float32
	PlayerMotionY float32
	PlayerMotionZ float32
}

type ChunkUnload struct {
	ChunkX int32
	ChunkZ int32
}

type GameEvent struct {
	Event uint8 // TODO: Enum
	Value float32
}

type OpenHorseScreen struct {
	WindowID  uint8
	SlotCount types.VarInt
	EntityID  int32
}

type InitializeWorldBorder struct {
	X                      float64
	Z                      float64
	OldDiameter            float64
	NewDiameter            float64
	Speed                  types.VarLong
	PortalTeleportBoundary types.VarInt
	WarningBlocks          types.VarInt
	WarningTime            types.VarInt
}

type KeepAlive struct {
	ID int64
}

type BlockEntity struct {
	PackedXZ byte
	Y        int16
	Type     types.VarInt
	Data     NBT // TODO: NBT
}

type ChunkDataAndUpdateLight struct {
	ChunkX              int32
	ChunkZ              int32
	Heightmaps          NBT
	Size                types.VarInt
	Data                []byte
	NumOfBlockEntities  types.VarInt
	BlockEntities       []BlockEntity
	TrustEdges          bool
	SkyLightMask        BitSet // TODO: BitSet
	BlockLightMask      BitSet // TODO: BitSet
	EmptySkyLightMask   BitSet // TODO: BitSet
	EmptyBlockLightMask BitSet // TODO: BitSet
	SkyLightArrayCount  types.VarInt
}
