package login

import (
	"github.com/google/uuid"
	"minecraftPackets/javaEdition/types"
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
	MessageID  types.VarInt
	Successful bool
	Data       []byte
}
