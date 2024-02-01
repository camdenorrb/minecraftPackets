package palette

import (
	"github.com/joomcode/errorx"
)

// TODO: Read from https://wiki.vg/Data_Generators

// TODO:
// 	generated blocks.json -> blocks.proto
// 	Have one struct representing original data
// 	Make map from struct to do state as string -> ID
// 	Store map in registry on registry load

// TODO:
// 	 For older versions which don't have blocks.json
//	 Have a way to go from map back to struct, to generate a ~"blocks.json"

// TODO: Move proto files to protocol folder based on version

type BlockPaletteJSON map[Material]BlockPaletteData

type BlockStateKey struct {
	Material   Material // Maybe use material instead
	Properties BlockStateProperties
}

type StateMap struct {
	stateToID map[uint64]uint32 // BlockStateKey as string for key
	IDToState map[uint32]BlockStateKey
}

func (m StateMap) StateToID(key *BlockStateKey) (*uint32, error) {

	hash, err := key.Properties.Hash()
	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "failed to hash block state")
	}

	id, exists := m.stateToID[*hash]
	if !exists {
		return nil, errorx.IllegalState.New("block state does not exist")
	}

	return &id, nil
}

func (p BlockPaletteJSON) AsBlockPalette() *BlockPalette {

	palette := &BlockPalette{Entries: make([]*BlockPaletteEntry, len(p))}

	for material, data := range p {
		entry := &BlockPaletteEntry{Material: material, Data: &data}
		palette.Entries = append(palette.Entries, entry)
	}

	return palette
}

func (p BlockPaletteJSON) AsStateMap() (*StateMap, error) {

	stateToID := make(map[uint64]uint32)        // Hash to ID
	idToState := make(map[uint32]BlockStateKey) // ID to BlockStateKey

	for material, data := range p {
		for _, state := range data.States {

			properties := state.Properties
			if properties == nil {
				properties = &BlockStateProperties{}
			}

			//goland:noinspection GoVetCopyLock
			key := BlockStateKey{
				Material:   material,
				Properties: *properties,
			}

			hash, err := state.Properties.Hash()
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "failed to hash block state %v", state)
			}

			stateToID[*hash] = state.Id
			idToState[state.Id] = key
		}
	}

	return &StateMap{
		stateToID: stateToID,
		IDToState: idToState,
	}, nil
}
