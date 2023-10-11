package handshake

import (
	common2 "github.com/camdenorrb/minecraftPackets/javaEdition/protocol/common"
	"github.com/camdenorrb/minecraftPackets/primitive"
)

type Handshake struct {
	ProtocolVersion primitive.VarInt
	ServerAddress   string
	ServerPort      uint16
	NextState       primitive.VarInt
}

func (*Handshake) PacketType() string {
	return string(common2.ServerBoundHandshake)
}

type LegacyServerListPing struct {
	Payload uint8
}

func (*LegacyServerListPing) PacketType() string {
	return string(common2.ServerBoundHandshakeLegacyServerListPing)
}
