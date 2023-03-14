package play

import (
	"github.com/google/uuid"
	"minecraftPackets/javaEdition/common"
	"minecraftPackets/nbt"
)

type SpawnEntity struct {
	EntityID   common.VarInt
	ObjectUUID string
	Type       common.VarInt
	X          float64
	Y          float64
	Z          float64
	Pitch      common.Angle
	Yaw        common.Angle
	HeadYaw    common.Angle
	Data       common.VarInt
	VelocityX  int16
	VelocityY  int16
	VelocityZ  int16
}

type SpawnExperienceOrb struct {
	EntityID common.VarInt
	X        float64
	Y        float64
	Z        float64
	Count    int16
}

type SpawnPlayer struct {
	EntityID   common.VarInt
	PlayerUUID uuid.UUID
	X          float64
	Y          float64
	Z          float64
	Yaw        common.Angle
	Pitch      common.Angle
}

type EntityAnimation struct {
	EntityID  common.VarInt
	Animation uint8 // TODO: Enum
}

type AwardStatistic struct {
	CategoryID  common.VarInt
	StatisticID common.VarInt
	Value       common.VarInt
}

type AwardStatistics []AwardStatistic

type AcknowledgeBlockChange struct {
	SequenceID common.VarInt
}

type SetBlockDestroyStage struct {
	EntityID common.VarInt
	Location common.BlockPosition
	Stage    byte // TODO: Enum
}

type BlockEntityData struct {
	Location common.BlockPosition
	Type     common.VarInt
	NBTData  nbt.NBT
}

type BlockAction struct {
	Location    common.BlockPosition
	ActionID    uint8
	ActionParam uint8
	BlockType   common.VarInt
}

type BlockUpdate struct {
	Location common.BlockPosition
	BlockID  common.VarInt
}

type BossBarActionAdd struct {
	Title    common.Chat
	Health   float32
	Color    common.VarInt // TODO: Enum
	Division common.VarInt // TODO: Enum
	Flags    uint8         // TODO: Enum
}

type BossBarActionRemove struct{}

type BossBarActionUpdateHealth struct {
	Health float32
}

type BossBarActionUpdateTitle struct {
	Title common.Chat
}

type BossBarActionUpdateStyle struct {
	Color    common.VarInt // TODO: Enum
	Division common.VarInt // TODO: Enum
}

type BossBarActionUpdateFlags struct {
	Flags uint8 // TODO: Enum
}

type BossBar struct {
	UUID         uuid.UUID
	Action       common.VarInt // TODO: Enum
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
	Tooltip    *common.Chat
}

type CommandSuggestionsResponse struct {
	ID      common.VarInt
	Start   common.VarInt
	Length  common.VarInt
	Count   common.VarInt
	Matches []CommandSuggestionsMatch
}

type Commands struct {
	Count     common.VarInt
	Nodes     []common.CommandNode
	RootIndex common.VarInt
}

type CloseContainer struct {
	WindowID uint8
}

type SetContainerContent struct {
	WindowID    uint8
	StateID     common.VarInt
	Count       common.VarInt
	SlotData    []common.Slot
	CarriedItem common.Slot
}

type SetContainerProperty struct {
	WindowID uint8
	Property int16
	Value    int16
}

type SetContainerSlot struct {
	WindowID uint8
	StateID  common.VarInt
	Slot     int16
	SlotData common.Slot
}

type SetCooldown struct {
	ItemID common.VarInt
	Ticks  common.VarInt
}

type ChatSuggestions struct {
	Action  common.VarInt // TODO: Enum
	Count   common.VarInt
	Entries []string
}

type PluginMessage struct {
	Channel common.Identifier
	Data    []byte
}

type DeleteMessage struct {
	SignatureLength common.VarInt
	Signature       []byte
}

type Disconnect struct {
	Reason common.Chat
}

type DisguisedChatMessage struct {
	Message       common.Chat
	ChatType      common.VarInt // TODO: Enum
	ChatTypeName  common.Chat
	HasTargetName bool
	TargetName    common.Chat
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
	RecordCount   common.VarInt
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
	SlotCount common.VarInt
	EntityID  int32
}

type InitializeWorldBorder struct {
	X                      float64
	Z                      float64
	OldDiameter            float64
	NewDiameter            float64
	Speed                  common.VarLong
	PortalTeleportBoundary common.VarInt
	WarningBlocks          common.VarInt
	WarningTime            common.VarInt
}

type KeepAlive struct {
	ID int64
}

