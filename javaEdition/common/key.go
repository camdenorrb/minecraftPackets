package common

import "github.com/camdenorrb/minecraftPackets/primitive"

// PublicKey is a struct that represents a public key.
// Refer to https://wiki.vg/Protocol#Player_Session
type PublicKey struct {
	ExpiresAt          int64
	PublicKeyLength    primitive.VarInt
	PublicKey          []byte
	KeySignatureLength primitive.VarInt
	KeySignature       []byte
}
