package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// UwuMap is a type that maps uwu keywords to React keywords
type UwuMap map[string]string

func main() {
	// Read uwu keyword pairs from a JSON file
	file, err := ioutil.ReadFile("uwuMap.json")
	if err != nil {
		fmt.Println("Error reading uwuMap.json:", err)
		return
	}

	// Parse JSON into a map of uwu keyword pairs
	var uwuMap UwuMap
	err = json.Unmarshal(file, &uwuMap)
	if err != nil {
		fmt.Println("Error parsing uwuMap.json:", err)
		return
	}

	// Print out the uwu keyword pairs
	for uwu, react := range uwuMap {
		fmt.Println(uwu, "->", react)
	}
}
