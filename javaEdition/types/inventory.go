package types

type Slot struct {
	IsPresent bool
	ItemID    VarInt
	ItemCount byte
	//NBTData   NBT // TODO: NBT
}
