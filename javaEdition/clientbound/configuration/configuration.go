package configuration

import (
	"github.com/camdenorrb/minecraftPackets/javaEdition/common"
	common2 "github.com/camdenorrb/minecraftPackets/javaEdition/protocol/common"
	"github.com/camdenorrb/minecraftPackets/nbt"
	"github.com/camdenorrb/minecraftPackets/primitive"
)

type PluginMessage struct {
	Channel primitive.Identifier
	Data    []byte
}

func (*PluginMessage) PacketType() string {
	return string(common2.ClientBoundConfigurationPluginMessage)
}

type Disconnect struct {
	Reason common.Chat
}

func (*Disconnect) PacketType() string {
	return string(common2.ClientBoundConfigurationDisconnect)
}

type Finish struct{}

func (*Finish) PacketType() string {
	return string(common2.ClientBoundConfigurationFinish)
}

type KeepAlive struct {
	ID int64
}

func (*KeepAlive) PacketType() string {
	return string(common2.ClientBoundConfigurationKeepAlive)
}

type Ping struct {
	ID int64
}

func (*Ping) PacketType() string {
	return string(common2.ClientBoundConfigurationPing)
}

type RegistryData struct {
	RegistryCodec nbt.NBT
}

func (*RegistryData) PacketType() string {
	return string(common2.ClientBoundConfigurationRegistryData)
}

type ResourcePack struct {
	URL           string
	Hash          string
	IsForced      bool
	PromptMessage *common.Chat
}

func (*ResourcePack) PacketType() string {
	return string(common2.ClientBoundConfigurationResourcePack)
}

type FeatureFlags struct {
	FeatureFlags []primitive.Identifier
}

func (*FeatureFlags) PacketType() string {
	return string(common2.ClientBoundConfigurationFeatureFlags)
}

type UpdateTags struct {
	TagCount primitive.VarInt
	//Tags     []Tag // TODO: Tag
}

func (*UpdateTags) PacketType() string {
	return string(common2.ClientBoundConfigurationUpdateTags)
}
