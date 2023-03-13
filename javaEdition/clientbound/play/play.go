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

type SetContainerContent struct {
	WindowID    uint8
	StateID     types.VarInt
	Count       types.VarInt
	SlotData    []types.Slot
	CarriedItem types.Slot
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
	SlotData types.Slot
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
	InputItem1      types.Slot
	OutputItem      types.Slot
	InputItem2      types.Slot
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

type RemoveEntities struct {
	NumOfEntities types.VarInt
	Entities      []types.VarInt
}

type RemoveEntityEffect struct {
	EntityID types.VarInt
	EffectID types.VarInt // TODO: Enum
}

type ResourcePack struct {
	URL              string
	Hash             string
	Forced           bool
	HasPromptMessage bool
	PromptMessage    *types.Chat
}

type Respawn struct {
	DimensionType      types.Identifier
	DimensionName      types.Identifier
	HashedSeed         int64
	GameMode           uint8 // TODO: Enum GameMode
	PreviousGameMode   byte
	IsDebug            bool
	IsFlat             bool
	CopyMetadata       bool
	HasDeathLocation   bool
	DeathDimensionName *types.Identifier
	DeathLocation      *types.BlockPosition
}

type SetHeadRotation struct {
	EntityID types.VarInt
	HeadYaw  types.Angle
}

type UpdateSectionBlocks struct {
	ChunkSectionPosition int64
	SuppressLightUpdates bool
	BlockArraySize       types.VarInt
	Blocks               []types.VarLong
}

type SelectAdvancementTab struct {
	HasID bool
	ID    *types.Identifier
}

type ServerData struct {
	HasMOTD            bool
	MOTD               *types.Chat
	HasIcon            bool
	Icon               *string
	EnforcesSecureChat bool
}

type SetActionBarText struct {
	Text types.Chat
}

type SetBorderCenter struct {
	X float64
	Z float64
}

type SetBorderLerpSize struct {
	OldDiameter float64
	NewDiameter float64
	Speed       types.VarLong
}

type SetBorderSize struct {
	Diameter float64
}

type SetBorderWarningDelay struct {
	WarningTime types.VarInt
}

type SetBorderWarningDistance struct {
	// WarningBlocks in meters
	WarningBlocks types.VarInt
}

type SetCamera struct {
	CameraID types.VarInt
}

type SetHeldItem struct {
	Slot byte
}

type SetCenterChunk struct {
	ChunkX int32
	ChunkZ int32
}

type SetRenderDistance struct {
	ViewDistance types.VarInt
}

type SetDefaultSpawnPosition struct {
	Location types.BlockPosition
	Angle    float32
}

type DisplayObjective struct {
	Position types.BlockPosition
	Angle    float32
}

type SetEntityMetadata struct {
	EntityID types.VarInt
	Metadata []types.EntityMetadata // TODO: EntityMetadata
}

type LinkEntities struct {
	AttachedEntityID int32
	HoldingEntityID  int32
}

type SetEntityVelocity struct {
	EntityID  types.VarInt
	VelocityX int16
	VelocityY int16
	VelocityZ int16
}

type SetEquipment struct {
	EntityID  types.VarInt
	Equipment types.Equipment // TODO: Equipment
}

type SetExperience struct {
	ExperienceBar   float32
	TotalExperience types.VarInt
	Level           types.VarInt
}

type SetHealth struct {
	Health         float32
	Food           types.VarInt
	FoodSaturation float32
}

type UpdateObjectives struct {
	ObjectiveName  string
	Mode           byte
	ObjectiveValue *types.Chat
	Type           *types.VarInt // TODO: Enum
}

type SetPassengers struct {
	EntityID       types.VarInt
	PassengerCount types.VarInt
	Passengers     []types.VarInt
}

type UpdateTeams struct {
	TeamName string
	Mode     byte
	// TODO: Complete
}

type UpdateScore struct {
	EntityName    string
	Action        types.VarInt // TODO: Enum
	ObjectiveName string
	Value         *types.VarInt
}

type SetSimulationDistance struct {
	SimulationDistance types.VarInt
}

type SetSubtitleText struct {
	SubtitleText types.Chat
}

type UpdateTime struct {
	WorldAge  int64
	TimeOfDay int64
}

type SetTitleText struct {
	TitleText types.Chat
}

type SetTitleAnimationTimes struct {
	FadeIn  int32
	Stay    int32
	FadeOut int32
}

type EntitySoundEffect struct {
	SoundID       types.VarInt // TODO: Enum
	SoundName     *types.Identifier
	HasFixedRange *bool
	Range         *float32
	SoundCategory types.VarInt // TODO: Enum
	EntityID      types.VarInt
	Volume        float32
	Pitch         float32
	Seed          int64
}

type SoundEffect struct {
	SoundID         types.VarInt
	SoundName       *types.Identifier
	HasFixedRange   *bool
	Range           *float32
	SoundCategory   types.VarInt // TODO: Enum
	EntityPositionX int32
	EntityPositionY int32
	EntityPositionZ int32
	Volume          float32
	Pitch           float32
	Seed            int64
}

type StopSound struct {
	Flags  byte
	Source *types.VarInt // TODO: Enum
	Sound  *types.Identifier
}

type SystemChatMessage struct {
	Content types.Chat
	Overlay bool
}

type SetTabListHeaderAndFooter struct {
	Header types.Chat
	Footer types.Chat
}

type TagQueryResponse struct {
	TransactionID types.VarInt
	NBT           types.NBT // TODO: NBT
}

type PickupItem struct {
	CollectedEntityID types.VarInt
	CollectorEntityID types.VarInt
	PickupItemCount   types.VarInt
}

type TeleportEntity struct {
	EntityID types.VarInt
	X        float64
	Y        float64
	Z        float64
	Yaw      types.Angle
	Pitch    types.Angle
	OnGround bool
}

type UpdateAdvancements struct {
	ResetOrClear       bool
	MappingSize        types.VarInt
	AdvancementMapping map[types.Identifier]Advancement // TODO: Advancement
	ListSize           types.VarInt
	Identifiers        []types.Identifier
	ProgressSize       types.VarInt
	ProgressMapping    map[types.Identifier]AdvancementProgress // TODO: AdvancementProgress
}

type UpdateAttributes struct {
	EntityID        types.VarInt
	NumOfProperties types.VarInt
	Properties      []AttributeProperty // TODO: AttributeProperty
}

type FeatureFlags struct {
	TotalFeatures types.VarInt
	FeatureFlags  []types.Identifier
}

type EntityEffect struct {
	EntityID      types.VarInt
	EffectID      types.VarInt // TODO: Enum
	Amplifier     byte
	Duration      types.VarInt
	Flags         byte
	HasFactorData bool
	FactorCodec   NBT // TODO: NBT
}

type UpdateRecipes struct {
	RecipeCount types.VarInt
	Recipes     []Recipe // TODO: Recipe
}

type UpdateTags struct {
	TagCount types.VarInt
	Tags     []Tag // TODO: Tag
}
