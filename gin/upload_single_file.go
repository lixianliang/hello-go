package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("%s uploaded", file.Filename))
	})
	// curl -v -X POST 'http://localhost:8080/upload' -F 'file=@/home/lxl/huiyi.md'
	// curl 会自动转换为Content-Type: multipart/form-data; boundary=------------------------5a40fba00cc3cfaa
	// Content-Type: multipart/form-data在程序中还是要加的

	router.Run(":8080")
}
