package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/png"
	"os"
)

func printUsage() {
	fmt.Printf("Usage: %s map.png format.json", os.Args[0])
}

func main() {
	if len(os.Args) != 3 {
		printUsage()
		panic(errors.New("Invalid number of arguments"))
	}

	mapFile, err := os.Open(os.Args[1])
	check(err)
	defer mapFile.Close()

	formatFile, err := os.Open(os.Args[2])
	check(err)
	defer formatFile.Close()

	mapImage, err := png.Decode(mapFile)
	check(err)

	formatJSON := map[string]interface{}{}
	decoder := json.NewDecoder(formatFile)
	err = decoder.Decode(&formatJSON)
	check(err)

	mapJSON := map[string]interface{}{}
	bounds := mapImage.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	mapJSON["width"] = w
	mapJSON["height"] = h
	tiles := make([]interface{}, 0, 0)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			rt, gt, bt, _ := mapImage.At(x, y).RGBA()
			r, g, b := uint8(rt), uint8(gt), uint8(bt)
			colorHex := fmt.Sprintf("0x%.2x%.2x%.2x", r, g, b)
			if formatJSON[colorHex] != nil {
				tile := map[string]interface{}{}
				tile["x"] = x
				tile["y"] = y
				tile["data"] = formatJSON[colorHex]
				tiles = append(tiles, tile)
			}
		}
	}
	mapJSON["tiles"] = tiles
	jsonBuf, err := json.MarshalIndent(mapJSON, "", "    ")
	fmt.Println(string(jsonBuf))
}

func check(e error) {
	if e != nil {
		printUsage()
		panic(e.Error())
	}
}
