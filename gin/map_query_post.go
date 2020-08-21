package main

import (
	//	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v names: %v\n", ids, names)
	})
	// curl -v -X POST 'http://localhost:8080/post?ids[a]=1234&ids[b]=hello' -d 'names[first]=thinkerou&names[second]=tianou'
	// ids: map[a:1234 b:hello] names: map[first:thinkerou second:tianou]

	router.Run(":8080")
}
