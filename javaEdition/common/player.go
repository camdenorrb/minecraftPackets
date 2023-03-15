package common

import "github.com/google/uuid"

type Hand VarInt

const (
	MainHand Hand = iota
	OffHand
)

// PlayerInfoAction is a player action.
// Refer to https://wiki.vg/Protocol#Player_Info_Update
type PlayerInfoAction struct {
	*PlayerInfoActionAddPlayer
	*PlayerInfoActionUpdateInitializeChat
	*PlayerInfoActionUpdateGameMode
	*PlayerInfoActionUpdateListed
	*PlayerInfoActionUpdateLatency
	*PlayerInfoActionUpdateDisplayName
}

type PlayerInfoActionAddPlayer struct {
	Name            string
	PropertiesCount VarInt
	Properties      []PlayerInfoActionAddPlayerProperty
}

type PlayerInfoActionAddPlayerProperty struct {
	Name      string
	Value     string
	IsSigned  bool
	Signature string
}

type PlayerInfoActionUpdateInitializeChat struct {
	HasSignatureData       bool
	ChatSessionID          uuid.UUID
	EncodedPublicKeySize   VarInt
	EncodedPublicKey       []byte
	PublicKeySignatureSize VarInt
	PublicKeySignature     []byte
}

type PlayerInfoActionUpdateGameMode struct {
	GameMode VarInt
}

type PlayerInfoActionUpdateListed struct {
	IsListed bool
}

type PlayerInfoActionUpdateLatency struct {
	Ping VarInt
}

type PlayerInfoActionUpdateDisplayName struct {
	HasDisplayName bool
	DisplayName    *Chat
}
