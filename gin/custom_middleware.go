package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("example", "12345")

		c.Next()

		lantency := time.Since(t)
		log.Println(lantency)

		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	router := gin.New()
	router.Use(Logger())

	router.GET("/test1", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		//log.Println("abcd")
		log.Println(example)
	})

	router.Run(":8080")
}
