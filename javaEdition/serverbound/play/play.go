package play

import (
	"github.com/camdenorrb/minecraftPackets/javaEdition/common"
	common2 "github.com/camdenorrb/minecraftPackets/javaEdition/protocol/common"
	"github.com/camdenorrb/minecraftPackets/primitive"
	"github.com/google/uuid"
)

type ConfirmTeleportation struct {
	TeleportID primitive.VarInt
}

func (*ConfirmTeleportation) PacketType() string {
	return string(common2.ServerBoundPlayConfirmTeleportation)
}

type QueryBlockEntityTag struct {
	TransactionID primitive.VarInt
	Location      common.BlockPosition
}

func (*QueryBlockEntityTag) PacketType() string {
	return string(common2.ServerBoundPlayQueryBlockEntityTag)
}

type ChangeDifficulty struct {
	NewDifficulty byte // TODO: Enum
}

func (*ChangeDifficulty) PacketType() string {
	return string(common2.ServerBoundPlayChangeDifficulty)
}

type MessageAcknowledgement struct {
	MessageCount primitive.VarInt
}

func (*MessageAcknowledgement) PacketType() string {
	return string(common2.ServerBoundPlayMessageAcknowledgement)
}

type ChatCommand struct {
	Command        string
	Timestamp      int64
	Salt           int64
	SignatureCount primitive.VarInt
	Signatures     []common.ChatCommandSignature
	MessageCount   primitive.VarInt
	Acknowledged   primitive.BitSet
}

func (*ChatCommand) PacketType() string {
	return string(common2.ServerBoundPlayChatCommand)
}

type ChatMessage struct {
	Message      string
	Timestamp    int64
	Salt         int64
	HasSignature bool
	Signature    []byte
	MessageCount primitive.VarInt
	Acknowledged primitive.BitSet
}

func (*ChatMessage) PacketType() string {
	return string(common2.ServerBoundPlayChatMessage)
}

type ClientCommand struct {
	ActionID primitive.VarInt // TODO: Enum
}

func (*ClientCommand) PacketType() string {
	return string(common2.ServerBoundPlayClientCommand)
}

type ClientInformation struct {
	Locale              string
	ViewDistance        byte
	ChatMode            primitive.VarInt // TODO: Enum
	ChatColors          bool
	DisplayedSkinParts  uint8 // TODO: Enum
	MainHand            common.Hand
	EnableTextFiltering bool
	AllowServerListings bool
}

func (*ClientInformation) PacketType() string {
	return string(common2.ServerBoundPlayClientInformation)
}

type CommandSuggestionsRequest struct {
	TransactionID primitive.VarInt
	Text          string
}

func (*CommandSuggestionsRequest) PacketType() string {
	return string(common2.ServerBoundPlayCommandSuggestionsRequest)
}

type ConfigurationAcknowledged struct{}

func (*ConfigurationAcknowledged) PacketType() string {
	return string(common2.ServerBoundPlayConfigurationAcknowledged)
}

type ClickContainerButton struct {
	WindowID byte
	ButtonID byte
}

func (*ClickContainerButton) PacketType() string {
	return string(common2.ServerBoundPlayClickContainerButton)
}

type ClickContainer struct {
	WindowID    uint8
	StateID     primitive.VarInt
	Slot        int16
	Button      byte
	Mode        primitive.VarInt // TODO: Enum
	SlotCount   primitive.VarInt
	Slots       []common.Slot
	CarriedItem common.Slot
}

func (*ClickContainer) PacketType() string {
	return string(common2.ServerBoundPlayClickContainer)
}

type CloseContainer struct {
	WindowID uint8
}

func (*CloseContainer) PacketType() string {
	return string(common2.ServerBoundPlayCloseContainer)
}

type PluginMessage struct {
	Channel primitive.Identifier
	Data    []byte
}

func (*PluginMessage) PacketType() string {
	return string(common2.ServerBoundPlayPluginMessage)
}

type EditBook struct {
	Slot     primitive.VarInt
	Count    primitive.VarInt
	Entries  []string
	HasTitle bool
	Title    *string
}

func (*EditBook) PacketType() string {
	return string(common2.ServerBoundPlayEditBook)
}

type QueryEntityTag struct {
	TransactionID primitive.VarInt
	EntityID      primitive.VarInt
}

func (*QueryEntityTag) PacketType() string {
	return string(common2.ServerBoundPlayQueryEntityTag)
}

