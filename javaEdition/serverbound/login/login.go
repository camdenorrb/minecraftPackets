package login

import (
	"github.com/camdenorrb/minecraftPackets/primitive"
	"github.com/google/uuid"
)

type Start struct {
	Name          string
	HasPlayerUUID bool
	PlayerUUID    uuid.UUID
}

type EncryptionResponse struct {
	SharedSecret []byte
	VerifyToken  []byte
}

type PluginResponse struct {
	MessageID  primitive.VarInt
	Successful bool
	Data       []byte
}
