package File

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"xiaowumin-SFM/Struct"
)

func CopyFile(date Struct.SandCopyFile) error {
	if date.Path[len(date.Path)-1] != '/' {
		date.Path += "/"
	}
	if date.ToPath[len(date.ToPath)-1] != '/' {
		date.ToPath += "/"
	}
	fmt.Println(date)
	for i := 0; i < len(date.CopyName); i++ {
		if date.CopyName[i].IsDir {
			var content string
			if date.ToPath[len(date.ToPath)-1] == '/' {
				content = date.ToPath[:len(date.ToPath)-1]
			}
			fmt.Println("[dir]path:", date.Path+date.CopyName[i].Name, "topath:", content)
			err := CopyDir(date.Path+date.CopyName[i].Name, content)
			if err != nil {
				return err
			}
			content = ""
		} else {
			var content string
			if date.ToPath[len(date.ToPath)-1] != '/' {
				content = date.ToPath + "/"
			}
			fmt.Println("[file]path:", date.Path+date.CopyName[i].Name, "topath:", content)
			err := copyFile(date.Path+date.CopyName[i].Name, content, 0)
			if err != nil {
				return err
			}
			content = ""
		}
	}
	return nil
}

// 使用os.Read()和os.Write()
func copyFile(src, des string, bufSize int) (err error) {
	if bufSize <= 0 {
		bufSize = 1 * 1024 * 1024 //1M
	}
	buf := make([]byte, bufSize)

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	//获取源文件的权限
	fi, _ := srcFile.Stat()
	perm := fi.Mode()

	desFile, err := os.OpenFile(des, os.O_CREATE|os.O_RDWR|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer desFile.Close()

	count := 0
	for {
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}

		if wn, err := desFile.Write(buf[:n]); err != nil {
			return err
		} else {
			count += wn
		}
	}

	return nil
}

func CopyDir(srcPath, desPath string) error {
	//检查目录是否正确
	if srcInfo, err := os.Stat(srcPath); err != nil {
		return err
	} else {
		if !srcInfo.IsDir() {
			return errors.New("源路径不是一个正确的目录！")
		}
	}

	if desInfo, err := os.Stat(desPath); err != nil {
		return err
	} else {
		if !desInfo.IsDir() {
			return errors.New("目标路径不是一个正确的目录！")
		}
	}

	if strings.TrimSpace(srcPath) == strings.TrimSpace(desPath) {
		return errors.New("源路径与目标路径不能相同！")
	}

	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		//复制目录是将源目录中的子目录复制到目标路径中，不包含源目录本身
		if path == srcPath {
			return nil
		}

		//生成新路径
		destNewPath := strings.Replace(path, srcPath, desPath, -1)

		if !f.IsDir() {
			copyFile(path, destNewPath, 0)
		} else {
			if !FileIsExisted(destNewPath) {
				return MakeDir(destNewPath)
			}
		}

		return nil
	})

	return err
}

func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func MakeDir(dir string) error {
	if !FileIsExisted(dir) {
		if err := os.MkdirAll(dir, 0777); err != nil { //os.ModePerm
			fmt.Println("MakeDir failed:", err)
			return err
		}
	}
	return nil
}
