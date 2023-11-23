package Struct

type GetFile struct {
	Main []struct {
		Isdir bool   `json:"isdir"`
		Name  string `json:"name"`
		Size  string `json:"size"`
		Time  string `json:"time"`
	} `json:"main"`
}

type SearchFile struct {
	KeyWord string `json:"KeyWord"`
	Number  int    `json:"Number"`
	Path    string `json:"Path"`
	Type    string `json:"Type"`
	Main    []struct {
		Isdir bool   `json:"isdir"`
		Name  string `json:"name"`
		Size  string `json:"size"`
		Time  string `json:"time"`
		Path  string `json:"path"`
	} `json:"main"`
}
