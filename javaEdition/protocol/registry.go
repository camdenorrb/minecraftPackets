package protocol

import (
	"github.com/camdenorrb/minecraftPackets/javaEdition/protocol/common"
	version119 "github.com/camdenorrb/minecraftPackets/javaEdition/protocol/version119"
)

type RegistryCache map[common.Version]map[common.Bound]common.PacketRegistry

func (r RegistryCache) RegistryFor(bound common.Bound, version common.Version) common.PacketRegistry {

	// If the map doesn't exist, create it
	if _, ok := r[version]; !ok {
		r[version] = make(map[common.Bound]common.PacketRegistry)
	}

	// If the registry doesn't exist, create it
	if _, ok := r[version][bound]; !ok {
		r[version][bound] = RegistryFor(bound, version)
	}

	return r[version][bound]
}

func RegistryFor(bound common.Bound, version common.Version) common.PacketRegistry {

	switch version {
	case common.Version_1_8:
		return version119.Registry(bound)
	}

	return nil
}
