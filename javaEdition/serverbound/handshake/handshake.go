package handshake

import "minecraftPackets/javaEdition/types"

type Handshake struct {
	ProtocolVersion types.VarInt
	ServerAddress   string
	ServerPort      uint16
	NextState       types.VarInt
}

type LegacyServerListPing struct {
	Payload uint8
}