type InteractEntity struct {
	EntityID primitive.VarInt
	Type     primitive.VarInt // TODO: Enum
	TargetX  float32
	TargetY  float32
	TargetZ  float32
	Hand     common.Hand
	Sneaking bool
}

func (*InteractEntity) PacketType() string {
	return string(common2.ServerBoundPlayInteractEntity)
}

type JigsawGenerate struct {
	Location    common.BlockPosition
	Levels      primitive.VarInt
	KeepJigsaws bool
}

func (*JigsawGenerate) PacketType() string {
	return string(common2.ServerBoundPlayJigsawGenerate)
}

type KeepAlive struct {
	KeepAliveID int64
}

func (*KeepAlive) PacketType() string {
	return string(common2.ServerBoundPlayKeepAlive)
}

type LockDifficulty struct {
	Locked bool
}

func (*LockDifficulty) PacketType() string {
	return string(common2.ServerBoundPlayLockDifficulty)
}

type SetPlayerPosition struct {
	X        float64
	Y        float64
	Z        float64
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (*SetPlayerPosition) PacketType() string {
	return string(common2.ServerBoundPlaySetPlayerPosition)
}

type SetPlayerPositionAndRotation struct {
	X        float64
	FeetY    float64
	Z        float64
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (*SetPlayerPositionAndRotation) PacketType() string {
	return string(common2.ServerBoundPlaySetPlayerPositionAndRotation)
}

type SetPlayerRotation struct {
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (*SetPlayerRotation) PacketType() string {
	return string(common2.ServerBoundPlaySetPlayerRotation)
}

type SetPlayerOnGround struct {
	OnGround bool
}

func (*SetPlayerOnGround) PacketType() string {
	return string(common2.ServerBoundPlaySetPlayerOnGround)
}

type MoveVehicle struct {
	X     float64
	Y     float64
	Z     float64
	Yaw   float32
	Pitch float32
}

func (*MoveVehicle) PacketType() string {
	return string(common2.ServerBoundPlayMoveVehicle)
}

type PaddleBoat struct {
	LeftPaddleTurning  bool
	RightPaddleTurning bool
}

func (*PaddleBoat) PacketType() string {
	return string(common2.ServerBoundPlayPaddleBoat)
}

type PickItem struct {
	Slot primitive.VarInt
}

func (*PickItem) PacketType() string {
	return string(common2.ServerBoundPlayPickItem)
}

type PlaceRecipe struct {
	WindowID uint8
	Recipe   primitive.Identifier
	MakeAll  bool
}

func (*PlaceRecipe) PacketType() string {
	return string(common2.ServerBoundPlayPlaceRecipe)
}

type PlayerAbilities struct {
	Flags uint8 // TODO: Enum
}

func (*PlayerAbilities) PacketType() string {
	return string(common2.ServerBoundPlayPlayerAbilities)
}

type PlayerAction struct {
	Status   primitive.VarInt // TODO: Enum
	Location common.BlockPosition
	Face     byte // TODO: Enum
	Sequence primitive.VarInt
}

func (*PlayerAction) PacketType() string {
	return string(common2.ServerBoundPlayPlayerAction)
}

type PlayerCommand struct {
	EntityID  primitive.VarInt
	ActionID  primitive.VarInt // TODO: Enum
	JumpBoost primitive.VarInt
}

func (*PlayerCommand) PacketType() string {
	return string(common2.ServerBoundPlayPlayerCommand)
}

type PlayerInput struct {
	Sideways float32
	Forward  float32
	Flags    uint8 // TODO: Enum
}

func (*PlayerInput) PacketType() string {
	return string(common2.ServerBoundPlayPlayerInput)
}

type Pong struct {
	ID int32
}

func (*Pong) PacketType() string {
	return string(common2.ServerBoundPlayPong)
}

type PlayerSession struct {
	SessionID uuid.UUID
	PublicKey common.PublicKey
}

func (*PlayerSession) PacketType() string {
	return string(common2.ServerBoundPlayPlayerSession)
}

type ChunkBatchReceived struct {
	ChunksPerTick float32
}

func (*ChunkBatchReceived) PacketType() string {
	return string(common2.ServerBoundPlayChunkBatchReceived)
}

type ChangeRecipeBookSettings struct {
	BookID       primitive.VarInt // TODO: Enum
	BookOpen     bool
	FilterActive bool
}

func (*ChangeRecipeBookSettings) PacketType() string {
	return string(common2.ServerBoundPlayChangeRecipeBookSettings)
}

type SetSeenRecipe struct {
	RecipeID primitive.Identifier
}

func (*SetSeenRecipe) PacketType() string {
	return string(common2.ServerBoundPlaySetSeenRecipe)
}

type RenameItem struct {
	ItemName string
}

func (*RenameItem) PacketType() string {
	return string(common2.ServerBoundPlayRenameItem)
}

type ResourcePack struct {
	Result primitive.VarInt // TODO: Enum
}

func (*ResourcePack) PacketType() string {
	return string(common2.ServerBoundPlayResourcePack)
}

type SeenAdvancements struct {
	Action primitive.VarInt // TODO: Enum
	TabID  *primitive.Identifier
}

func (*SeenAdvancements) PacketType() string {
	return string(common2.ServerBoundPlaySeenAdvancements)
}

type SelectTrade struct {
	SelectedSlot primitive.VarInt
}

func (*SelectTrade) PacketType() string {
	return string(common2.ServerBoundPlaySelectTrade)
}

type SetBeaconEffect struct {
	HasPrimaryEffect   bool
	PrimaryEffect      primitive.VarInt // TODO: Enum
	HasSecondaryEffect bool
	SecondaryEffect    primitive.VarInt // TODO: Enum
}

func (*SetBeaconEffect) PacketType() string {
	return string(common2.ServerBoundPlaySetBeaconEffect)
}

type SetHeldItem struct {
	Slot int16
}

func (*SetHeldItem) PacketType() string {
	return string(common2.ServerBoundPlaySetHeldItem)
}

type ProgramCommandBlock struct {
	Location common.BlockPosition
	Command  string
	Mode     primitive.VarInt // TODO: Enum
	Flags    byte             // TODO: Enum
}

func (*ProgramCommandBlock) PacketType() string {
	return string(common2.ServerBoundPlayProgramCommandBlock)
}

type ProgramCommandBlockMinecart struct {
	EntityID    primitive.VarInt
	Command     string
	TrackOutput bool
}

func (*ProgramCommandBlockMinecart) PacketType() string {
	return string(common2.ServerBoundPlayProgramCommandBlockMinecart)
}

type SetCreativeModeSlot struct {
	Slot        int16
	ClickedItem common.Slot
}

func (*SetCreativeModeSlot) PacketType() string {
	return string(common2.ServerBoundPlaySetCreativeModeSlot)
}

type ProgramJigsawBlock struct {
	Location   common.BlockPosition
	Name       primitive.Identifier
	Target     primitive.Identifier
	Pool       primitive.Identifier
	FinalState string
	JointType  string
}

func (*ProgramJigsawBlock) PacketType() string {
	return string(common2.ServerBoundPlayProgramJigsawBlock)
}

type ProgramStructureBlock struct {
	Location  common.BlockPosition
	Action    primitive.VarInt // TODO: Enum
	Mode      primitive.VarInt // TODO: Enum
	Name      string
	OffsetX   byte
	OffsetY   byte
	OffsetZ   byte
	SizeX     byte
	SizeY     byte
	SizeZ     byte
	Mirror    primitive.VarInt // TODO: Enum
	Rotation  primitive.VarInt // TODO: Enum
	Metadata  string
	Integrity float32
	Seed      primitive.VarLong
	Flags     byte // TODO: Enum
}

func (*ProgramStructureBlock) PacketType() string {
	return string(common2.ServerBoundPlayProgramStructureBlock)
}

type UpdateSign struct {
	Location common.BlockPosition
	Line1    string
	Line2    string
	Line3    string
	Line4    string
}

func (*UpdateSign) PacketType() string {
	return string(common2.ServerBoundPlayUpdateSign)
}

type SwingArm struct {
	Hand common.Hand
}

func (*SwingArm) PacketType() string {
	return string(common2.ServerBoundPlaySwingArm)
}

type TeleportToEntity struct {
	TargetPlayer uuid.UUID
}

func (*TeleportToEntity) PacketType() string {
	return string(common2.ServerBoundPlayTeleportToEntity)
}

type UseItemOn struct {
	Hand        common.Hand
	Location    common.BlockPosition
	Face        primitive.VarInt // TODO: Enum
	CursorX     float32
	CursorY     float32
	CursorZ     float32
	InsideBlock bool
	Sequence    primitive.VarInt
}

func (*UseItemOn) PacketType() string {
	return string(common2.ServerBoundPlayUseItemOn)
}

type UseItem struct {
	Hand     common.Hand
	Sequence primitive.VarInt
}

func (*UseItem) PacketType() string {
	return string(common2.ServerBoundPlayUseItem)
}
