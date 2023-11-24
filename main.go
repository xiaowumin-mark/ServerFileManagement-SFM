package main

import (
	"fmt"
	"xiaowumin-SFM/chief"
	"xiaowumin-SFM/chief/File"
	"xiaowumin-SFM/chief/ToJson"

	"github.com/gin-gonic/gin"
)

func main() {
	//File.SearchFile("D:/xiaowumin more", "main", true)
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
		decodedPerson, err := ToJson.HostStateJson(chief.GetHostState())
		if err != nil {
			fmt.Println("解码错误:", err)
			return
		}
		c.JSON(200, decodedPerson)

	})
	File_.POST("/", func(c *gin.Context) {
		path := c.PostForm("path")

		decodedPerson, err := ToJson.GetFileJson(File.GetFile(path))
		if err != nil {
			fmt.Println("解码错误:", err)
			return
		}
		c.JSON(200, decodedPerson)
		//fmt.Println(path)
	})

	File_.POST("/SearchFile", func(c *gin.Context) {
		path := c.PostForm("path")
		keyword := c.PostForm("keyword")
		Type := c.PostForm("type")

		decodedPerson, err := ToJson.SearchFileJson(File.SearchFile(path, keyword, Type))
		if err != nil {
			fmt.Println("解码错误:", err)
			return
		}
		c.JSON(200, decodedPerson)
		//fmt.Println(path)
	})

	r.Run(":8080")

}
