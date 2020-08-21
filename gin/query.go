package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {
		// 区分大小写，如果没有对应的参数，default为Guest
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run(":8080")
}
