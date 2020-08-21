package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name       string    `form:"name"`
	Address    string    `form:"address"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", startPage)

	route.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person

	// GET only form binding engine query
	// POST content-type json xml使用form(form-data)
	if c.ShouldBind(&person) == nil {
		log.Printf("%s %s %v %v %v", person.Name, person.Address, person.Birthday, person.CreateTime, person.UnixTime)
	}
	c.String(200, "success")
}