type ChunkDataAndUpdateLight struct {
	ChunkX               int32
	ChunkZ               int32
	Heightmaps           nbt.NBT
	Size                 common.VarInt
	Data                 []byte
	NumOfBlockEntities   common.VarInt
	BlockEntities        []common.BlockEntity
	TrustEdges           bool
	SkyLightMask         common.BitSet
	BlockLightMask       common.BitSet
	EmptySkyLightMask    common.BitSet
	EmptyBlockLightMask  common.BitSet
	SkyLightArrayCount   common.VarInt
	SkyLightArrays       []common.SkyLightArray
	BlockLightArrayCount common.VarInt
	BlockLightArrays     []common.BlockLightArray
}

type WorldEvent struct {
	Event                 int32
	Location              common.BlockPosition
	Data                  int32 // TODO: Enum
	DisableRelativeVolume bool
}

type Particle struct {
	ParticleID     common.VarInt
	IsLongDistance bool
	X              float64
	Y              float64
	Z              float64
	OffsetX        float32
	OffsetY        float32
	OffsetZ        float32
	MaxSpeed       float32
	ParticleCount  int32
	Data           interface{} // TODO: Complete data common
}

type UpdateLight struct {
	ChunkX               common.VarInt
	ChunkZ               common.VarInt
	TrustEdges           bool
	SkyLightMask         common.BitSet
	BlockLightMask       common.BitSet
	EmptySkyLightMask    common.BitSet
	EmptyBlockLightMask  common.BitSet
	SkyLightArrayCount   common.VarInt
	SkyLightArrays       []common.SkyLightArray
	BlockLightArrayCount common.VarInt
	BlockLightArrays     []common.BlockLightArray
}

type Login struct {
	EntityID            int32
	IsHardcore          bool
	GameMode            uint8 // TODO: Enum GameMode
	PreviousGameMode    uint8 // TODO: Enum GameMode
	DimensionCount      common.VarInt
	DimensionNames      []common.Identifier
	RegistryCodec       nbt.NBT
	DimensionType       common.Identifier
	DimensionName       common.Identifier
	HashedSeed          int64
	MaxPlayers          common.VarInt
	ViewDistance        common.VarInt
	SimulationDistance  common.VarInt
	ReducedDebugInfo    bool
	EnableRespawnScreen bool
	IsDebug             bool
	IsFlat              bool
	HasDeathLocation    bool
	DeathDimension      *common.Identifier
	DeathLocation       *common.BlockPosition
}

type Icon struct {
	Type           common.VarInt // TODO: Enum
	X              byte
	Z              byte
	Direction      byte
	HasDisplayName bool
	DisplayName    *common.Chat
}

type MapData struct {
	MapID     common.VarInt
	Scale     uint8
	IsLocked  bool
	HasIcons  bool
	ItemCount *common.VarInt
	Icons     []Icon
	Columns   uint8
	Rows      *uint8
	X         *byte
	Z         *byte
	Length    *common.VarInt
	Data      []byte
}

type Trade struct {
	InputItem1      common.Slot
	OutputItem      common.Slot
	InputItem2      common.Slot
	IsDisabled      bool
	NumOfTradeUses  int32
	MaxTradeUses    int32
	Experience      int32
	SpecialPrice    int32
	PriceMultiplier float32
	Demand          int32
}

type MerchantOffers struct {
	WindowID          common.VarInt
	Size              common.VarInt
	Trades            []Trade
	VillagerLevel     common.VarInt
	Experience        common.VarInt
	IsRegularVillager bool
	CanRestock        bool
}

type UpdateEntityPosition struct {
	EntityID common.VarInt
	DeltaX   int16
	DeltaY   int16
	DeltaZ   int16
	OnGround bool
}

type UpdateEntityPositionAndRotation struct {
	EntityID common.VarInt
	DeltaX   int16
	DeltaY   int16
	DeltaZ   int16
	Yaw      common.Angle
	Pitch    common.Angle
	OnGround bool
}

type UpdateEntityRotation struct {
	EntityID common.VarInt
	Yaw      common.Angle
	Pitch    common.Angle
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
	Hand common.VarInt // TODO: Enum
}

type OpenWindow struct {
	WindowID    common.VarInt
	WindowType  common.VarInt // TODO: Enum
	WindowTitle common.Chat
}

type OpenSignEditor struct {
	Location common.BlockPosition
}

type Ping struct {
	ID int32
}

type PlaceGhostRecipe struct {
	WindowID byte
	RecipeID common.Identifier
}

type PlayerAbilities struct {
	Flags               byte // TODO: Enum
	FlyingSpeed         float32
	FieldOfViewModifier float32
}

type PreviousMessage struct {
	MessageID common.VarInt
	Signature []byte
}

