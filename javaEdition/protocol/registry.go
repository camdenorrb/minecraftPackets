package protocol

import (
	"github.com/camdenorrb/minecraftPackets/javaEdition/protocol/common"
	"github.com/camdenorrb/minecraftPackets/javaEdition/protocol/version119"
	"github.com/camdenorrb/minecraftPackets/javaEdition/protocol/version120"
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
	case common.Version_1_19, common.Version_1_19_2, common.Version_1_19_3, common.Version_1_19_4:
		return version119.Registry(bound)
	case common.Version_1_20: // 1.20.1 has same protocol as 1.20
		return version120.Registry(bound)
	}

	return nil
}
