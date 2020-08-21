package main

import (
	"fmt"
	//	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	//  curl -v 'http://192.168.199.109:8080/cookie'
	// 响应的header含有 Set-Cookie: gin_cookie=test; Path=/; Domain=localhost; Max-Age=3600; HttpOnly
	router.Run(":8080")
}
