package handshake

import (
	"github.com/camdenorrb/minecraftPackets/primitive"
)

type Handshake struct {
	ProtocolVersion primitive.VarInt
	ServerAddress   string
	ServerPort      uint16
	NextState       primitive.VarInt
}

type LegacyServerListPing struct {
	Payload uint8
}
