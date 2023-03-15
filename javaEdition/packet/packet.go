package packet

import (
	"github.com/camdenorrb/minecraftPackets/javaEdition/common"
	"github.com/camdenorrb/minecraftPackets/javaEdition/protocol"
)

type CompressedPacket struct {
	PacketLength common.VarInt
	*Packet
}

type Packet struct {
	Length common.VarInt
	ID     common.VarInt
	Data   []byte
}

type Payload interface {
	ID(version protocol.Version) common.VarInt
}
