package main

import (
	"log"
	//"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/long_async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("done in path " + cCp.Request.URL.Path)
		}()
	})

	router.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("done in path " + c.Request.URL.Path)
	})

	router.Run(":8080")
}
