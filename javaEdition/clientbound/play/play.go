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
	// TODO: Complete
}

type WorldEvent struct {
	Event                 int32
	Location              types.BlockPosition
	Data                  int32 // TODO: Enum
	DisableRelativeVolume bool
}

type Particle struct {
	ParticleID     types.VarInt
	IsLongDistance bool
	X              float64
	Y              float64
	Z              float64
	OffsetX        float32
	OffsetY        float32
	OffsetZ        float32
	MaxSpeed       float32
	ParticleCount  int32
	Data           interface{} // TODO: Complete data types
}

type UpdateLight struct {
	ChunkX              types.VarInt
	ChunkZ              types.VarInt
	TrustEdges          bool
	SkyLightMask        BitSet // TODO: BitSet
	BlockLightMask      BitSet // TODO: BitSet
	EmptySkyLightMask   BitSet // TODO: BitSet
	EmptyBlockLightMask BitSet // TODO: BitSet
	SkyLightArrayCount  types.VarInt
	// TODO: Complete
}

type Login struct {
	EntityID            int32
	IsHardcore          bool
	GameMode            uint8 // TODO: Enum GameMode
	PreviousGameMode    uint8 // TODO: Enum GameMode
	DimensionCount      types.VarInt
	DimensionNames      []types.Identifier
	RegistryCodec       NBT // TODO: NBT
	DimensionType       types.Identifier
	DimensionName       types.Identifier
	HashedSeed          int64
	MaxPlayers          types.VarInt
	ViewDistance        types.VarInt
	SimulationDistance  types.VarInt
	ReducedDebugInfo    bool
	EnableRespawnScreen bool
	IsDebug             bool
	IsFlat              bool
	HasDeathLocation    bool
	DeathDimension      *types.Identifier
	DeathLocation       *types.BlockPosition
}

type Icon struct {
	Type           types.VarInt // TODO: Enum
	X              byte
	Z              byte
	Direction      byte
	HasDisplayName bool
	DisplayName    *types.Chat
}

type MapData struct {
	MapID     types.VarInt
	Scale     uint8
	IsLocked  bool
	HasIcons  bool
	ItemCount *types.VarInt
	Icons     []Icon
	Columns   uint8
	Rows      *uint8
	X         *byte
	Z         *byte
	Length    *types.VarInt
	Data      []byte
}

type Trade struct {
	InputItem1      Slot
	OutputItem      Slot
	InputItem2      Slot
	IsDisabled      bool
	NumOfTradeUses  int32
	MaxTradeUses    int32
	Experience      int32
	SpecialPrice    int32
	PriceMultiplier float32
	Demand          int32
}

type MerchantOffers struct {
	WindowID          types.VarInt
	Size              types.VarInt
	Trades            []Trade
	VillagerLevel     types.VarInt
	Experience        types.VarInt
	IsRegularVillager bool
	CanRestock        bool
}

type UpdateEntityPosition struct {
	EntityID types.VarInt
	DeltaX   int16
	DeltaY   int16
	DeltaZ   int16
	OnGround bool
}

type UpdateEntityPositionAndRotation struct {
	EntityID types.VarInt
	DeltaX   int16
	DeltaY   int16
	DeltaZ   int16
	Yaw      types.Angle
	Pitch    types.Angle
	OnGround bool
}

type UpdateEntityRotation struct {
	EntityID types.VarInt
	Yaw      types.Angle
	Pitch    types.Angle
	OnGround bool
}

type MoveVehicle struct {
	X     float64
	Y     float64
	Z     float64
	Yaw   float32
	Pitch float32
}

type OpenBook struct {
	Hand types.VarInt // TODO: Enum
}

type OpenWindow struct {
	WindowID    types.VarInt
	WindowType  types.VarInt // TODO: Enum
	WindowTitle types.Chat
}

type OpenSignEditor struct {
	Location types.BlockPosition
}

type Ping struct {
	ID int32
}

type PlaceGhostRecipe struct {
	WindowID byte
	RecipeID types.Identifier
}

type PlayerAbilities struct {
	Flags               byte // TODO: Enum
	FlyingSpeed         float32
	FieldOfViewModifier float32
}

type PreviousMessage struct {
	MessageID types.VarInt
	Signature []byte
}

type PlayerChatMessage struct {
	Sender                     uuid.UUID
	Index                      types.VarInt
	MessageSignaturePresent    bool
	MessageSignature           []byte
	Message                    string
	Timestamp                  int64
	Salt                       int64
	TotalPreviousMessages      types.VarInt
	PreviousMessages           []PreviousMessage
	UnsignedContentPresent     bool
	UnsignedContent            *types.Chat
	FilterType                 types.VarInt // TODO: Enum
	FilterTypeBits             *BitSet      // TODO: BitSet
	ChatType                   types.VarInt // TODO: Enum
	NetworkName                types.Chat
	IsNetworkTargetNamePresent bool
	NetworkTargetName          *types.Chat
}

// EndCombat "Unused by the Notchian client. This data was once used for twitch.tv metadata circa 1.8.f"
type EndCombat struct {
	Duration types.VarInt
	EntityID int32
}

// EnterCombat "Unused by the Notchian client. This data was once used for twitch.tv metadata circa 1.8.f"
type EnterCombat struct{}

type CombatDeath struct {
	PlayerID types.VarInt
	EntityID int32
	Message  types.Chat
}

type PlayerInfoRemove struct {
	NumOfPlayers types.VarInt
	Players      []uuid.UUID
}

type PlayerInfoUpdate struct {
	// TODO: Complete
}

type LookAt struct {
	FeetOrEyes       types.VarInt // TODO: Enum
	TargetX          float64
	TargetY          float64
	TargetZ          float64
	IsEntity         bool
	EntityID         *types.VarInt
	EntityFeetOrEyes *types.VarInt // TODO: Enum
}

type SynchronizePlayerPosition struct {
	X               float64
	Y               float64
	Z               float64
	Yaw             float32
	Pitch           float32
	Flags           byte // TODO: Enum
	TeleportID      types.VarInt
	DismountVehicle bool
}

type UpdateRecipeBook struct {
	Action                       types.VarInt // TODO: Enum
	CraftingBookOpen             bool
	CraftingBookFilterActive     bool
	SmeltingBookOpen             bool
	SmeltingBookFilterActive     bool
	BlastFurnaceBookOpen         bool
	BlastFurnaceBookFilterActive bool
	SmokerBookOpen               bool
	SmokerBookFilterActive       bool
	RecipeIDSize1                types.VarInt
	RecipeIDs1                   []types.Identifier
	RecipeIDSize2                types.VarInt
	RecipeIDs2                   []types.Identifier
}
