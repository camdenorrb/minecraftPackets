package play

import (
	"github.com/camdenorrb/minecraftPackets/javaEdition/common"
	"github.com/camdenorrb/minecraftPackets/primitive"
	"github.com/google/uuid"
)

type ConfirmTeleportation struct {
	TeleportID primitive.VarInt
}

type QueryBlockEntityTag struct {
	TransactionID primitive.VarInt
	Location      common.BlockPosition
}

type ChangeDifficulty struct {
	NewDifficulty byte // TODO: Enum
}

type MessageAcknowledgement struct {
	MessageCount primitive.VarInt
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

type ChatMessage struct {
	Message      string
	Timestamp    int64
	Salt         int64
	HasSignature bool
	Signature    []byte
	MessageCount primitive.VarInt
	Acknowledged primitive.BitSet
}

type ClientCommand struct {
	ActionID primitive.VarInt // TODO: Enum
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

type CommandSuggestionsRequest struct {
	TransactionID primitive.VarInt
	Text          string
}

type ConfigurationAcknowledged struct{}

type ClickContainerButton struct {
	WindowID byte
	ButtonID byte
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

type CloseContainer struct {
	WindowID uint8
}

type PluginMessage struct {
	Channel primitive.Identifier
	Data    []byte
}

type EditBook struct {
	Slot     primitive.VarInt
	Count    primitive.VarInt
	Entries  []string
	HasTitle bool
	Title    *string
}

type QueryEntityTag struct {
	TransactionID primitive.VarInt
	EntityID      primitive.VarInt
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

type JigsawGenerate struct {
	Location    common.BlockPosition
	Levels      primitive.VarInt
	KeepJigsaws bool
}

type KeepAlive struct {
	KeepAliveID int64
}

type LockDifficulty struct {
	Locked bool
}

type SetPlayerPosition struct {
	X        float64
	Y        float64
	Z        float64
	Yaw      float32
	Pitch    float32
	OnGround bool
}

type SetPlayerPositionAndRotation struct {
	X        float64
	FeetY    float64
	Z        float64
	Yaw      float32
	Pitch    float32
	OnGround bool
}

type SetPlayerRotation struct {
	Yaw      float32
	Pitch    float32
	OnGround bool
}

type SetPlayerOnGround struct {
	OnGround bool
}

type MoveVehicle struct {
	X     float64
	Y     float64
	Z     float64
	Yaw   float32
	Pitch float32
}

type PaddleBoat struct {
	LeftPaddleTurning  bool
	RightPaddleTurning bool
}

type PickItem struct {
	Slot primitive.VarInt
}

type PlaceRecipe struct {
	WindowID uint8
	Recipe   primitive.Identifier
	MakeAll  bool
}

type PlayerAbilities struct {
	Flags uint8 // TODO: Enum
}

type PlayerAction struct {
	Status   primitive.VarInt // TODO: Enum
	Location common.BlockPosition
	Face     byte // TODO: Enum
	Sequence primitive.VarInt
}

type PlayerCommand struct {
	EntityID  primitive.VarInt
	ActionID  primitive.VarInt // TODO: Enum
	JumpBoost primitive.VarInt
}

type PlayerInput struct {
	Sideways float32
	Forward  float32
	Flags    uint8 // TODO: Enum
}

type Pong struct {
	ID int32
}

type PlayerSession struct {
	SessionID uuid.UUID
	PublicKey common.PublicKey
}

type ChunkBatchReceived struct {
	ChunksPerTick float32
}

type ChangeRecipeBookSettings struct {
	BookID       primitive.VarInt // TODO: Enum
	BookOpen     bool
	FilterActive bool
}

type SetSeenRecipe struct {
	RecipeID primitive.Identifier
}

type RenameItem struct {
	ItemName string
}

type ResourcePack struct {
	Result primitive.VarInt // TODO: Enum
}

type SeenAdvancements struct {
	Action primitive.VarInt // TODO: Enum
	TabID  *primitive.Identifier
}

type SelectTrade struct {
	SelectedSlot primitive.VarInt
}

type SetBeaconEffect struct {
	HasPrimaryEffect   bool
	PrimaryEffect      primitive.VarInt // TODO: Enum
	HasSecondaryEffect bool
	SecondaryEffect    primitive.VarInt // TODO: Enum
}

type SetHeldItem struct {
	Slot int16
}

type ProgramCommandBlock struct {
	Location common.BlockPosition
	Command  string
	Mode     primitive.VarInt // TODO: Enum
	Flags    byte             // TODO: Enum
}

type ProgramCommandBlockMinecart struct {
	EntityID    primitive.VarInt
	Command     string
	TrackOutput bool
}

type SetCreativeModeSlot struct {
	Slot        int16
	ClickedItem common.Slot
}

type ProgramJigsawBlock struct {
	Location   common.BlockPosition
	Name       primitive.Identifier
	Target     primitive.Identifier
	Pool       primitive.Identifier
	FinalState string
	JointType  string
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

type UpdateSign struct {
	Location common.BlockPosition
	Line1    string
	Line2    string
	Line3    string
	Line4    string
}

type SwingArm struct {
	Hand common.Hand
}

type TeleportToEntity struct {
	TargetPlayer uuid.UUID
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

type UseItem struct {
	Hand     common.Hand
	Sequence primitive.VarInt
}
