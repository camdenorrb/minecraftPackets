package play

import (
	"github.com/camdenorrb/minecraftPackets/javaEdition/common"
	"github.com/camdenorrb/minecraftPackets/nbt"
	"github.com/camdenorrb/minecraftPackets/primitive"
	"github.com/google/uuid"
)

type SpawnEntity struct {
	EntityID   primitive.VarInt
	ObjectUUID string
	Type       primitive.VarInt
	X          float64
	Y          float64
	Z          float64
	Pitch      common.Angle
	Yaw        common.Angle
	HeadYaw    common.Angle
	Data       primitive.VarInt
	VelocityX  int16
	VelocityY  int16
	VelocityZ  int16
}

type SpawnExperienceOrb struct {
	EntityID primitive.VarInt
	X        float64
	Y        float64
	Z        float64
	Count    int16
}

type SpawnPlayer struct {
	EntityID   primitive.VarInt
	PlayerUUID uuid.UUID
	X          float64
	Y          float64
	Z          float64
	Yaw        common.Angle
	Pitch      common.Angle
}

type EntityAnimation struct {
	EntityID  primitive.VarInt
	Animation uint8 // TODO: Enum
}

type AwardStatistics struct {
	Count      primitive.VarInt
	Statistics common.AwardStatistics
}

type AcknowledgeBlockChange struct {
	SequenceID primitive.VarInt
}

type SetBlockDestroyStage struct {
	EntityID primitive.VarInt
	Location common.BlockPosition
	Stage    byte // TODO: Enum
}

type BlockEntityData struct {
	Location common.BlockPosition
	Type     primitive.VarInt
	NBTData  nbt.NBT
}

type BlockAction struct {
	Location    common.BlockPosition
	ActionID    uint8
	ActionParam uint8
	BlockType   primitive.VarInt
}

type BlockUpdate struct {
	Location common.BlockPosition
	BlockID  primitive.VarInt
}

type BossBarActionAdd struct {
	Title    common.Chat
	Health   float32
	Color    primitive.VarInt // TODO: Enum
	Division primitive.VarInt // TODO: Enum
	Flags    uint8            // TODO: Enum
}

type BossBarActionRemove struct{}

type BossBarActionUpdateHealth struct {
	Health float32
}

type BossBarActionUpdateTitle struct {
	Title common.Chat
}

type BossBarActionUpdateStyle struct {
	Color    primitive.VarInt // TODO: Enum
	Division primitive.VarInt // TODO: Enum
}

type BossBarActionUpdateFlags struct {
	Flags uint8 // TODO: Enum
}

type BossBar struct {
	UUID         uuid.UUID
	Action       primitive.VarInt // TODO: Enum
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
	ID      primitive.VarInt
	Start   primitive.VarInt
	Length  primitive.VarInt
	Count   primitive.VarInt
	Matches []CommandSuggestionsMatch
}

type Commands struct {
	Count     primitive.VarInt
	Nodes     []common.CommandNode
	RootIndex primitive.VarInt
}

type CloseContainer struct {
	WindowID uint8
}

type SetContainerContent struct {
	WindowID    uint8
	StateID     primitive.VarInt
	Count       primitive.VarInt
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
	StateID  primitive.VarInt
	Slot     int16
	SlotData common.Slot
}

type SetCooldown struct {
	ItemID primitive.VarInt
	Ticks  primitive.VarInt
}

type ChatSuggestions struct {
	Action  primitive.VarInt // TODO: Enum
	Count   primitive.VarInt
	Entries []string
}

type PluginMessage struct {
	Channel primitive.Identifier
	Data    []byte
}

type DeleteMessage struct {
	SignatureLength primitive.VarInt
	Signature       []byte
}

type Disconnect struct {
	Reason common.Chat
}

type DisguisedChatMessage struct {
	Message       common.Chat
	ChatType      primitive.VarInt // TODO: Enum
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
	RecordCount   primitive.VarInt
	Records       [][]byte
	PlayerMotionX float32
	PlayerMotionY float32
	PlayerMotionZ float32
}

type UnloadChunk struct {
	ChunkX int32
	ChunkZ int32
}

type GameEvent struct {
	Event uint8 // TODO: Enum
	Value float32
}

type OpenHorseScreen struct {
	WindowID  uint8
	SlotCount primitive.VarInt
	EntityID  int32
}

