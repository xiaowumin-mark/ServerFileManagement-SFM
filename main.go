package main

import (
	"encoding/json"
	"fmt"
	"io"
	"xiaowumin-SFM/Struct"
	"xiaowumin-SFM/chief"
	"xiaowumin-SFM/chief/File"

	"github.com/gin-gonic/gin"
)

func main() {
	//File.SearchFile("D:/xiaowumin more", "main", true)
	r := gin.Default()
	Setting := r.Group("config/")
	File_ := r.Group("file/")
	Setting.POST("/disk", func(c *gin.Context) {

		decodedPerson, err := chief.Config()
		if err != nil {
			c.JSON(200, gin.H{
				"err": err,
			})
			return
		}
		c.JSON(200, decodedPerson)
	})

	Setting.POST("/state", func(c *gin.Context) {
		decodedPerson, err := chief.GetHostState()
		if err != nil {
			c.JSON(200, gin.H{
				"err": err,
			})
			return
		}
		c.JSON(200, decodedPerson)

	})
	File_.POST("/", func(c *gin.Context) {
		GetFile := Struct.SandGetFile{}
		if err := c.BindJSON(&GetFile); err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{
				"err": "传入参数错误！",
			})
			return
		}

		decodedPerson, err := File.GetFile(GetFile.Path)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{
				"err": "解码错误！",
			})
			return
		}
		c.JSON(200, decodedPerson)
		//fmt.Println(path)
	})

	File_.POST("/SearchFile", func(c *gin.Context) {
		var SearchFile Struct.SandSearchFile
		if err := c.BindJSON(&SearchFile); err != nil {
			c.JSON(200, gin.H{
				"err": err,
			})
			return
		}

		decodedPerson, err := File.SearchFile(SearchFile.Path, SearchFile.KeyWord, SearchFile.Type)
		if err != nil {
			c.JSON(200, gin.H{
				"err": err,
			})
			return
		}
		c.JSON(200, decodedPerson)
		//fmt.Println(path)
	})
	File_.POST("/RemoveFile", func(c *gin.Context) {
		var RemoveFilest Struct.SandRemoveFile
		if err := c.ShouldBind(&RemoveFilest); err != nil {
			c.JSON(200, gin.H{
				"err": "传入参数错误！",
			})
			return
		}
		if err := File.RemoveFile(RemoveFilest); err != nil {
			c.JSON(200, gin.H{
				"main": "err",
			})
		} else {
			c.JSON(200, gin.H{
				"main": "ok",
			})
		}

	})
	File_.POST("/RenameFile", func(c *gin.Context) {
		var RenameFile = Struct.SandRenameFile{}
		if err := c.ShouldBind(&RenameFile); err != nil {
			c.JSON(200, gin.H{
				"err": "传入参数错误！",
			})
			return
		}
		if err := File.RenameFile(RenameFile); err != nil {
			c.JSON(200, gin.H{
				"main": err,
			})
		} else {
			c.JSON(200, gin.H{
				"main": "ok",
			})
		}
	})
	File_.POST("/CopyFile", func(c *gin.Context) {
		var CopyFile Struct.SandCopyFile
		Data, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(200, gin.H{
				"err": "解析参数错误！",
			})
			return
		}

		if err := json.Unmarshal(Data, &CopyFile); err != nil {
			c.JSON(200, gin.H{
				"err": "解析参数错误！",
			})
			return
		}
		fmt.Println(CopyFile)
		if err := File.CopyFile(CopyFile); err != nil {
			c.JSON(200, gin.H{
				"main": err,
			})
		} else {
			c.JSON(200, gin.H{
				"main": err,
			})
		}
	})
	r.Run(":8080")

}
