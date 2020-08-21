package main

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileForm struct {
	Name   string                `form:"name" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/profile", func(c *gin.Context) {
		var form ProfileForm

		if err := c.ShouldBind(&form); err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			return
		}

		err := c.SaveUploadedFile(form.Avatar, form.Avatar.Filename)
		if err != nil {
			c.String(http.StatusInternalServerError, "unkonow error")
			return
		}

		c.String(http.StatusOK, "or")
	})

	router.Run(":8080")
	//  curl -v -X POST --form name=user --form "avatar=@./xnpijiws0s0khgic0e8bh4q.jpg" 'http://localhost:8080/profile'
}