type InitializeWorldBorder struct {
	X                      float64
	Z                      float64
	OldDiameter            float64
	NewDiameter            float64
	Speed                  primitive.VarLong
	PortalTeleportBoundary primitive.VarInt
	WarningBlocks          primitive.VarInt
	WarningTime            primitive.VarInt
}

type KeepAlive struct {
	ID int64
}

type ChunkDataAndUpdateLight struct {
	ChunkX               int32
	ChunkZ               int32
	Heightmaps           nbt.NBT
	Size                 primitive.VarInt
	Data                 []common.ChunkSection
	NumOfBlockEntities   primitive.VarInt
	BlockEntities        []common.BlockEntity
	TrustEdges           bool
	SkyLightMask         primitive.BitSet
	BlockLightMask       primitive.BitSet
	EmptySkyLightMask    primitive.BitSet
	EmptyBlockLightMask  primitive.BitSet
	SkyLightArrayCount   primitive.VarInt
	SkyLightArrays       []common.SkyLightArray
	BlockLightArrayCount primitive.VarInt
	BlockLightArrays     []common.BlockLightArray
}

type WorldEvent struct {
	Event                 int32
	Location              common.BlockPosition
	Data                  int32 // TODO: Enum
	DisableRelativeVolume bool
}

type Particle struct {
	ParticleID     primitive.VarInt
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
	ChunkX               primitive.VarInt
	ChunkZ               primitive.VarInt
	TrustEdges           bool
	SkyLightMask         primitive.BitSet
	BlockLightMask       primitive.BitSet
	EmptySkyLightMask    primitive.BitSet
	EmptyBlockLightMask  primitive.BitSet
	SkyLightArrayCount   primitive.VarInt
	SkyLightArrays       []common.SkyLightArray
	BlockLightArrayCount primitive.VarInt
	BlockLightArrays     []common.BlockLightArray
}

type Login struct {
	EntityID            int32
	IsHardcore          bool
	GameMode            uint8 // TODO: Enum GameMode
	PreviousGameMode    uint8 // TODO: Enum GameMode
	DimensionCount      primitive.VarInt
	DimensionNames      []primitive.Identifier
	RegistryCodec       nbt.NBT
	DimensionType       primitive.Identifier
	DimensionName       primitive.Identifier
	HashedSeed          int64
	MaxPlayers          primitive.VarInt
	ViewDistance        primitive.VarInt
	SimulationDistance  primitive.VarInt
	ReducedDebugInfo    bool
	EnableRespawnScreen bool
	IsDebug             bool
	IsFlat              bool
	HasDeathLocation    bool
	DeathDimension      *primitive.Identifier
	DeathLocation       *common.BlockPosition
	PortalCooldown      primitive.VarInt
}

type Icon struct {
	Type           primitive.VarInt // TODO: Enum
	X              byte
	Z              byte
	Direction      byte
	HasDisplayName bool
	DisplayName    *common.Chat
}

type MapData struct {
	MapID     primitive.VarInt
	Scale     uint8
	IsLocked  bool
	HasIcons  bool
	ItemCount *primitive.VarInt
	Icons     []Icon
	Columns   uint8
	Rows      *uint8
	X         *byte
	Z         *byte
	Length    *primitive.VarInt
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
	WindowID          primitive.VarInt
	Size              primitive.VarInt
	Trades            []Trade
	VillagerLevel     primitive.VarInt
	Experience        primitive.VarInt
	IsRegularVillager bool
	CanRestock        bool
}

type UpdateEntityPosition struct {
	EntityID primitive.VarInt
	DeltaX   int16
	DeltaY   int16
	DeltaZ   int16
	OnGround bool
}

type UpdateEntityPositionAndRotation struct {
	EntityID primitive.VarInt
	DeltaX   int16
	DeltaY   int16
	DeltaZ   int16
	Yaw      common.Angle
	Pitch    common.Angle
	OnGround bool
}

type UpdateEntityRotation struct {
	EntityID primitive.VarInt
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
	Hand common.Hand
}

type OpenWindow struct {
	WindowID    primitive.VarInt
	WindowType  primitive.VarInt // TODO: Enum
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
	RecipeID primitive.Identifier
}

