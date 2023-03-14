package common

import "minecraftPackets/nbt"

type Slot struct {
	IsPresent bool
	ItemID    VarInt
	ItemCount byte
	NBTData   nbt.NBT
}
