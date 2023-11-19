package Struct

type GetFile struct {
	Main []struct {
		Isdir bool   `json:"isdir"`
		Name  string `json:"name"`
		Size  string `json:"size"`
		Time  string `json:"time"`
	} `json:"main"`
}
