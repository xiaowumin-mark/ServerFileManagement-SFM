package Struct

type GetFile struct {
	Err    string `json:"err"`
	Number int    `json:"Number"`
	Path   string `json:"Path"`
	Main   []struct {
		Isdir bool   `json:"isdir"`
		Name  string `json:"name"`
		Size  string `json:"size"`
		Time  string `json:"time"`
		Path  string `json:"path"`
	} `json:"main"`
}

type SearchFile struct {
	Err     string `json:"err"`
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

type SandRemoveFile struct {
	Path       string `json:"path"`
	RemoveName []struct {
		Isdir bool   `json:"isdir"`
		Name  string `json:"name"`
	} `json:"removename"`
	User `json:"user"`
}

type SandSearchFile struct {
	Path    string `json:"path"`
	KeyWord string `json:"keyword"`
	Type    string `json:"type"`
	User    `json:"user"`
}

type SandGetFile struct {
	Path string `json:"path"`
	User `json:"user"`
}

type SandRenameFile struct {
	Path   string `json:"path"`
	Name   string `json:"name"`
	Rename string `json:"rename"`
	User   `json:"user"`
}

type SandCopyFile struct {
	Path     string `json:"path"`
	CopyName []struct {
		Name  string `json:"name"`
		IsDir bool   `json:"isdir"`
	} `json:"copyname"`
	ToPath string `json:"topath"`
	User   `json:"user"`
}
