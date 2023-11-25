package File

import (
	"fmt"
	"os"
	"xiaowumin-SFM/Struct"
)

func RenameFile(date Struct.SandRenameFile) error {
	if date.Path[len(date.Path)-1] != '/' {
		date.Path += "/"
	}
	fmt.Println(date)
	err := os.Rename(date.Path+date.Name, date.Path+date.Rename)
	if err != nil {
		return err
	}
	return nil
}
