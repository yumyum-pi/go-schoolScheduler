package utils

import (
	"encoding/json"
	"io/ioutil"
)

// WriteFile write data to the given path
// from the given interface
// TODO add error check
func WriteFile(path string, t interface{}) {
	file, _ := json.Marshal(t)

	_ = ioutil.WriteFile(path, file, 0644)
}
