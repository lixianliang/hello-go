package main

import (
	"log"
	//	"net/http"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(router, &m))
}
