package common

import (
	"github.com/camdenorrb/minecraftPackets/primitive"
	"github.com/google/uuid"
)

type Hand primitive.VarInt

const (
	MainHand Hand = iota
	OffHand
)

// PlayerInfoAction is a player action.
// Refer to https://wiki.vg/Protocol#Player_Info_Update
type PlayerInfoAction struct {
	AddPlayer         *PlayerInfoActionAddPlayer
	InitializeChat    *PlayerInfoActionUpdateInitializeChat
	UpdateGameMode    *PlayerInfoActionUpdateGameMode
	UpdateListed      *PlayerInfoActionUpdateListed
	UpdateLatency     *PlayerInfoActionUpdateLatency
	UpdateDisplayName *PlayerInfoActionUpdateDisplayName
}

type PlayerInfoActionAddPlayer struct {
	Name       string
	Properties []PlayerInfoActionAddPlayerProperty
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
	EncodedPublicKeySize   primitive.VarInt
	EncodedPublicKey       []byte
	PublicKeySignatureSize primitive.VarInt
	PublicKeySignature     []byte
}

type PlayerInfoActionUpdateGameMode struct {
	GameMode primitive.VarInt
}

type PlayerInfoActionUpdateListed struct {
	IsListed bool
}

type PlayerInfoActionUpdateLatency struct {
	Ping primitive.VarInt
}

type PlayerInfoActionUpdateDisplayName struct {
	DisplayName *Chat
}
