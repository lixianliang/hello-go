package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
	// 绑定uuid类型
}

func main() {
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})

	route.Run(":8080")
	// curl -v -X GET 'http://localhost:8080/thinkerou/not-uuid'
	//  curl -v -X GET 'http://localhost:8080/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3'
}
