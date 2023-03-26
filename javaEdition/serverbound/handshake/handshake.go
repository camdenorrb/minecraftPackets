package handshake

import "github.com/camdenorrb/minecraftPackets/javaEdition/common"

type Handshake struct {
	ProtocolVersion common.VarInt
	ServerAddress   string
	ServerPort      uint16
	NextState       common.VarInt
}

type LegacyServerListPing struct {
	Payload uint8
}
