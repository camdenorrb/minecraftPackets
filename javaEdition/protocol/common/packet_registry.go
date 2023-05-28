package common

import "github.com/camdenorrb/minecraftPackets/javaEdition/common"

type PacketRegistry interface {
	Register(state State, id common.VarInt, packet interface{})
	GetPacket(state State, id common.VarInt) any
	GetID(state State, packet interface{}) common.VarInt
}

type packetRegistry struct {
	KeyToPacket map[State]map[common.VarInt]interface{}
	PacketToID  map[interface{}]map[State]common.VarInt
}

// Ideally we would get via something like PacketRegistry.GetPacketID(nil.(Type))
func NewPacketRegistry() PacketRegistry {
	return &packetRegistry{
		KeyToPacket: make(map[State]map[common.VarInt]interface{}),
		PacketToID:  make(map[interface{}]map[State]common.VarInt),
	}
}

func (r *packetRegistry) Register(state State, id common.VarInt, packet interface{}) {

	if r.KeyToPacket[state] == nil {
		r.KeyToPacket[state] = make(map[common.VarInt]interface{})
	}
	if r.PacketToID[packet] == nil {
		r.PacketToID[packet] = make(map[State]common.VarInt)
	}

	r.KeyToPacket[state][id] = packet
	r.PacketToID[packet][state] = id
}

func (r *packetRegistry) GetPacket(state State, id common.VarInt) any {
	return r.KeyToPacket[state][id]
}

func (r *packetRegistry) GetID(state State, packet interface{}) common.VarInt {
	return r.PacketToID[packet][state]
}