type PlayerChatMessage struct {
	Sender                     uuid.UUID
	Index                      common.VarInt
	MessageSignaturePresent    bool
	MessageSignature           []byte
	Message                    string
	Timestamp                  int64
	Salt                       int64
	TotalPreviousMessages      common.VarInt
	PreviousMessages           []PreviousMessage
	UnsignedContentPresent     bool
	UnsignedContent            *common.Chat
	FilterType                 common.VarInt // TODO: Enum
	FilterTypeBits             *common.BitSet
	ChatType                   common.VarInt // TODO: Enum
	NetworkName                common.Chat
	IsNetworkTargetNamePresent bool
	NetworkTargetName          *common.Chat
}

// EndCombat "Unused by the Notchian client. This data was once used for twitch.tv metadata circa 1.8.f"
type EndCombat struct {
	Duration common.VarInt
	EntityID int32
}

// EnterCombat "Unused by the Notchian client. This data was once used for twitch.tv metadata circa 1.8.f"
type EnterCombat struct{}

type CombatDeath struct {
	PlayerID common.VarInt
	EntityID int32
	Message  common.Chat
}

type PlayerInfoRemove struct {
	NumOfPlayers common.VarInt
	Players      []uuid.UUID
}

type PlayerInfoUpdate struct {
	Action       byte // TODO: Enum
	NumOfActions common.VarInt
	Actions      []common.PlayerInfoAction
}

type LookAt struct {
	FeetOrEyes       common.VarInt // TODO: Enum
	TargetX          float64
	TargetY          float64
	TargetZ          float64
	IsEntity         bool
	EntityID         *common.VarInt
	EntityFeetOrEyes *common.VarInt // TODO: Enum
}

type SynchronizePlayerPosition struct {
	X               float64
	Y               float64
	Z               float64
	Yaw             float32
	Pitch           float32
	Flags           byte // TODO: Enum
	TeleportID      common.VarInt
	DismountVehicle bool
}

type UpdateRecipeBook struct {
	Action                       common.VarInt // TODO: Enum
	CraftingBookOpen             bool
	CraftingBookFilterActive     bool
	SmeltingBookOpen             bool
	SmeltingBookFilterActive     bool
	BlastFurnaceBookOpen         bool
	BlastFurnaceBookFilterActive bool
	SmokerBookOpen               bool
	SmokerBookFilterActive       bool
	RecipeIDSize1                common.VarInt
	RecipeIDs1                   []common.Identifier
	RecipeIDSize2                common.VarInt
	RecipeIDs2                   []common.Identifier
}

type RemoveEntities struct {
	NumOfEntities common.VarInt
	Entities      []common.VarInt
}

type RemoveEntityEffect struct {
	EntityID common.VarInt
	EffectID common.VarInt // TODO: Enum
}

type ResourcePack struct {
	URL              string
	Hash             string
	Forced           bool
	HasPromptMessage bool
	PromptMessage    *common.Chat
}

type Respawn struct {
	DimensionType      common.Identifier
	DimensionName      common.Identifier
	HashedSeed         int64
	GameMode           uint8 // TODO: Enum GameMode
	PreviousGameMode   byte
	IsDebug            bool
	IsFlat             bool
	CopyMetadata       bool
	HasDeathLocation   bool
	DeathDimensionName *common.Identifier
	DeathLocation      *common.BlockPosition
}

type SetHeadRotation struct {
	EntityID common.VarInt
	HeadYaw  common.Angle
}

type UpdateSectionBlocks struct {
	ChunkSectionPosition int64
	SuppressLightUpdates bool
	BlockArraySize       common.VarInt
	Blocks               []common.VarLong
}

type SelectAdvancementTab struct {
	HasID bool
	ID    *common.Identifier
}

type ServerData struct {
	HasMOTD            bool
	MOTD               *common.Chat
	HasIcon            bool
	Icon               *string
	EnforcesSecureChat bool
}

type SetActionBarText struct {
	Text common.Chat
}

type SetBorderCenter struct {
	X float64
	Z float64
}

type SetBorderLerpSize struct {
	OldDiameter float64
	NewDiameter float64
	Speed       common.VarLong
}

type SetBorderSize struct {
	Diameter float64
}

type SetBorderWarningDelay struct {
	WarningTime common.VarInt
}

type SetBorderWarningDistance struct {
	// WarningBlocks in meters
	WarningBlocks common.VarInt
}

type SetCamera struct {
	CameraID common.VarInt
}

type SetHeldItem struct {
	Slot byte
}

type SetCenterChunk struct {
	ChunkX int32
	ChunkZ int32
}

