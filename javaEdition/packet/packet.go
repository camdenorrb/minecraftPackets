package packet

import "minecraftPackets/javaEdition/types"

type CompressedPacket struct {
	PacketLength types.VarInt
	*Packet
}

type Packet struct {
	Length types.VarInt
	ID     types.VarInt
	Data   []byte
}
