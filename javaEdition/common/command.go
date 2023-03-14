package common

// CommandNode is a node in a command tree.
// Refer to https://wiki.vg/Command_Data
type CommandNode struct {
	Flags         byte // TODO: Enum
	ChildrenCount VarInt
	Children      []VarInt
	RedirectNode  *VarInt
	Name          *string
	ParserID      *VarInt
	//Properties    *CommandNodeProperties // TODO: Figure out
	SuggestionsType *Identifier
}

type ChatCommandSignature struct {
	ArgumentName string
	Signature    []byte
}
