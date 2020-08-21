package main

import (
	"fmt"
	_ "net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(FirstMiddleware(), SecondMiddleware())
	router.GET("/", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}

func FirstMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("first middleware before next()")
		isAbort := c.Query("isAbort")
		bAbort, err := strconv.ParseBool(isAbort)
		if err != nil {
			fmt.Printf("is abort value err, value %s\n", isAbort)
			c.Next() // (2)
			// 这里回执行到第2个middelware
		}
		if bAbort {
			fmt.Println("first middleware abort") //(3)
			c.Abort()
			//c.AbortWithStatusJSON(http.StatusOK, "abort is true")
			// abort最好是也返回对应的错误码
			return
		} else {
			fmt.Println("first middleware doesnot abort") //(4)
			return
		}

		fmt.Println("first middleware after next()")
		// 这一行执行不到
	}
}

func SecondMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("current inside of second middleware")
	}
}
