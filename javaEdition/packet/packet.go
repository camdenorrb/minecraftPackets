package packet

import "github.com/camdenorrb/minecraftPackets/javaEdition/common"

type CompressedPacket struct {
	PacketLength common.VarInt
	*Packet
}

type Packet struct {
	Length common.VarInt
	ID     common.VarInt
	Data   []byte
}
