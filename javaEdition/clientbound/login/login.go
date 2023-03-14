package login

import (
	"github.com/google/uuid"
	"minecraftPackets/javaEdition/common"
)

type Disconnect struct {
	Reason common.Chat
}

type EncryptionRequest struct {
	ServerID    string
	PublicKey   []byte
	VerifyToken []byte
}

type Success struct {
	UUID            uuid.UUID
	UserName        string
	NumOfProperties common.VarInt
	Properties      []SuccessProperty
}

type SuccessProperty struct {
	Name      string
	Value     string
	IsSigned  bool
	Signature string
}

type SetCompression struct {
	Threshold common.VarInt
}

type PluginRequest struct {
	MessageID common.VarInt
	Channel   string
	Data      []byte
}
