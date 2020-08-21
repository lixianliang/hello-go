package main

import (
	"fmt"
	//"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// curl -v -X POST 'http://localhost:8080/post?id=123&page=1' -d 'name=manu&message=this_is_great'
	// post默认为Content-Type为application/x-www-form-urlencoded
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id:%s page:%s name:%s message:%s\n", id, page, name, message)
	})

	router.Run(":8080")
}
