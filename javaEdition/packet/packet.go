package packet

import (
	"bytes"
	"github.com/camdenorrb/minecraftPackets/primitive"
)

type Body interface {
	PacketType() string
}

type CompressedPacket struct {
	*Packet
	PacketLength primitive.VarInt
}

type Packet struct {
	Length primitive.VarInt
	ID     primitive.VarInt
	Data   bytes.Buffer
}
