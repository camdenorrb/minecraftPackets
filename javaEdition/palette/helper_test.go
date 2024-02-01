package palette

import (
	"encoding/json"
	"github.com/joomcode/errorx"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFromBlockPalette(t *testing.T) {

	a := assert.New(t)

	file, err := os.ReadFile("testdata/blocks1_20_4.json")
	a.NoError(err)

	var blocks BlockPaletteJSON
	err = json.Unmarshal(file, &blocks)
	if err != nil {
		a.NoError(errorx.Decorate(err, "failed to unmarshal json"))

	}
	//thing, err := json.Marshal(blocks)
	//a.NoError(err)
	//fmt.Println(string(thing))
	//fmt.Println(len(blocks))
	//fmt.Println(len(blocks))
	/*
		var blocks palette.BlockPaletteJSON
		a.NoError(json.Unmarshal(file, &blocks))


		blocksAsProto, err := FromBlockPalette(blocks)
		a.NoError(err)

		blockAsJSON, err := blocksAsProto.AsJSONStruct()
		a.NoError(err)

		for _, entry := range blocksAsProto.Entries {
			if entry.Material == Material_ACACIA_BUTTON {
				spew.Sdump(entry.Data.Properties)
			}
		}

	*/
	//a.Equal(blocks, blockAsJSON)
	/*a.Len(blocks["minecraft:acacia_button"].Properties.Face, 3)
	 */

	//goland:noinspection GoVetCopyLock
	indent, err := json.MarshalIndent(blocks, "", "  ")
	a.NoError(err)

	// Write to file
	a.NoError(os.WriteFile("testdata/blocks1_20_4.proto.json", indent, 0644))

	a.JSONEq(string(file), string(indent))

	asStateMap, err := blocks.AsStateMap()
	a.NoError(err)

	assert.NotEmpty(t, asStateMap.stateToID)
}
