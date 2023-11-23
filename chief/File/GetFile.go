package File

import (
 "encoding/json"
 "fmt"
 "os"
 "time"
 "xiaowumin-SFM/Struct"
)

var gf = Struct.GetFile{} // 将gf的定义移出函数之外，使其成为全局变量

func GetFile(dirname string) []byte {
 // 在每次调用GetFile之前将gf重置为初始状态
 gf = Struct.GetFile{}

 // 读取目录内容
 entries, err := os.ReadDir(dirname)
 if err != nil {
 fmt.Println("读取目录失败:", err)
 return nil
 }

 for _, entry := range entries {
 fileInfo, err := entry.Info()
 if err != nil {
 fmt.Println("无法获取文件信息:", err)
 continue
 }

 modTime := fileInfo.ModTime()
 var FileTime_ string
 if modTime.Year() == time.Now().Year() {
 FileTime_ = modTime.Format("01月02日")
 } else {
 FileTime_ = modTime.Format("2006年01月02日")
 }

 newEntry := struct {
 Isdir bool   `json:"isdir"`
 Name  string `json:"name"`
 Size  string `json:"size"`
 Time  string `json:"time"`
 }{
 Isdir: fileInfo.IsDir(),
 Name:  entry.Name(),
 Size:  formatFileSize(fileInfo.Size()),
 Time:  FileTime_,
 }

 gf.Main = append(gf.Main, newEntry)
 }

 jsonData, err := json.Marshal(gf)
 if err != nil {
 fmt.Println(err)
 return nil
 }
 fmt.Println(string(jsonData))
 return jsonData
}

func formatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fPB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}
