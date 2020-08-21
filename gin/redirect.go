package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://google.com/")
	})

	router.GET("/test-foo", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/foo") // 浏览器会发起两次请求 uri最终显示为foo
	})
	router.GET("/foo", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "foo"})
	})

	router.GET("test-test2", func(c *gin.Context) {
		c.Request.URL.Path = "/test2" // 浏览器只有一次请求，但gin会先显示test2路由请求，然后在显示test-test2路由，浏览器只请求一次，且uri显示为test-test2
		router.HandleContext(c)
	})
	router.GET("test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	router.Run(":8080")
}
