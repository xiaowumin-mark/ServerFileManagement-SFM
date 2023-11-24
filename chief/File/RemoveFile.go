package File

import (
	"fmt"
	"os"
	"xiaowumin-SFM/Struct"
)

func RemoveFile(date Struct.SandRemoveFile) error {
	for i := 0; i < len(date.RemoveName); i++ {
		if date.Path[len(date.Path)-1] != '/' {
			date.Path += "/"
		}
		if date.RemoveName[i].Isdir == true {
			err := os.RemoveAll(date.Path + date.RemoveName[i].Name)
			if err != nil {
				fmt.Println(err)
				return err
			}
		} else if date.RemoveName[i].Isdir == false {
			err := os.Remove(date.Path + date.RemoveName[i].Name)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}
	return nil
}
