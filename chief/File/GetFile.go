package File

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"xiaowumin-SFM/Struct"
)

var gf Struct.GetFile

func GetFile(dirname string) []byte {
	//dir := "D:/xiaowumin more" // 替换成目标文件夹的路径
	entries, err := os.ReadDir(dirname)
	if err != nil {
		//fmt.Println("读取目录失败:", err)
		//return
	}

	for _, entry := range entries {
		fileInfo, err := entry.Info()
		if err != nil {
			//	fmt.Println("无法获取文件信息:", err)

		}

		//fmt.Println("文件名:", entry.Name())
		//fmt.Println("大小:", formatFileSize(fileInfo.Size()))
		modTime := fileInfo.ModTime()
		// 判断年份是否为今年
		var FileTime_ string
		if modTime.Year() == time.Now().Year() {
			//	fmt.Println("修改时间:", modTime.Format("01月02日"))
			FileTime_ = modTime.Format("01月02日")
		} else {
			//	fmt.Println("修改时间:", modTime.Format("2006年01月02日"))
			FileTime_ = modTime.Format("2006年01月02日")
		}

		// 创建一个新的条目
		newEntry := struct {
			Isdir bool   `json:"isdir"`
			Name  string `json:"name"`
			Size  string `json:"size"`
			Time  string `json:"time"`
		}{
			Isdir: fileInfo.IsDir(),                // 这里可以设置你想要的值
			Name:  entry.Name(),                    // 同上
			Size:  formatFileSize(fileInfo.Size()), // 同上
			Time:  FileTime_,                       // 同上
		}

		// 将新的条目添加到Main字段中
		gf.Main = append(gf.Main, newEntry)

	}
	jsonData, err := json.Marshal(gf)
	if err != nil {
		fmt.Println(err)

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
