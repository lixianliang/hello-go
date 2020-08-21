package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/some", getting)
	router.POST("/some", getting)
	router.PUT("/some", getting)
	router.DELETE("/some", getting)
	router.PATCH("/some", getting)
	router.HEAD("/some", getting)
	router.OPTIONS("/some", getting)

	router.Run()
}

func getting(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodOptions:
		// body,_ := ioutil.ReadAll(C.Request.Body)
		c.JSON(http.StatusOK, gin.H{
			"messge": "hello",
		})
	case http.MethodHead:
		//	c.JSON(http.StatusOK, "") not need body
		//c.JSON(http.StatusOK, nil)
		// http head only response header
		fmt.Printf("head http \r\n")
		c.String(http.StatusOK, "")
		//c.String(http.StatusOK, "ok")
	}
}
