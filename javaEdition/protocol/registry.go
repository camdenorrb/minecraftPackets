package protocol

import (
	"github.com/camdenorrb/minecraftPackets/javaEdition/common"
	version119 "github.com/camdenorrb/minecraftPackets/javaEdition/protocol/version119"
)

type RegistryCache map[Version]map[Bound]PacketRegistry

type PacketRegistry interface {
	Register(state State, id common.VarInt, packet interface{})
	GetPacket(state State, id common.VarInt) any
	GetID(state State, packet interface{}) common.VarInt
}

type packetRegistry struct {
	KeyToPacket map[State]map[common.VarInt]interface{}
	PacketToID  map[interface{}]map[State]common.VarInt
}

func NewPacketRegistry() PacketRegistry {
	return &packetRegistry{
		KeyToPacket: make(map[State]map[common.VarInt]interface{}),
		PacketToID:  make(map[interface{}]map[State]common.VarInt),
	}
}

func (r *packetRegistry) Register(state State, id common.VarInt, packet interface{}) {
	r.KeyToPacket[state][id] = packet
	r.PacketToID[packet][state] = id
}

func (r *packetRegistry) GetPacket(state State, id common.VarInt) any {
	return r.KeyToPacket[state][id]
}

func (r *packetRegistry) GetID(state State, packet interface{}) common.VarInt {
	return r.PacketToID[packet][state]
}

func (r RegistryCache) RegistryFor(bound Bound, version Version) PacketRegistry {

	// If the map doesn't exist, create it
	if _, ok := r[version]; !ok {
		r[version] = make(map[Bound]PacketRegistry)
	}

	// If the registry doesn't exist, create it
	if _, ok := r[version][bound]; !ok {
		r[version][bound] = RegistryFor(bound, version)
	}

	return r[version][bound]
}

func RegistryFor(bound Bound, version Version) PacketRegistry {

	switch version {
	case Version_1_8:
		return version119.Registry(bound)
	}

	return nil
}