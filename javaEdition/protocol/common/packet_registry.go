package common

import "github.com/camdenorrb/minecraftPackets/javaEdition/common"

type PacketRegistry interface {
	Register(state State, id common.VarInt, packet string)
	GetPacket(state State, id common.VarInt) *string
	GetID(state State, packet string) common.VarInt
}

type packetRegistry struct {
	KeyToPacket map[State]map[common.VarInt]*string
	PacketToID  map[string]map[State]common.VarInt
}

// Ideally we would get via something like PacketRegistry.GetPacketID(nil.(Type))
func NewPacketRegistry() PacketRegistry {
	return &packetRegistry{
		KeyToPacket: make(map[State]map[common.VarInt]*string),
		PacketToID:  make(map[string]map[State]common.VarInt),
	}
}

func (r *packetRegistry) Register(state State, id common.VarInt, packet string) {

	if r.KeyToPacket[state] == nil {
		r.KeyToPacket[state] = make(map[common.VarInt]*string)
	}
	if r.PacketToID[packet] == nil {
		r.PacketToID[packet] = make(map[State]common.VarInt)
	}

	r.KeyToPacket[state][id] = &packet
	r.PacketToID[packet][state] = id
}

func (r *packetRegistry) GetPacket(state State, id common.VarInt) *string {

	value, exists := r.KeyToPacket[state][id]
	if !exists {
		return nil
	}

	return value
}

func (r *packetRegistry) GetID(state State, packet string) common.VarInt {
	return r.PacketToID[packet][state]
}
