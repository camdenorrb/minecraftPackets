package packet

import (
	"bytes"
	"github.com/camdenorrb/minecraftPackets/javaEdition/v1/common"
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
