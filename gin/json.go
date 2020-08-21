package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.SecureJsonPrefix(")]}',\n")

	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		c.SecureJSON(http.StatusOK, names)
	})

	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		c.JSONP(http.StatusOK, data)
	})
	// call({"foo":"bar"});

	r.GET("/AsciiJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "go语言",
			"tag":  "<br>",
		}

		c.AsciiJSON(http.StatusOK, data)
	})
	// {"lang":"go\u8bed\u8a00","tag":"\u003cbr\u003e"}

	r.GET("/JSON", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>hello world</b>",
		})
	})
	// {"html":"\u003cb\u003ehello world\u003c/b\u003e"}l

	r.GET("/PureJSON", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>hello world</b>",
		})
	})
	// {"html":"<b>hello world</b>"}

	r.Run(":8080")
}