type SetRenderDistance struct {
	ViewDistance common.VarInt
}

type SetDefaultSpawnPosition struct {
	Location common.BlockPosition
	Angle    float32
}

type DisplayObjective struct {
	Position common.BlockPosition
	Angle    float32
}

type SetEntityMetadata struct {
	EntityID common.VarInt
	//Metadata []common.EntityMetadata // TODO: EntityMetadata
}

type LinkEntities struct {
	AttachedEntityID int32
	HoldingEntityID  int32
}

type SetEntityVelocity struct {
	EntityID  common.VarInt
	VelocityX int16
	VelocityY int16
	VelocityZ int16
}

type SetEquipment struct {
	EntityID common.VarInt
	//Equipment common.Equipment // TODO: Equipment
}

type SetExperience struct {
	ExperienceBar   float32
	TotalExperience common.VarInt
	Level           common.VarInt
}

type SetHealth struct {
	Health         float32
	Food           common.VarInt
	FoodSaturation float32
}

type UpdateObjectives struct {
	ObjectiveName  string
	Mode           byte
	ObjectiveValue *common.Chat
	Type           *common.VarInt // TODO: Enum
}

type SetPassengers struct {
	EntityID       common.VarInt
	PassengerCount common.VarInt
	Passengers     []common.VarInt
}

type UpdateTeams struct {
	TeamName string
	Mode     byte
	// TODO: Complete
}

type UpdateScore struct {
	EntityName    string
	Action        common.VarInt // TODO: Enum
	ObjectiveName string
	Value         *common.VarInt
}

type SetSimulationDistance struct {
	SimulationDistance common.VarInt
}

type SetSubtitleText struct {
	SubtitleText common.Chat
}

type UpdateTime struct {
	WorldAge  int64
	TimeOfDay int64
}

type SetTitleText struct {
	TitleText common.Chat
}

type SetTitleAnimationTimes struct {
	FadeIn  int32
	Stay    int32
	FadeOut int32
}

type EntitySoundEffect struct {
	SoundID       common.VarInt // TODO: Enum
	SoundName     *common.Identifier
	HasFixedRange *bool
	Range         *float32
	SoundCategory common.VarInt // TODO: Enum
	EntityID      common.VarInt
	Volume        float32
	Pitch         float32
	Seed          int64
}

type SoundEffect struct {
	SoundID         common.VarInt
	SoundName       *common.Identifier
	HasFixedRange   *bool
	Range           *float32
	SoundCategory   common.VarInt // TODO: Enum
	EntityPositionX int32
	EntityPositionY int32
	EntityPositionZ int32
	Volume          float32
	Pitch           float32
	Seed            int64
}

type StopSound struct {
	Flags  byte
	Source *common.VarInt // TODO: Enum
	Sound  *common.Identifier
}

type SystemChatMessage struct {
	Content common.Chat
	Overlay bool
}

type SetTabListHeaderAndFooter struct {
	Header common.Chat
	Footer common.Chat
}

type TagQueryResponse struct {
	TransactionID common.VarInt
	NBT           nbt.NBT
}

type PickupItem struct {
	CollectedEntityID common.VarInt
	CollectorEntityID common.VarInt
	PickupItemCount   common.VarInt
}

type TeleportEntity struct {
	EntityID common.VarInt
	X        float64
	Y        float64
	Z        float64
	Yaw      common.Angle
	Pitch    common.Angle
	OnGround bool
}

type UpdateAdvancements struct {
	ResetOrClear bool
	MappingSize  common.VarInt
	//AdvancementMapping map[common.Identifier]Advancement // TODO: Advancement
	ListSize     common.VarInt
	Identifiers  []common.Identifier
	ProgressSize common.VarInt
	//ProgressMapping    map[common.Identifier]AdvancementProgress // TODO: AdvancementProgress
}

type UpdateAttributes struct {
	EntityID        common.VarInt
	NumOfProperties common.VarInt
	//Properties      []AttributeProperty // TODO: AttributeProperty
}

type FeatureFlags struct {
	TotalFeatures common.VarInt
	FeatureFlags  []common.Identifier
}

type EntityEffect struct {
	EntityID      common.VarInt
	EffectID      common.VarInt // TODO: Enum
	Amplifier     byte
	Duration      common.VarInt
	Flags         byte
	HasFactorData bool
	FactorCodec   nbt.NBT
}

type UpdateRecipes struct {
	RecipeCount common.VarInt
	//Recipes     []Recipe // TODO: Recipe
}

type UpdateTags struct {
	TagCount common.VarInt
	//Tags     []Tag // TODO: Tag
}
