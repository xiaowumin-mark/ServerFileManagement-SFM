package File

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"xiaowumin-SFM/Struct"
	"xiaowumin-SFM/chief"
	"xiaowumin-SFM/chief/ToJson"
)

var sf = Struct.SearchFile{}
var jsonData []byte

func SearchFile(path string, name string, all bool) []byte {
	if path[len(path)-1] != '/' {
		path += "/"
	}
	if all == false {
		entries, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("读取目录失败:", err)

		}
		var SearchFilenum int
		for _, entry := range entries {

			file := path + entry.Name()
			nenn := strings.Split(file, `\`)
			//fmt.Println(nenn[len(nenn)-1])
			if strings.Contains(nenn[len(nenn)-1], name) == true {
				fileInfo, err := entry.Info()
				if err != nil {
					fmt.Println("无法获取文件信息:", err)

				}

				modTime := fileInfo.ModTime()
				var FileTime_ string
				if modTime.Year() == time.Now().Year() {
					FileTime_ = modTime.Format("01月02日")
				} else {
					FileTime_ = modTime.Format("2006年01月02日")
				}

				SearchFileMain := struct {
					Isdir bool   `json:"isdir"`
					Name  string `json:"name"`
					Size  string `json:"size"`
					Time  string `json:"time"`
					Path  string `json:"path"`
				}{
					Isdir: fileInfo.IsDir(),
					Name:  entry.Name(),
					Size:  formatFileSize(fileInfo.Size()),
					Time:  FileTime_,
					Path:  file,
				}

				sf.Main = append(sf.Main, SearchFileMain)
				SearchFilenum++
			}
		}
		sf.KeyWord = name
		sf.Path = path
		sf.Type = "AtPersrnt"
		sf.Number = SearchFilenum
		jsonData, err = json.Marshal(sf)
		if err != nil {
			fmt.Println(err)

		}
		//fmt.Println(string(jsonData))

	} else {
		decodedPerson, err := ToJson.ConfJson(chief.Config())
		if err != nil {
			fmt.Println("解码错误:", err)

		}
		var SearchFilenum int
		fmt.Println(decodedPerson.Disks.Name)
		for i := 0; i < len(decodedPerson.Disks.Name); i++ {
			content := decodedPerson.Disks.Name[i][:len(decodedPerson.Disks.Name[i])-1] + "/"

			root := content // 替换为你要遍历的目录的路径

			err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					fmt.Println(err)
					return nil
				}

				// 打印目录信息

				nenn := strings.Split(path, `\`)
				//fmt.Println(nenn[len(nenn)-1])
				if strings.Contains(nenn[len(nenn)-1], name) == true {

					modTime := info.ModTime()
					var FileTime_ string
					if modTime.Year() == time.Now().Year() {
						FileTime_ = modTime.Format("01月02日")
					} else {
						FileTime_ = modTime.Format("2006年01月02日")
					}

					SearchFileMain := struct {
						Isdir bool   `json:"isdir"`
						Name  string `json:"name"`
						Size  string `json:"size"`
						Time  string `json:"time"`
						Path  string `json:"path"`
					}{
						Isdir: info.IsDir(),
						Name:  info.Name(),
						Size:  formatFileSize(info.Size()),
						Time:  FileTime_,
						Path:  path,
					}

					sf.Main = append(sf.Main, SearchFileMain)
					SearchFilenum++
				}

				return nil
			})
			if err != nil {
				fmt.Println(err)
			}
		}
		sf.KeyWord = name
		sf.Path = path
		sf.Type = "All"
		sf.Number = SearchFilenum
		jsonData, err = json.Marshal(sf)
		if err != nil {
			fmt.Println(err)

		}
		//fmt.Println(string(jsonData))

	}
	return jsonData
}
