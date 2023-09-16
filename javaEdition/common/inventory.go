package common

import (
	"github.com/camdenorrb/minecraftPackets/nbt"
	"github.com/camdenorrb/minecraftPackets/primitive"
)

type Slot struct {
	IsPresent bool
	ItemID    primitive.VarInt
	ItemCount byte
	NBTData   nbt.NBT
}
