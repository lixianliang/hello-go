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
	router.Static("/", "./pulibc")
	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		files := form.File["files"]

		for _, file := range files {
			log.Println(file.Filename)
			filename := filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}

		c.String(http.StatusOK, fmt.Sprintf("upload success %d files with fileds name=%s email=%s", len(files), name, email))
	})
	// curl -v -X POST 'http://localhost:8080/upload' -F 'files=@/home/lxl/huiyi.md' -F 'files=@/home/lxl/LICENSE' -H "Content-Type: multipart/form-data" -F 'name=lxl;email=826028251@qq.com' -H "Content-Type: multipart/form-data"
	// curl -F表示multipart formpost -d表示encode url请求方式，两者只能选一种
	// curl的files和代码中form.File["files"]对应

	router.Run(":8080")
}
