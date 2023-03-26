package login

import (
	"github.com/google/uuid"
	"javaEdition/common"
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
	MessageID  common.VarInt
	Successful bool
	Data       []byte
}
