package main

import (
	"fmt"
	"xiaowumin-SFM/chief"
	"xiaowumin-SFM/chief/ToJson"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World!!")
	fmt.Println()
	r := gin.Default()
	Setting := r.Group("config/")
	Setting.POST("/disk", func(c *gin.Context) {
		decodedPerson, err := ToJson.ConfJson(chief.Config())
		if err != nil {
			fmt.Println("解码错误:", err)
			return
		}
		c.JSON(200, decodedPerson)
	})

	Setting.POST("/state", func(c *gin.Context) {

		c.JSON(200, chief.GetHostState())

	})

	r.Run(":8080")

}
