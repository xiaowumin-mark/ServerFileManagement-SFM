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
	Setting.GET("/disk", func(c *gin.Context) {
		decodedPerson, err := ToJson.ConfJson(chief.Config())
		if err != nil {
			fmt.Println("Error decoding person:", err)
			return
		}
		c.JSON(200, decodedPerson)
	})
	r.Run(":8080")

}
