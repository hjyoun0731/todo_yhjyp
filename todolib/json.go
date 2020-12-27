package todolib

import (
	"encoding/json"
	"fmt"
)

type todoData struct {
	Version    string
	Name       string
	UpdateTime int
}

// JSONEnc string to json
func JSONEnc(version string, name string, updateTime int) string {

	mem := todoData{version, name, updateTime}

	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}
	jsonString := string(jsonBytes)

	return jsonString
}

// JSONDec json to string
func JSONDec() {
	fmt.Println("Decoding...")
}
