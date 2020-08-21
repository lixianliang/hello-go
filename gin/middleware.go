package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// create a router without any middleware by default
	// gin.Default	// default with the logger and recovery middleware already attached
	r := gin.New()

	// Logger中间件将日志写入gin.DefaultWriter，即使你将GIN MODE设置为release
	r.Use(gin.Logger())
	// Recovery中间件会recover任何panic，如果有panic的华，会写入500
	r.Use(gin.Recovery())

	// 可以为每个路由添加任意数量的中间件
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	authorized := r.Group("/")
	// 路由组中间件，authorized自定义了AuthRequired中间件
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/admin", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// 嵌套路由
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	r.Run(":8080")
}

func analyticsEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, "analytics")
}

func readEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, "read")
}

func submitEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, "submit")
}

func loginEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, "login")
}

func benchEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, "bench")
}

func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before bench logger middleware")
		c.Set("request", "client_request")
		c.Next()
		fmt.Println("after bench logger middleware")
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before auth middleware")
		//c.Set("request", "client_request")
		c.Next()
		fmt.Println("after auth middleware")
	}
}
