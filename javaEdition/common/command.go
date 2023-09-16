package common

import "github.com/camdenorrb/minecraftPackets/primitive"

// CommandNode is a node in a command tree.
// Refer to https://wiki.vg/Command_Data
type CommandNode struct {
	Flags         byte // TODO: Enum
	ChildrenCount primitive.VarInt
	Children      []primitive.VarInt
	RedirectNode  *primitive.VarInt
	Name          *string
	ParserID      *primitive.VarInt
	//Properties    *CommandNodeProperties // TODO: Figure out
	SuggestionsType *primitive.Identifier
}

type ChatCommandSignature struct {
	ArgumentName string
	Signature    []byte
}
