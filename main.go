package main

import (
	"fmt"
	"xiaowumin-SFM/chief"
	"xiaowumin-SFM/chief/File"
	"xiaowumin-SFM/chief/ToJson"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Setting := r.Group("config/")
	File_ := r.Group("file/")
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
	File_.POST("/", func(c *gin.Context) {
		path := c.PostForm("path")
		
		decodedPerson, err := ToJson.GetFileJson(File.GetFile(path))
		if err != nil {
			fmt.Println("解码错误:", err)
			return
		}
		c.JSON(200, decodedPerson)
		fmt.Println(path)
	})
	r.Run(":8080")

}
