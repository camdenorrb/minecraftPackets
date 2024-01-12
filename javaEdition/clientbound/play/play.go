package play

import (
	"github.com/camdenorrb/minecraftPackets/javaEdition/common"
	common2 "github.com/camdenorrb/minecraftPackets/javaEdition/protocol/common"
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

func (*SpawnEntity) PacketType() string {
	return string(common2.ClientBoundPlaySpawnEntity)
}

type SpawnExperienceOrb struct {
	EntityID primitive.VarInt
	X        float64
	Y        float64
	Z        float64
	Count    int16
}

func (*SpawnExperienceOrb) PacketType() string {
	return string(common2.ClientBoundPlaySpawnExperienceOrb)
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

func (*SpawnPlayer) PacketType() string {
	return string(common2.ClientBoundPlaySpawnPlayer)
}

type EntityAnimation struct {
	EntityID  primitive.VarInt
	Animation uint8 // TODO: Enum
}

func (*EntityAnimation) PacketType() string {
	return string(common2.ClientBoundPlayEntityAnimation)
}

type AwardStatistics struct {
	Count      primitive.VarInt
	Statistics common.AwardStatistics
}

func (*AwardStatistics) PacketType() string {
	return string(common2.ClientBoundPlayAwardStatistics)
}

type AcknowledgeBlockChange struct {
	SequenceID primitive.VarInt
}

func (*AcknowledgeBlockChange) PacketType() string {
	return string(common2.ClientBoundPlayAcknowledgeBlockChange)
}

type SetBlockDestroyStage struct {
	EntityID primitive.VarInt
	Location common.BlockPosition
	Stage    byte // TODO: Enum
}

func (*SetBlockDestroyStage) PacketType() string {
	return string(common2.ClientBoundPlaySetBlockDestroyStage)
}

type BlockEntityData struct {
	Location common.BlockPosition
	Type     primitive.VarInt
	NBTData  nbt.NBT
}

func (*BlockEntityData) PacketType() string {
	return string(common2.ClientBoundPlayBlockEntityData)
}

type BlockAction struct {
	Location    common.BlockPosition
	ActionID    uint8
	ActionParam uint8
	BlockType   primitive.VarInt
}

func (*BlockAction) PacketType() string {
	return string(common2.ClientBoundPlayBlockAction)
}

type BlockUpdate struct {
	Location common.BlockPosition
	BlockID  primitive.VarInt
}

func (*BlockUpdate) PacketType() string {
	return string(common2.ClientBoundPlayBlockUpdate)
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

func (*BossBar) PacketType() string {
	return string(common2.ClientBoundPlayBossBar)
}

type ChangeDifficulty struct {
	Difficulty         uint8 // TODO: Enum
	IsDifficultyLocked bool
}

func (*ChangeDifficulty) PacketType() string {
	return string(common2.ClientBoundPlayChangeDifficulty)
}

type ChunkBatchFinished struct {
	NumChunks int64
}

func (*ChunkBatchFinished) PacketType() string {
	return string(common2.ClientBoundPlayChunkBatchFinished)
}

type ChunkBatchStart struct{}

func (*ChunkBatchStart) PacketType() string {
	return string(common2.ClientBoundPlayChunkBatchStart)
}

type ClearTitles struct {
	Reset bool
}

func (*ClearTitles) PacketType() string {
	return string(common2.ClientBoundPlayClearTitles)
}

type CommandSuggestionsResponse struct {
	ID      primitive.VarInt
	Start   primitive.VarInt
	Length  primitive.VarInt
	Count   primitive.VarInt
	Matches []CommandSuggestionsMatch
}

func (*CommandSuggestionsResponse) PacketType() string {
	return string(common2.ClientBoundPlayCommandSuggestionsResponse)
}

type Commands struct {
	Count     primitive.VarInt
	Nodes     []common.CommandNode
	RootIndex primitive.VarInt
}

func (*Commands) PacketType() string {
	return string(common2.ClientBoundPlayCommands)
}

type CloseContainer struct {
	WindowID uint8
}

func (*CloseContainer) PacketType() string {
	return string(common2.ClientBoundPlayCloseContainer)
}

type SetContainerContent struct {
	WindowID    uint8
	StateID     primitive.VarInt
	Count       primitive.VarInt
	SlotData    []common.Slot
	CarriedItem common.Slot
}

func (*SetContainerContent) PacketType() string {
	return string(common2.ClientBoundPlaySetContainerContent)
}

type SetContainerProperty struct {
	WindowID uint8
	Property int16
	Value    int16
}

func (*SetContainerProperty) PacketType() string {
	return string(common2.ClientBoundPlaySetContainerProperty)
}

type SetContainerSlot struct {
	WindowID uint8
	StateID  primitive.VarInt
	Slot     int16
	SlotData common.Slot
}

func (*SetContainerSlot) PacketType() string {
	return string(common2.ClientBoundPlaySetContainerSlot)
}

type SetCooldown struct {
	ItemID primitive.VarInt
	Ticks  primitive.VarInt
}

func (*SetCooldown) PacketType() string {
	return string(common2.ClientBoundPlaySetCooldown)
}

type ChatSuggestions struct {
	Action  primitive.VarInt // TODO: Enum
	Count   primitive.VarInt
	Entries []string
}

func (*ChatSuggestions) PacketType() string {
	return string(common2.ClientBoundPlayChatSuggestions)
}

type PluginMessage struct {
	Channel primitive.Identifier
	Data    []byte
}

func (*PluginMessage) PacketType() string {
	return string(common2.ClientBoundPlayPluginMessage)
}

type DeleteMessage struct {
	SignatureLength primitive.VarInt
	Signature       []byte
}

func (*DeleteMessage) PacketType() string {
	return string(common2.ClientBoundPlayDeleteMessage)
}

type Disconnect struct {
	Reason common.Chat
}

func (*Disconnect) PacketType() string {
	return string(common2.ClientBoundPlayDisconnect)
}

type DisguisedChatMessage struct {
	Message       common.Chat
	ChatType      primitive.VarInt // TODO: Enum
	ChatTypeName  common.Chat
	HasTargetName bool
	TargetName    common.Chat
}

func (*DisguisedChatMessage) PacketType() string {
	return string(common2.ClientBoundPlayDisguisedChatMessage)
}

type EntityEvent struct {
	EntityID     int32
	EntityStatus uint8 // TODO: Enum
}

func (*EntityEvent) PacketType() string {
	return string(common2.ClientBoundPlayEntityEvent)
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

func (*Explosion) PacketType() string {
	return string(common2.ClientBoundPlayExplosion)
}

type UnloadChunk struct {
	ChunkX int32
	ChunkZ int32
}

func (*UnloadChunk) PacketType() string {
	return string(common2.ClientBoundPlayUnloadChunk)
}

type GameEvent struct {
	Event uint8 // TODO: Enum
	Value float32
}

func (*GameEvent) PacketType() string {
	return string(common2.ClientBoundPlayGameEvent)
}

type OpenHorseScreen struct {
	WindowID  uint8
	SlotCount primitive.VarInt
	EntityID  int32
}

func (*OpenHorseScreen) PacketType() string {
	return string(common2.ClientBoundPlayOpenHorseScreen)
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

func (*InitializeWorldBorder) PacketType() string {
	return string(common2.ClientBoundPlayInitializeWorldBorder)
}

type KeepAlive struct {
	ID int64
}

func (*KeepAlive) PacketType() string {
	return string(common2.ClientBoundPlayKeepAlive)
}

type ChunkDataAndUpdateLight struct {
	ChunkX              int32
	ChunkZ              int32
	Heightmaps          nbt.NBT
	Data                []common.ChunkSection
	BlockEntities       []common.BlockEntity
	TrustEdges          bool
	SkyLightMask        primitive.BitSet
	BlockLightMask      primitive.BitSet
	EmptySkyLightMask   primitive.BitSet
	EmptyBlockLightMask primitive.BitSet
	SkyLightArrays      []common.SkyLightArray
	BlockLightArrays    []common.BlockLightArray
}

func (*ChunkDataAndUpdateLight) PacketType() string {
	return string(common2.ClientBoundPlayChunkDataAndUpdateLight)
}

type WorldEvent struct {
	Event                 int32
	Location              common.BlockPosition
	Data                  int32 // TODO: Enum
	DisableRelativeVolume bool
}

func (*WorldEvent) PacketType() string {
	return string(common2.ClientBoundPlayWorldEvent)
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

func (*Particle) PacketType() string {
	return string(common2.ClientBoundPlayParticle)
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

func (*UpdateLight) PacketType() string {
	return string(common2.ClientBoundPlayUpdateLight)
}

type Login struct {
	EntityID            int32
	IsHardcore          bool
	GameMode            uint8 // TODO: Enum GameMode
	PreviousGameMode    uint8 // TODO: Enum GameMode
	DimensionCount      primitive.VarInt
	DimensionNames      []primitive.Identifier
	RegistryCodec       *nbt.NBT // Moved to configuration in 1.20.2
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
	DoLimitCrafting     bool
	DeathDimension      *primitive.Identifier
	DeathLocation       *common.BlockPosition
	PortalCooldown      primitive.VarInt
}

func (*Login) PacketType() string {
	return string(common2.ClientBoundPlayLogin)
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

func (*MapData) PacketType() string {
	return string(common2.ClientBoundPlayMapData)
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

func (*MerchantOffers) PacketType() string {
	return string(common2.ClientBoundPlayMerchantOffers)
}

type UpdateEntityPosition struct {
	EntityID primitive.VarInt
	DeltaX   int16
	DeltaY   int16
	DeltaZ   int16
	OnGround bool
}

func (*UpdateEntityPosition) PacketType() string {
	return string(common2.ClientBoundPlayUpdateEntityPosition)
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

func (*UpdateEntityPositionAndRotation) PacketType() string {
	return string(common2.ClientBoundPlayUpdateEntityPositionAndRotation)
}

type UpdateEntityRotation struct {
	EntityID primitive.VarInt
	Yaw      common.Angle
	Pitch    common.Angle
	OnGround bool
}

func (*UpdateEntityRotation) PacketType() string {
	return string(common2.ClientBoundPlayUpdateEntityRotation)
}

type MoveVehicle struct {
	X     float64
	Y     float64
	Z     float64
	Yaw   float32
	Pitch float32
}

func (*MoveVehicle) PacketType() string {
	return string(common2.ClientBoundPlayMoveVehicle)
}

type OpenBook struct {
	Hand common.Hand
}

func (*OpenBook) PacketType() string {
	return string(common2.ClientBoundPlayOpenBook)
}

type OpenScreen struct {
	WindowID    primitive.VarInt
	WindowType  primitive.VarInt // TODO: Enum
	WindowTitle common.Chat
}

func (*OpenScreen) PacketType() string {
	return string(common2.ClientBoundPlayOpenScreen)
}

type OpenSignEditor struct {
	Location common.BlockPosition
}

func (*OpenSignEditor) PacketType() string {
	return string(common2.ClientBoundPlayOpenSignEditor)
}

type Ping struct {
	ID int32
}

func (*Ping) PacketType() string {
	return string(common2.ClientBoundPlayPing)
}

type PingResponse struct {
	Payload int32
}

func (*PingResponse) PacketType() string {
	return string(common2.ClientBoundPlayPingResponse)
}

type PlaceGhostRecipe struct {
	WindowID byte
	RecipeID primitive.Identifier
}

func (*PlaceGhostRecipe) PacketType() string {
	return string(common2.ClientBoundPlayPlaceGhostRecipe)
}

type PlayerAbilities struct {
	Flags               byte // TODO: Enum
	FlyingSpeed         float32
	FieldOfViewModifier float32
}

func (*PlayerAbilities) PacketType() string {
	return string(common2.ClientBoundPlayPlayerAbilities)
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

func (*PlayerChatMessage) PacketType() string {
	return string(common2.ClientBoundPlayPlayerChatMessage)
}

// EndCombat "Unused by the Notchian client. This data was once used for twitch.tv metadata circa 1.8.f"
type EndCombat struct {
	Duration primitive.VarInt
	EntityID int32
}

func (*EndCombat) PacketType() string {
	return string(common2.ClientBoundPlayEndCombat)
}

// EnterCombat "Unused by the Notchian client. This data was once used for twitch.tv metadata circa 1.8.f"
type EnterCombat struct{}

func (*EnterCombat) PacketType() string {
	return string(common2.ClientBoundPlayEnterCombat)
}

type CombatDeath struct {
	PlayerID primitive.VarInt
	EntityID int32
	Message  common.Chat
}

func (*CombatDeath) PacketType() string {
	return string(common2.ClientBoundPlayCombatDeath)
}

type PlayerInfoRemove struct {
	NumOfPlayers primitive.VarInt
	Players      []uuid.UUID
}

func (*PlayerInfoRemove) PacketType() string {
	return string(common2.ClientBoundPlayPlayerInfoRemove)
}

type PlayerInfoUpdate struct {
	Actions PlayerInfoMask
	Players PlayerInfoUpdatePlayersData
}

func (*PlayerInfoUpdate) PacketType() string {
	return string(common2.ClientBoundPlayPlayerInfoUpdate)
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

func (*LookAt) PacketType() string {
	return string(common2.ClientBoundPlayLookAt)
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

func (*SynchronizePlayerPosition) PacketType() string {
	return string(common2.ClientBoundPlaySynchronizePlayerPosition)
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

func (*UpdateRecipeBook) PacketType() string {
	return string(common2.ClientBoundPlayUpdateRecipeBook)
}

type RemoveEntities struct {
	NumOfEntities primitive.VarInt
	Entities      []primitive.VarInt
}

func (*RemoveEntities) PacketType() string {
	return string(common2.ClientBoundPlayRemoveEntities)
}

type RemoveEntityEffect struct {
	EntityID primitive.VarInt
	EffectID primitive.VarInt // TODO: Enum
}

func (*RemoveEntityEffect) PacketType() string {
	return string(common2.ClientBoundPlayRemoveEntityEffect)
}

type ResourcePack struct {
	URL              string
	Hash             string
	Forced           bool
	HasPromptMessage bool
	PromptMessage    *common.Chat
}

func (*ResourcePack) PacketType() string {
	return string(common2.ClientBoundPlayResourcePack)
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

func (*Respawn) PacketType() string {
	return string(common2.ClientBoundPlayRespawn)
}

type SetHeadRotation struct {
	EntityID primitive.VarInt
	HeadYaw  common.Angle
}

func (*SetHeadRotation) PacketType() string {
	return string(common2.ClientBoundPlaySetHeadRotation)
}

type UpdateSectionBlocks struct {
	ChunkSectionPosition int64
	SuppressLightUpdates bool
	BlockArraySize       primitive.VarInt
	Blocks               []primitive.VarLong
}

func (*UpdateSectionBlocks) PacketType() string {
	return string(common2.ClientBoundPlayUpdateSectionBlocks)
}

type SelectAdvancementTab struct {
	HasID bool
	ID    *primitive.Identifier
}

func (*SelectAdvancementTab) PacketType() string {
	return string(common2.ClientBoundPlaySelectAdvancementTab)
}

type ServerData struct {
	HasMOTD            bool
	MOTD               *common.Chat
	HasIcon            bool
	Icon               *string
	EnforcesSecureChat bool
}

func (*ServerData) PacketType() string {
	return string(common2.ClientBoundPlayServerData)
}

type SetActionBarText struct {
	Text common.Chat
}

func (*SetActionBarText) PacketType() string {
	return string(common2.ClientBoundPlaySetActionBarText)
}

type SetBorderCenter struct {
	X float64
	Z float64
}

func (*SetBorderCenter) PacketType() string {
	return string(common2.ClientBoundPlaySetBorderCenter)
}

type SetBorderLerpSize struct {
	OldDiameter float64
	NewDiameter float64
	Speed       primitive.VarLong
}

func (*SetBorderLerpSize) PacketType() string {
	return string(common2.ClientBoundPlaySetBorderLerpSize)
}

type SetBorderSize struct {
	Diameter float64
}

func (*SetBorderSize) PacketType() string {
	return string(common2.ClientBoundPlaySetBorderSize)
}

type SetBorderWarningDelay struct {
	WarningTime primitive.VarInt
}

func (*SetBorderWarningDelay) PacketType() string {
	return string(common2.ClientBoundPlaySetBorderWarningDelay)
}

type SetBorderWarningDistance struct {
	// WarningBlocks in meters
	WarningBlocks primitive.VarInt
}

func (*SetBorderWarningDistance) PacketType() string {
	return string(common2.ClientBoundPlaySetBorderWarningDistance)
}

type SetCamera struct {
	CameraID primitive.VarInt
}

func (*SetCamera) PacketType() string {
	return string(common2.ClientBoundPlaySetCamera)
}

type SetHeldItem struct {
	Slot byte
}

func (*SetHeldItem) PacketType() string {
	return string(common2.ClientBoundPlaySetHeldItem)
}

type SetCenterChunk struct {
	ChunkX int32
	ChunkZ int32
}

func (*SetCenterChunk) PacketType() string {
	return string(common2.ClientBoundPlaySetCenterChunk)
}

type SetRenderDistance struct {
	ViewDistance primitive.VarInt
}

func (*SetRenderDistance) PacketType() string {
	return string(common2.ClientBoundPlaySetRenderDistance)
}

type SetDefaultSpawnPosition struct {
	Location common.BlockPosition
	Angle    float32
}

func (*SetDefaultSpawnPosition) PacketType() string {
	return string(common2.ClientBoundPlaySetDefaultSpawnPosition)
}

type DisplayObjective struct {
	Position common.BlockPosition
	Angle    float32
}

func (*DisplayObjective) PacketType() string {
	return string(common2.ClientBoundPlayDisplayObjective)
}

type SetEntityMetadata struct {
	EntityID primitive.VarInt
	//Metadata []common.EntityMetadata // TODO: EntityMetadata
}

func (*SetEntityMetadata) PacketType() string {
	return string(common2.ClientBoundPlaySetEntityMetadata)
}

type LinkEntities struct {
	AttachedEntityID int32
	HoldingEntityID  int32
}

func (*LinkEntities) PacketType() string {
	return string(common2.ClientBoundPlayLinkEntities)
}

type SetEntityVelocity struct {
	EntityID  primitive.VarInt
	VelocityX int16
	VelocityY int16
	VelocityZ int16
}

func (*SetEntityVelocity) PacketType() string {
	return string(common2.ClientBoundPlaySetEntityVelocity)
}

type SetEquipment struct {
	EntityID primitive.VarInt
	//Equipment common.Equipment // TODO: Equipment
}

func (*SetEquipment) PacketType() string {
	return string(common2.ClientBoundPlaySetEquipment)
}

type SetExperience struct {
	ExperienceBar   float32
	TotalExperience primitive.VarInt
	Level           primitive.VarInt
}

func (*SetExperience) PacketType() string {
	return string(common2.ClientBoundPlaySetExperience)
}

type SetHealth struct {
	Health         float32
	Food           primitive.VarInt
	FoodSaturation float32
}

func (*SetHealth) PacketType() string {
	return string(common2.ClientBoundPlaySetHealth)
}

type UpdateObjectives struct {
	ObjectiveName  string
	Mode           byte
	ObjectiveValue *common.Chat
	Type           *primitive.VarInt // TODO: Enum
}

func (*UpdateObjectives) PacketType() string {
	return string(common2.ClientBoundPlayUpdateObjectives)
}

type SetPassengers struct {
	EntityID       primitive.VarInt
	PassengerCount primitive.VarInt
	Passengers     []primitive.VarInt
}

func (*SetPassengers) PacketType() string {
	return string(common2.ClientBoundPlaySetPassengers)
}

type UpdateTeams struct {
	TeamName string
	Mode     byte
	// TODO: Complete
}

func (*UpdateTeams) PacketType() string {
	return string(common2.ClientBoundPlayUpdateTeams)
}

type UpdateScore struct {
	EntityName    string
	Action        primitive.VarInt // TODO: Enum
	ObjectiveName string
	Value         *primitive.VarInt
}

func (*UpdateScore) PacketType() string {
	return string(common2.ClientBoundPlayUpdateScore)
}

type SetSimulationDistance struct {
	SimulationDistance primitive.VarInt
}

func (*SetSimulationDistance) PacketType() string {
	return string(common2.ClientBoundPlaySetSimulationDistance)
}

type SetSubtitleText struct {
	SubtitleText common.Chat
}

func (*SetSubtitleText) PacketType() string {
	return string(common2.ClientBoundPlaySetSubtitleText)
}

type UpdateTime struct {
	WorldAge  int64
	TimeOfDay int64
}

func (*UpdateTime) PacketType() string {
	return string(common2.ClientBoundPlayUpdateTime)
}

type SetTitleText struct {
	TitleText common.Chat
}

func (*SetTitleText) PacketType() string {
	return string(common2.ClientBoundPlaySetTitleText)
}

type SetTitleAnimationTimes struct {
	FadeIn  int32
	Stay    int32
	FadeOut int32
}

func (*SetTitleAnimationTimes) PacketType() string {
	return string(common2.ClientBoundPlaySetTitleAnimationTimes)
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

func (*EntitySoundEffect) PacketType() string {
	return string(common2.ClientBoundPlayEntitySoundEffect)
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

func (*SoundEffect) PacketType() string {
	return string(common2.ClientBoundPlaySoundEffect)
}

type StopSound struct {
	Flags  byte
	Source *primitive.VarInt // TODO: Enum
	Sound  *primitive.Identifier
}

func (*StopSound) PacketType() string {
	return string(common2.ClientBoundPlayStopSound)
}

type SystemChatMessage struct {
	Content common.Chat
	Overlay bool
}

func (*SystemChatMessage) PacketType() string {
	return string(common2.ClientBoundPlaySystemChatMessage)
}

type SetTabListHeaderAndFooter struct {
	Header common.Chat
	Footer common.Chat
}

func (*SetTabListHeaderAndFooter) PacketType() string {
	return string(common2.ClientBoundPlaySetTabListHeaderAndFooter)
}

type TagQueryResponse struct {
	TransactionID primitive.VarInt
	NBT           nbt.NBT
}

func (*TagQueryResponse) PacketType() string {
	return string(common2.ClientBoundPlayTagQueryResponse)
}

type PickupItem struct {
	CollectedEntityID primitive.VarInt
	CollectorEntityID primitive.VarInt
	PickupItemCount   primitive.VarInt
}

func (*PickupItem) PacketType() string {
	return string(common2.ClientBoundPlayPickupItem)
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

func (*TeleportEntity) PacketType() string {
	return string(common2.ClientBoundPlayTeleportEntity)
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

func (*UpdateAdvancements) PacketType() string {
	return string(common2.ClientBoundPlayUpdateAdvancements)
}

type UpdateAttributes struct {
	EntityID        primitive.VarInt
	NumOfProperties primitive.VarInt
	//Properties      []AttributeProperty // TODO: AttributeProperty
}

func (*UpdateAttributes) PacketType() string {
	return string(common2.ClientBoundPlayUpdateAttributes)
}

type FeatureFlags struct {
	TotalFeatures primitive.VarInt
	FeatureFlags  []primitive.Identifier
}

func (*FeatureFlags) PacketType() string {
	return string(common2.ClientBoundPlayFeatureFlags)
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

func (*EntityEffect) PacketType() string {
	return string(common2.ClientBoundPlayEntityEffect)
}

type UpdateRecipes struct {
	RecipeCount primitive.VarInt
	//Recipes     []Recipe // TODO: Recipe
}

func (*UpdateRecipes) PacketType() string {
	return string(common2.ClientBoundPlayUpdateRecipes)
}

type UpdateTags struct {
	TagCount primitive.VarInt
	//Tags     []Tag // TODO: Tag
}

func (*UpdateTags) PacketType() string {
	return string(common2.ClientBoundPlayUpdateTags)
}
