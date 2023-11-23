package chief

import (
	"encoding/json"
	"fmt"
	"xiaowumin-SFM/Struct"
)

func InFoError(errs string) []byte {
	fmt.Println("错误: " + errs)
	errr := Struct.Error{
		Code: 200,
		Err:  errs,
	}

	// 将结构体转换为 JSON 格式
	jsonData, err := json.Marshal(errr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return jsonData
}
