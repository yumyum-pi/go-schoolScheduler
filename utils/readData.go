package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Path for resources
const Path = "./resource/"

// ResourceFilePath returns the file path
func ResourceFilePath(name string) string {
	return fmt.Sprintf("%v%v.json", Path, name)
}

// ReadFile read data from the given path
// and adds it to the given interface
func ReadFile(path string, t interface{}) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		// throw an error
		fmt.Printf("> Error:\n> %v", err)
	}

	fmt.Printf("\t>> File in path:\"%v\", successfully opened.\n", path)

	// get byte data from file
	byteValue, _ := ioutil.ReadAll(file)
	// unmarshal the byteValue
	json.Unmarshal(byteValue, t)
}
