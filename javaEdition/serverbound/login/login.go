package login

import (
	common2 "github.com/camdenorrb/minecraftPackets/javaEdition/protocol/common"
	"github.com/camdenorrb/minecraftPackets/primitive"
	"github.com/google/uuid"
)

type Start struct {
	Name          string
	HasPlayerUUID bool
	PlayerUUID    uuid.UUID
}

func (*Start) PacketType() string {
	return string(common2.ServerBoundLoginStart)
}

type EncryptionResponse struct {
	SharedSecret []byte
	VerifyToken  []byte
}

func (*EncryptionResponse) PacketType() string {
	return string(common2.ServerBoundLoginEncryptionResponse)
}

type PluginResponse struct {
	MessageID  primitive.VarInt
	Successful bool
	Data       []byte
}

func (*PluginResponse) PacketType() string {
	return string(common2.ServerBoundLoginPluginResponse)
}

type Acknowledged struct{}

func (*Acknowledged) PacketType() string {
	return string(common2.ServerBoundLoginAcknowledged)
}
