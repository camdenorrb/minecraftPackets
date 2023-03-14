package common

// PublicKey is a struct that represents a public key.
// Refer to https://wiki.vg/Protocol#Player_Session
type PublicKey struct {
	ExpiresAt          int64
	PublicKeyLength    VarInt
	PublicKey          []byte
	KeySignatureLength VarInt
	KeySignature       []byte
}
