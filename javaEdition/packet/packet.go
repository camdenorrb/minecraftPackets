package packet

import (
	"github.com/camdenorrb/minecraftPackets/javaEdition/common"
)

type CompressedPacket struct {
	*Packet
	PacketLength common.VarInt
}

type Packet struct {
	Length common.VarInt
	ID     common.VarInt
	Data   []byte
}
