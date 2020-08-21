package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123456"},
	"austin": gin.H{"email": "austin@bar.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@bar.com", "phone": "555"},
}

func main() {
	router := gin.Default()

	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello",
		"manu":   "4321",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET"})
		}
	})
	// curl  -v -X GET 'http://192.168.199.109:8080/admin/secrets' --user foo:bar
	// Authorization: Basic Zm9vOmJhcg==

	// failed: HTTP/1.1 401 Unauthorized
	// Www-Authenticate: Basic realm="Authorization Required"

	router.Run(":8080")
}
