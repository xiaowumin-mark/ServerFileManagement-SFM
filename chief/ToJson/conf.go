package ToJson

import (
	"encoding/json"
	"fmt"
	"xiaowumin-SFM/Struct"
)

func ConfJson(jsonData []byte) (*Struct.Config, error) {
	var person Struct.Config
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println("JSON decoding error:", err)
		return nil, err
	}

	return &person, nil
}

func GetFileJson(jsonData []byte) (*Struct.GetFile, error) {
	var person Struct.GetFile
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println("JSON decoding error:", err)
		return nil, err
	}

	return &person, nil
}

func HostStateJson(jsonData []byte) (*Struct.HostState, error) {
	var person Struct.HostState
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println("JSON decoding error:", err)
		return nil, err
	}

	return &person, nil
}

func SearchFileJson(jsonData []byte) (*Struct.SearchFile, error) {
	var person Struct.SearchFile
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println("JSON decoding error:", err)
		return nil, err
	}

	return &person, nil
}
