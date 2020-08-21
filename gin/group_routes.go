package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Run() {
	v1 := router.Group("/v1")
	addUserRoutes(v1)
	addPingRoutes(v1)

	v2 := router.Group("/v2")
	addPingRoutes(v2)

	router.Run(":8080")
}

func main() {
	Run()
	/*router := gin.Deafult()

	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")*/
}

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users")
	})

	users.GET("/comments", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users comments")
	})

	users.GET("/pictures", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users pictures")
	})
}

func addPingRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")

	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}
