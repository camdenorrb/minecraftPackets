package packet

import (
	"bytes"
	"javaEdition/common"
)

type CompressedPacket struct {
	*Packet
	PacketLength common.VarInt
}

type Packet struct {
	Length common.VarInt
	ID     common.VarInt
	Data   bytes.Buffer
}
