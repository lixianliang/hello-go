package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/loacl/file", func(c *gin.Context) {
		c.File("./serving_file.go")
	})

	//var fs http.FileSystem 后续再研究，新的接口
	var fs http.FileSystem = router.GET("/fs/file", func(c *gin.Context) {
		c.FileFromFS("fs/file.go", fs)
	})
}
