package common

import "github.com/camdenorrb/minecraftPackets/nbt"

type BlockEntity struct {
	PackedXZ byte
	Y        int16
	Type     VarInt
	Data     nbt.NBT
}
