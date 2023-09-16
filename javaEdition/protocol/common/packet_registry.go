package common

import (
	"github.com/camdenorrb/minecraftPackets/primitive"
)

type PacketRegistry interface {
	Register(state State, id primitive.VarInt, packet string)
	GetPacket(state State, id primitive.VarInt) *string
	GetID(state State, packet string) *primitive.VarInt
}

type packetRegistry struct {
	KeyToPacket map[State]map[primitive.VarInt]*string
	PacketToID  map[string]map[State]primitive.VarInt
}

// Ideally we would get via something like PacketRegistry.GetPacketID(nil.(Type))
func NewPacketRegistry() PacketRegistry {
	return &packetRegistry{
		KeyToPacket: make(map[State]map[primitive.VarInt]*string),
		PacketToID:  make(map[string]map[State]primitive.VarInt),
	}
}

func (r *packetRegistry) Register(state State, id primitive.VarInt, packet string) {

	if r.KeyToPacket[state] == nil {
		r.KeyToPacket[state] = make(map[primitive.VarInt]*string)
	}
	if r.PacketToID[packet] == nil {
		r.PacketToID[packet] = make(map[State]primitive.VarInt)
	}

	r.KeyToPacket[state][id] = &packet
	r.PacketToID[packet][state] = id
}

func (r *packetRegistry) GetPacket(state State, id primitive.VarInt) *string {

	value, exists := r.KeyToPacket[state][id]
	if !exists {
		return nil
	}

	return value
}

func (r *packetRegistry) GetID(state State, packet string) *primitive.VarInt {

	value, exists := r.PacketToID[packet][state]
	if !exists {
		return nil
	}

	return &value
}
