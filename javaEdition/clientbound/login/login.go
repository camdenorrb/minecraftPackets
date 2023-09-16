package login

import (
	"github.com/camdenorrb/minecraftPackets/javaEdition/common"
	"github.com/camdenorrb/minecraftPackets/primitive"
	"github.com/google/uuid"
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
	NumOfProperties primitive.VarInt
	Properties      []common.SuccessProperty
}

type SetCompression struct {
	Threshold primitive.VarInt
}

type PluginRequest struct {
	MessageID primitive.VarInt
	Channel   string
	Data      []byte
}
