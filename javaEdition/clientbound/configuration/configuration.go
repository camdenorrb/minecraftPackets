package configuration

import (
	"github.com/camdenorrb/minecraftPackets/javaEdition/common"
	"github.com/camdenorrb/minecraftPackets/nbt"
	"github.com/camdenorrb/minecraftPackets/primitive"
)

type PluginMessage struct {
	Channel primitive.Identifier
	Data    []byte
}

type Disconnect struct {
	Reason common.Chat
}

type Finish struct{}

type KeepAlive struct {
	ID int64
}

type Ping struct {
	ID int64
}

type RegistryData struct {
	RegistryCodec nbt.NBT
}

type ResourcePack struct {
	URL           string
	Hash          string
	IsForced      bool
	PromptMessage *common.Chat
}

type FeatureFlags struct {
	FeatureFlags []primitive.Identifier
}
