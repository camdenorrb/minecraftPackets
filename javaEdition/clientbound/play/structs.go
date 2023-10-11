package play

import (
	"github.com/camdenorrb/minecraftPackets/javaEdition/common"
	"github.com/camdenorrb/minecraftPackets/primitive"
	"github.com/google/uuid"
)

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

type CommandSuggestionsMatch struct {
	Match      string
	HasToolTip bool
	Tooltip    *common.Chat
}

type Icon struct {
	Type           primitive.VarInt // TODO: Enum
	X              byte
	Z              byte
	Direction      byte
	HasDisplayName bool
	DisplayName    *common.Chat
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

type PreviousMessage struct {
	MessageID primitive.VarInt
	Signature []byte
}

type PlayerInfoUpdatePlayersData struct {
	UUID          uuid.UUID
	PlayerActions []common.PlayerInfoAction
}
