package play

import (
	"github.com/camdenorrb/minecraftPackets/javaEdition/common"
	"github.com/google/uuid"
)

type ConfirmTeleportation struct {
	TeleportID common.VarInt
}

type QueryBlockEntityTag struct {
	TransactionID common.VarInt
	Location      common.BlockPosition
}

type ChangeDifficulty struct {
	NewDifficulty byte // TODO: Enum
}

type MessageAcknowledgement struct {
	MessageCount common.VarInt
}

type ChatCommand struct {
	Command        string
	Timestamp      int64
	Salt           int64
	SignatureCount common.VarInt
	Signatures     []common.ChatCommandSignature
	MessageCount   common.VarInt
	Acknowledged   common.BitSet
}

type ChatMessage struct {
	Message      string
	Timestamp    int64
	Salt         int64
	HasSignature bool
	Signature    []byte
	MessageCount common.VarInt
	Acknowledged common.BitSet
}

type ClientCommand struct {
	ActionID common.VarInt // TODO: Enum
}

type ClientInformation struct {
	Locale              string
	ViewDistance        byte
	ChatMode            common.VarInt // TODO: Enum
	ChatColors          bool
	DisplayedSkinParts  uint8         // TODO: Enum
	MainHand            common.VarInt // TODO: Enum
	EnableTextFiltering bool
	AllowServerListings bool
}

type CommandSuggestionsRequest struct {
	TransactionID common.VarInt
	Text          string
}

type ClickContainerButton struct {
	WindowID byte
	ButtonID byte
}

type ClickContainer struct {
	WindowID    uint8
	StateID     common.VarInt
	Slot        int16
	Button      byte
	Mode        common.VarInt // TODO: Enum
	SlotCount   common.VarInt
	Slots       []common.Slot
	CarriedItem common.Slot
}

type CloseContainer struct {
	WindowID uint8
}

type PluginMessage struct {
	Channel common.Identifier
	Data    []byte
}

type EditBook struct {
	Slot     common.VarInt
	Count    common.VarInt
	Entries  []string
	HasTitle bool
	Title    *string
}

type QueryEntityTag struct {
	TransactionID common.VarInt
	EntityID      common.VarInt
}

type InteractEntity struct {
	EntityID common.VarInt
	Type     common.VarInt // TODO: Enum
	TargetX  float32
	TargetY  float32
	TargetZ  float32
	Hand     common.VarInt // TODO: Enum
	Sneaking bool
}

type JigsawGenerate struct {
	Location    common.BlockPosition
	Levels      common.VarInt
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
	Slot common.VarInt
}

type PlayerAbilities struct {
	Flags uint8 // TODO: Enum
}

type PlayerAction struct {
	Status   common.VarInt // TODO: Enum
	Location common.BlockPosition
	Face     byte // TODO: Enum
	Sequence common.VarInt
}

type PlayerCommand struct {
	EntityID  common.VarInt
	ActionID  common.VarInt // TODO: Enum
	JumpBoost common.VarInt
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

type ChangeRecipeBookSettings struct {
	BookID       common.VarInt // TODO: Enum
	BookOpen     bool
	FilterActive bool
}

type SetSeenRecipe struct {
	RecipeID common.Identifier
}

type RenameItem struct {
	ItemName string
}

type ResourcePack struct {
	Result common.VarInt // TODO: Enum
}

type SeenAdvancements struct {
	Action common.VarInt // TODO: Enum
	TabID  *common.Identifier
}

type SelectTrade struct {
	SelectedSlot common.VarInt
}

type SetBeaconEffect struct {
	HasPrimaryEffect   bool
	PrimaryEffect      common.VarInt // TODO: Enum
	HasSecondaryEffect bool
	SecondaryEffect    common.VarInt // TODO: Enum
}

type SetHeldItem struct {
	Slot int16
}

type ProgramCommandBlock struct {
	Location common.BlockPosition
	Command  string
	Mode     common.VarInt // TODO: Enum
	Flags    byte          // TODO: Enum
}

type ProgramCommandBlockMinecart struct {
	EntityID    common.VarInt
	Command     string
	TrackOutput bool
}

type SetCreativeModeSlot struct {
	Slot        int16
	ClickedItem common.Slot
}

type ProgramJigsawBlock struct {
	Location   common.BlockPosition
	Name       common.Identifier
	Target     common.Identifier
	Pool       common.Identifier
	FinalState string
	JointType  string
}

type ProgramStructureBlock struct {
	Location  common.BlockPosition
	Action    common.VarInt // TODO: Enum
	Mode      common.VarInt // TODO: Enum
	Name      string
	OffsetX   byte
	OffsetY   byte
	OffsetZ   byte
	SizeX     byte
	SizeY     byte
	SizeZ     byte
	Mirror    common.VarInt // TODO: Enum
	Rotation  common.VarInt // TODO: Enum
	Metadata  string
	Integrity float32
	Seed      common.VarLong
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
	Hand common.VarInt // TODO: Enum
}

type TeleportToEntity struct {
	TargetPlayer uuid.UUID
}

type UseItemOn struct {
	Hand        common.VarInt // TODO: Enum
	Location    common.BlockPosition
	Face        common.VarInt // TODO: Enum
	CursorX     float32
	CursorY     float32
	CursorZ     float32
	InsideBlock bool
	Sequence    common.VarInt
}

type UseItem struct {
	Hand     common.VarInt // TODO: Enum
	Sequence common.VarInt
}
