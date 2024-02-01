package main

import (
	"encoding/json"
	"fmt"
	"github.com/camdenorrb/minecraftPackets/javaEdition/palette"
	"github.com/joomcode/errorx"
	"google.golang.org/protobuf/proto"
	"os"
	"sort"
	"strings"
)

type blocksJSON map[string]map[string]any

func main() {

	file, err := os.ReadFile("palette/testdata/blocks1_20_4.json")
	if err != nil {
		panic(errorx.Panic(err))
	}

	var blocks palette.BlockPaletteJSON
	if err = json.Unmarshal(file, &blocks); err != nil {
		panic(errorx.Panic(err))
	}

	fmt.Println(len(blocks))

	for _, block := range blocks {
		fmt.Printf("%+v\n", block)
	}

	//stateMap := blocks.AsStateMap()

	//for key, id := range stateMap.StateToID {
	//fmt.Println(key, id)
	//}

	//printOutPaletteProperties("palette/testdata/blocks1_20_4.json")

	blockProtoFile, err := os.Create("palette/blocks1_20_4.proto.bin")
	if err != nil {
		panic(errorx.Panic(err))
	}
	defer blockProtoFile.Close()

	marshalled, err := proto.Marshal(blocks.AsBlockPalette())
	if err != nil {
		panic(errorx.Panic(err))
	}

	if _, err = blockProtoFile.Write(marshalled); err != nil {
		panic(errorx.Panic(err))
	}

	//fmt.Println(size.Of(jsonStruct.AsStateMap()))
	//printOutPaletteMaterials("palette/blocks1_20_4.json", true)
}

func stringArraysEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func printOutPaletteMaterials(file string, asProto bool) {

	var blocks blocksJSON

	data, err := os.ReadFile(file)
	if err != nil {
		panic(errorx.Panic(err))
	}

	// Read bytes
	if err = json.Unmarshal(data, &blocks); err != nil {
		panic(errorx.Panic(err))
	}

	index := 0
	for material := range blocks {

		if asProto {
			key := strings.ToUpper(strings.TrimPrefix(material, "minecraft:"))
			fmt.Println(fmt.Sprintf("%v = %d;", key, index))
		} else {
			fmt.Println(material)
		}

		index++
	}
}

func printOutPaletteProperties(file string) {

	var blocks blocksJSON

	data, err := os.ReadFile(file)
	if err != nil {
		panic(errorx.Panic(err))
	}

	// Read bytes
	if err = json.Unmarshal(data, &blocks); err != nil {
		panic(errorx.Panic(err))
	}

	allProperties := map[string][]string{}
	for _, mapThing := range blocks {

		thing := mapThing["properties"]
		if thing == nil {
			continue
		}

		thingMap := thing.(map[string]any)
		for key, value := range thingMap {

			// Unmarshal []interface{} to []string
			marshal, err := json.Marshal(value)
			if err != nil {
				panic(errorx.Panic(err))
			}

			var values []string
			if err = json.Unmarshal(marshal, &values); err != nil {
				panic(errorx.Panic(err))
			}

			// Check if values is the same
			if allProperties[key] != nil && !stringArraysEqual(allProperties[key], values) {
				// Pick the longer one
				if len(allProperties[key]) > len(values) {
					values = allProperties[key]
				}
			}

			allProperties[key] = values
		}

	}

	// Sort keys
	var keys []string
	for key := range allProperties {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for key := range keys {
		values := allProperties[keys[key]]
		print(keys[key] + ": ")
		for i, value := range values {
			print(value)
			if i != len(values)-1 {
				print(", ")
			}
		}
		println()
	}
}
