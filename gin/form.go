package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	// curl -XPOST -H 'Content-Type: application/x-www-form-urlencoded' -d'nick=manu&message=this_is_great' localhost:8080/form_post
	// curl -XPOST -d'nick=manu&message=this_is_great' localhost:8080/form_post

	router.Run(":8080")
}
