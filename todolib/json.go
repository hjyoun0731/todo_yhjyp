package todolib

import (
	"encoding/json"
)

type todoData struct {
	version	version
	account account
}
type version struct {
	Version	string
	Name	string
	UpdateTime	string
}
type account struct {
	UserId			string
	Password	string
}

// JSONEnc string to json
func JSONEnc(arg interface{}) string {
	jsonBytes, err := json.Marshal(arg)
	if err != nil {
		panic(err)
	}
	jsonString := string(jsonBytes)

	MakeLog("JSONEnc success")

	return jsonString
}

// JSONDecVer json to string
func JSONDecVer(data []byte) version{
	var mem version
	err := json.Unmarshal(data, &mem)
	if err != nil {
		panic(err)
	}
	MakeLog("JSONDecVer success")
	return mem
}

// JSONDecAcct json to string
func JSONDecAcct(data []byte) account{
	var mem account
	err := json.Unmarshal(data, &mem)
	if err != nil {
		panic(err)
	}
	MakeLog("JSONDecAcct success")
	return mem
}