type PlayerAbilities struct {
	Flags               byte // TODO: Enum
	FlyingSpeed         float32
	FieldOfViewModifier float32
}

type PreviousMessage struct {
	MessageID primitive.VarInt
	Signature []byte
}

type PlayerChatMessage struct {
	Sender                     uuid.UUID
	Index                      primitive.VarInt
	MessageSignaturePresent    bool
	MessageSignature           []byte
	Message                    string
	Timestamp                  int64
	Salt                       int64
	TotalPreviousMessages      primitive.VarInt
	PreviousMessages           []PreviousMessage
	UnsignedContentPresent     bool
	UnsignedContent            *common.Chat
	FilterType                 primitive.VarInt // TODO: Enum
	FilterTypeBits             *primitive.BitSet
	ChatType                   primitive.VarInt // TODO: Enum
	NetworkName                common.Chat
	IsNetworkTargetNamePresent bool
	NetworkTargetName          *common.Chat
}

// EndCombat "Unused by the Notchian client. This data was once used for twitch.tv metadata circa 1.8.f"
type EndCombat struct {
	Duration primitive.VarInt
	EntityID int32
}

// EnterCombat "Unused by the Notchian client. This data was once used for twitch.tv metadata circa 1.8.f"
type EnterCombat struct{}

type CombatDeath struct {
	PlayerID primitive.VarInt
	EntityID int32
	Message  common.Chat
}

type PlayerInfoRemove struct {
	NumOfPlayers primitive.VarInt
	Players      []uuid.UUID
}

type PlayerInfoUpdate struct {
	Action       byte // TODO: Enum
	NumOfActions primitive.VarInt
	Actions      []common.PlayerInfoAction
}

type LookAt struct {
	FeetOrEyes       primitive.VarInt // TODO: Enum
	TargetX          float64
	TargetY          float64
	TargetZ          float64
	IsEntity         bool
	EntityID         *primitive.VarInt
	EntityFeetOrEyes *primitive.VarInt // TODO: Enum
}

type SynchronizePlayerPosition struct {
	X               float64
	Y               float64
	Z               float64
	Yaw             float32
	Pitch           float32
	Flags           byte // TODO: Enum
	TeleportID      primitive.VarInt
	DismountVehicle bool
}

type UpdateRecipeBook struct {
	Action                       primitive.VarInt // TODO: Enum
	CraftingBookOpen             bool
	CraftingBookFilterActive     bool
	SmeltingBookOpen             bool
	SmeltingBookFilterActive     bool
	BlastFurnaceBookOpen         bool
	BlastFurnaceBookFilterActive bool
	SmokerBookOpen               bool
	SmokerBookFilterActive       bool
	RecipeIDSize1                primitive.VarInt
	RecipeIDs1                   []primitive.Identifier
	RecipeIDSize2                primitive.VarInt
	RecipeIDs2                   []primitive.Identifier
}

type RemoveEntities struct {
	NumOfEntities primitive.VarInt
	Entities      []primitive.VarInt
}

type RemoveEntityEffect struct {
	EntityID primitive.VarInt
	EffectID primitive.VarInt // TODO: Enum
}

type ResourcePack struct {
	URL              string
	Hash             string
	Forced           bool
	HasPromptMessage bool
	PromptMessage    *common.Chat
}

type Respawn struct {
	DimensionType      primitive.Identifier
	DimensionName      primitive.Identifier
	HashedSeed         int64
	GameMode           uint8 // TODO: Enum GameMode
	PreviousGameMode   byte
	IsDebug            bool
	IsFlat             bool
	CopyMetadata       bool
	HasDeathLocation   bool
	DeathDimensionName *primitive.Identifier
	DeathLocation      *common.BlockPosition
}

type SetHeadRotation struct {
	EntityID primitive.VarInt
	HeadYaw  common.Angle
}

type UpdateSectionBlocks struct {
	ChunkSectionPosition int64
	SuppressLightUpdates bool
	BlockArraySize       primitive.VarInt
	Blocks               []primitive.VarLong
}

type SelectAdvancementTab struct {
	HasID bool
	ID    *primitive.Identifier
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
	Speed       primitive.VarLong
}

type SetBorderSize struct {
	Diameter float64
}

type SetBorderWarningDelay struct {
	WarningTime primitive.VarInt
}

