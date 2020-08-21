package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHanlders int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHanlders)
	}

	router.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})
	router.POST("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})
	router.POST("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "status")
	})

	router.Run(":8080")
}
