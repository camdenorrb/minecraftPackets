package login

import (
	"github.com/camdenorrb/minecraftPackets/javaEdition/common"
	common2 "github.com/camdenorrb/minecraftPackets/javaEdition/protocol/common"
	"github.com/camdenorrb/minecraftPackets/primitive"
	"github.com/google/uuid"
)

type Disconnect struct {
	Reason common.Chat
}

func (*Disconnect) PacketType() string {
	return string(common2.ClientBoundLoginDisconnect)
}

type EncryptionRequest struct {
	ServerID    string
	PublicKey   []byte
	VerifyToken []byte
}

func (*EncryptionRequest) PacketType() string {
	return string(common2.ClientBoundLoginEncryptionRequest)
}

type Success struct {
	UUID            uuid.UUID
	UserName        string
	NumOfProperties primitive.VarInt
	Properties      []common.SuccessProperty
}

func (*Success) PacketType() string {
	return string(common2.ClientBoundLoginSuccess)
}

type SetCompression struct {
	Threshold primitive.VarInt
}

func (*SetCompression) PacketType() string {
	return string(common2.ClientBoundLoginSetCompression)
}

type PluginRequest struct {
	MessageID primitive.VarInt
	Channel   string
	Data      []byte
}

func (*PluginRequest) PacketType() string {
	return string(common2.ClientBoundLoginPluginRequest)
}
