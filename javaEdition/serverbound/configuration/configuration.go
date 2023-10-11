package configuration

import (
	common2 "github.com/camdenorrb/minecraftPackets/javaEdition/protocol/common"
	"github.com/camdenorrb/minecraftPackets/primitive"
)

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

func (*ClientInformation) PacketType() string {
	return string(common2.ServerBoundConfigurationClientInformation)
}

type PluginMessage struct {
	Channel primitive.Identifier
	Data    []byte
}

func (*PluginMessage) PacketType() string {
	return string(common2.ServerBoundConfigurationPluginMessage)
}

type Finish struct{}

func (*Finish) PacketType() string {
	return string(common2.ServerBoundConfigurationFinish)
}

type KeepAlive struct {
	ID int64
}

func (*KeepAlive) PacketType() string {
	return string(common2.ServerBoundConfigurationKeepAlive)
}

type Pong struct {
	ID int64
}

func (*Pong) PacketType() string {
	return string(common2.ServerBoundConfigurationPong)
}

type ResourcePackResponse struct {
	Result primitive.VarInt // TODO: enum
}

func (*ResourcePackResponse) PacketType() string {
	return string(common2.ServerBoundConfigurationResourcePackResponse)
}