type SetBorderWarningDistance struct {
	// WarningBlocks in meters
	WarningBlocks primitive.VarInt
}

type SetCamera struct {
	CameraID primitive.VarInt
}

type SetHeldItem struct {
	Slot byte
}

type SetCenterChunk struct {
	ChunkX int32
	ChunkZ int32
}

type SetRenderDistance struct {
	ViewDistance primitive.VarInt
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
	EntityID primitive.VarInt
	//Metadata []common.EntityMetadata // TODO: EntityMetadata
}

type LinkEntities struct {
	AttachedEntityID int32
	HoldingEntityID  int32
}

type SetEntityVelocity struct {
	EntityID  primitive.VarInt
	VelocityX int16
	VelocityY int16
	VelocityZ int16
}

type SetEquipment struct {
	EntityID primitive.VarInt
	//Equipment common.Equipment // TODO: Equipment
}

type SetExperience struct {
	ExperienceBar   float32
	TotalExperience primitive.VarInt
	Level           primitive.VarInt
}

type SetHealth struct {
	Health         float32
	Food           primitive.VarInt
	FoodSaturation float32
}

type UpdateObjectives struct {
	ObjectiveName  string
	Mode           byte
	ObjectiveValue *common.Chat
	Type           *primitive.VarInt // TODO: Enum
}

type SetPassengers struct {
	EntityID       primitive.VarInt
	PassengerCount primitive.VarInt
	Passengers     []primitive.VarInt
}

type UpdateTeams struct {
	TeamName string
	Mode     byte
	// TODO: Complete
}

type UpdateScore struct {
	EntityName    string
	Action        primitive.VarInt // TODO: Enum
	ObjectiveName string
	Value         *primitive.VarInt
}

type SetSimulationDistance struct {
	SimulationDistance primitive.VarInt
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
	SoundID       primitive.VarInt // TODO: Enum
	SoundName     *primitive.Identifier
	HasFixedRange *bool
	Range         *float32
	SoundCategory primitive.VarInt // TODO: Enum
	EntityID      primitive.VarInt
	Volume        float32
	Pitch         float32
	Seed          int64
}

type SoundEffect struct {
	SoundID         primitive.VarInt
	SoundName       *primitive.Identifier
	HasFixedRange   *bool
	Range           *float32
	SoundCategory   primitive.VarInt // TODO: Enum
	EntityPositionX int32
	EntityPositionY int32
	EntityPositionZ int32
	Volume          float32
	Pitch           float32
	Seed            int64
}

type StopSound struct {
	Flags  byte
	Source *primitive.VarInt // TODO: Enum
	Sound  *primitive.Identifier
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
	TransactionID primitive.VarInt
	NBT           nbt.NBT
}

type PickupItem struct {
	CollectedEntityID primitive.VarInt
	CollectorEntityID primitive.VarInt
	PickupItemCount   primitive.VarInt
}

type TeleportEntity struct {
	EntityID primitive.VarInt
	X        float64
	Y        float64
	Z        float64
	Yaw      common.Angle
	Pitch    common.Angle
	OnGround bool
}

type UpdateAdvancements struct {
	ResetOrClear bool
	MappingSize  primitive.VarInt
	//AdvancementMapping map[common.Identifier]Advancement // TODO: Advancement
	ListSize     primitive.VarInt
	Identifiers  []primitive.Identifier
	ProgressSize primitive.VarInt
	//ProgressMapping    map[common.Identifier]AdvancementProgress // TODO: AdvancementProgress
}

type UpdateAttributes struct {
	EntityID        primitive.VarInt
	NumOfProperties primitive.VarInt
	//Properties      []AttributeProperty // TODO: AttributeProperty
}

type FeatureFlags struct {
	TotalFeatures primitive.VarInt
	FeatureFlags  []primitive.Identifier
}

type EntityEffect struct {
	EntityID      primitive.VarInt
	EffectID      primitive.VarInt // TODO: Enum
	Amplifier     byte
	Duration      primitive.VarInt
	Flags         byte
	HasFactorData bool
	FactorCodec   nbt.NBT
}

type UpdateRecipes struct {
	RecipeCount primitive.VarInt
	//Recipes     []Recipe // TODO: Recipe
}

type UpdateTags struct {
	TagCount primitive.VarInt
	//Tags     []Tag // TODO: Tag
}
