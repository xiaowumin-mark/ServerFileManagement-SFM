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
