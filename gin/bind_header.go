package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type testHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
	// 绑定header
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		h := testHeader{}

		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(200, err)
		}

		fmt.Printf("%#v\n", h)
		c.JSON(200, gin.H{"Rate": h.Rate, "Domain": h.Domain})
	})

	r.Run()
	// curl -v -X GET 'http://localhost:8080' -H "rate:300" -H "domain:music"
}
