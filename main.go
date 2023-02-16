import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Define the flags
	inputFile := flag.String("input", "react-code.js", "the input file name")
	outputFile := flag.String("output", "transpiled-code.js", "the output file name")
	flag.Parse()

	// Load the uwu keyword pairs from the uwu.json file
	file, err := os.Open("uwu.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var uwuKeywords map[string]string
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &uwuKeywords)

	// Load the React code to be transpiled
	file, err = os.Open(*inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteValue, _ = ioutil.ReadAll(file)
	reactCode := string(byteValue)

	// Replace the uwu keywords with the React keywords
	for uwuKeyword, reactKeyword := range uwuKeywords {
		reactCode = strings.ReplaceAll(reactCode, uwuKeyword, reactKeyword)
	}

	// Write the transpiled code to a new file
	err = ioutil.WriteFile(*outputFile, []byte(reactCode), 0644)
	if err != nil {
		panic(err)
	}
}
