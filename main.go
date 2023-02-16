package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// Load the file to transpile
	filename := "example.jsx"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Define the UWU keywords and their replacements
	replacements := map[string]string{
		"uwu": "React",
		"nuzzles": "render",
		"boops": "props",
		"snuggles": "state",
		"glomp": "componentDidMount",
	}

	// Replace the keywords in the file content
	for uwu, react := range replacements {
		content = []byte(strings.ReplaceAll(string(content), uwu, react))
	}

	// Write the transpiled file back to disk
	err = ioutil.WriteFile(filename, content, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Transpilation complete!")
}
