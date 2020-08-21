package main

import (
	"log"
	//"net/http"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	log.Fatal(autotls.Run(router, "example1.com", "example2.com"))
}
