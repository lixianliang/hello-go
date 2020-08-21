package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())

	// CustomRecovery为master的功能，在v1.6.3最新版本还不包含
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error:%s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	r.GET("/panic", func(c *gin.Context) {
		panic("foo")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ohai")
	})

	r.Run(":8080")
}
