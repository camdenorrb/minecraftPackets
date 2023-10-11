package configuration

import "github.com/camdenorrb/minecraftPackets/primitive"

type ClientInformation struct {
	Locale              string
	ViewDistance        byte
	ChatMode            primitive.VarInt // TODO: enum
	ChatColors          bool
	DisplayedSkinParts  byte             // TODO: enum
	MainHand            primitive.VarInt // TODO: enum
	EnableTextFiltering bool
	AllowServerListings bool
}

type PluginMessage struct {
	Channel primitive.Identifier
	Data    []byte
}

type Finish struct{}

type KeepAlive struct {
	ID int64
}

type Pong struct {
	ID int64
}

type ResourcePackResponse struct {
	Result primitive.VarInt // TODO: enum
}
