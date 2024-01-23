package proto

import (
	"encoding/json"
	"github.com/camdenorrb/minecraftPackets/javaEdition/palette"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFromBlockPalette(t *testing.T) {

	a := assert.New(t)

	file, err := os.ReadFile("testdata/blocks1_20_4.json")
	a.NoError(err)

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

	//a.Equal(blocks, blockAsJSON)
	a.Len(blocks["minecraft:acacia_button"].Properties.Face, 3)

	blockAsJSONBytes, err := json.Marshal(blockAsJSON)
	a.NoError(err)

	indent, err := json.MarshalIndent(blockAsJSON, "", "  ")
	a.NoError(err)

	// Write to file
	a.NoError(os.WriteFile("testdata/blocks1_20_4.proto.json", indent, 0644))

	a.JSONEq(string(file), string(blockAsJSONBytes))
}
