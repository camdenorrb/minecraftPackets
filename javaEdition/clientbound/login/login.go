package login

import (
	"github.com/google/uuid"
	"minecraftPackets/javaEdition/types"
)

type Disconnect struct {
	Reason types.Chat
}

type EncryptionRequest struct {
	ServerID    string
	PublicKey   []byte
	VerifyToken []byte
}

type Success struct {
	UUID            uuid.UUID
	UserName        string
	NumOfProperties types.VarInt
	Properties      []SuccessProperty
}

type SuccessProperty struct {
	Name      string
	Value     string
	IsSigned  bool
	Signature string
}

type SetCompression struct {
	Threshold types.VarInt
}

type PluginRequest struct {
	MessageID types.VarInt
	Channel   string
	Data      []byte
}
