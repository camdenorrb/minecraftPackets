package common

import (
	"github.com/camdenorrb/minecraftPackets/nbt"
	"github.com/camdenorrb/minecraftPackets/primitive"
)

type BlockEntity struct {
	PackedXZ byte
	Y        int16
	Type     primitive.VarInt
	Data     nbt.NBT
}
