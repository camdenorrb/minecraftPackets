package play

import (
	"github.com/google/uuid"
	"minecraftPackets/javaEdition/types"
)

type ConfirmTeleportation struct {
	TeleportID types.VarInt
}

type QueryBlockEntityTag struct {
	TransactionID types.VarInt
	Location      types.BlockPosition
}

type ChangeDifficulty struct {
	NewDifficulty byte // TODO: Enum
}

type MessageAcknowledgement struct {
	MessageCount types.VarInt
}

type ChatCommand struct {
	Command        string
	Timestamp      int64
	Salt           int64
	SignatureCount types.VarInt
	//Signatures     []ChatCommandSignature // TODO: Signature
	MessageCount types.VarInt
	Acknowledged types.BitSet
}

type ChatMessage struct {
	Message      string
	Timestamp    int64
	Salt         int64
	HasSignature bool
	Signature    []byte
	MessageCount types.VarInt
	Acknowledged types.BitSet
}

type ClientCommand struct {
	ActionID types.VarInt // TODO: Enum
}

type ClientInformation struct {
	Locale              string
	ViewDistance        byte
	ChatMode            types.VarInt // TODO: Enum
	ChatColors          bool
	DisplayedSkinParts  uint8        // TODO: Enum
	MainHand            types.VarInt // TODO: Enum
	EnableTextFiltering bool
	AllowServerListings bool
}

type CommandSuggestionsRequest struct {
	TransactionID types.VarInt
	Text          string
}

type ClickContainerButton struct {
	WindowID byte
	ButtonID byte
}

type ClickContainer struct {
	WindowID    uint8
	StateID     types.VarInt
	Slot        int16
	Button      byte
	Mode        types.VarInt // TODO: Enum
	SlotCount   types.VarInt
	Slots       []types.Slot
	CarriedItem types.Slot
}

type CloseContainer struct {
	WindowID uint8
}

type PluginMessage struct {
	Channel types.Identifier
	Data    []byte
}

type EditBook struct {
	Slot     types.VarInt
	Count    types.VarInt
	Entries  []string
	HasTitle bool
	Title    *string
}

type QueryEntityTag struct {
	TransactionID types.VarInt
	EntityID      types.VarInt
}

type InteractEntity struct {
	EntityID types.VarInt
	Type     types.VarInt // TODO: Enum
	TargetX  float32
	TargetY  float32
	TargetZ  float32
	Hand     types.VarInt // TODO: Enum
	Sneaking bool
}

type JigsawGenerate struct {
	Location    types.BlockPosition
	Levels      types.VarInt
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
	Slot types.VarInt
}

type PlayerAbilities struct {
	Flags uint8 // TODO: Enum
}

type PlayerAction struct {
	Status   types.VarInt // TODO: Enum
	Location types.BlockPosition
	Face     byte // TODO: Enum
	Sequence types.VarInt
}

type PlayerCommand struct {
	EntityID  types.VarInt
	ActionID  types.VarInt // TODO: Enum
	JumpBoost types.VarInt
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
	//PublicKey types.PublicKey // TODO: PublicKey
}

type ChangeRecipeBookSettings struct {
	BookID       types.VarInt // TODO: Enum
	BookOpen     bool
	FilterActive bool
}

type SetSeenRecipe struct {
	RecipeID types.Identifier
}

type RenameItem struct {
	ItemName string
}

type ResourcePack struct {
	Result types.VarInt // TODO: Enum
}

type SeenAdvancements struct {
	Action types.VarInt // TODO: Enum
	TabID  *types.Identifier
}

type SelectTrade struct {
	SelectedSlot types.VarInt
}

type SetBeaconEffect struct {
	HasPrimaryEffect   bool
	PrimaryEffect      types.VarInt // TODO: Enum
	HasSecondaryEffect bool
	SecondaryEffect    types.VarInt // TODO: Enum
}

type SetHeldItem struct {
	Slot int16
}

type ProgramCommandBlock struct {
	Location types.BlockPosition
	Command  string
	Mode     types.VarInt // TODO: Enum
	Flags    byte         // TODO: Enum
}

type ProgramCommandBlockMinecart struct {
	EntityID    types.VarInt
	Command     string
	TrackOutput bool
}

type SetCreativeModeSlot struct {
	Slot        int16
	ClickedItem types.Slot
}

type ProgramJigsawBlock struct {
	Location   types.BlockPosition
	Name       types.Identifier
	Target     types.Identifier
	Pool       types.Identifier
	FinalState string
	JointType  string
}

type ProgramStructureBlock struct {
	Location  types.BlockPosition
	Action    types.VarInt // TODO: Enum
	Mode      types.VarInt // TODO: Enum
	Name      string
	OffsetX   byte
	OffsetY   byte
	OffsetZ   byte
	SizeX     byte
	SizeY     byte
	SizeZ     byte
	Mirror    types.VarInt // TODO: Enum
	Rotation  types.VarInt // TODO: Enum
	Metadata  string
	Integrity float32
	Seed      types.VarLong
	Flags     byte // TODO: Enum
}

type UpdateSign struct {
	Location types.BlockPosition
	Line1    string
	Line2    string
	Line3    string
	Line4    string
}

type SwingArm struct {
	Hand types.VarInt // TODO: Enum
}

type TeleportToEntity struct {
	TargetPlayer uuid.UUID
}

type UseItemOn struct {
	Hand        types.VarInt // TODO: Enum
	Location    types.BlockPosition
	Face        types.VarInt // TODO: Enum
	CursorX     float32
	CursorY     float32
	CursorZ     float32
	InsideBlock bool
	Sequence    types.VarInt
}

type UseItem struct {
	Hand     types.VarInt // TODO: Enum
	Sequence types.VarInt
}
