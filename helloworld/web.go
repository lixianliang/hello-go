package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"context"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-web"
	proto "helloworld/proto"
)

type Say struct{}

var (
	cl proto.GreeterService
)

func (s *Say) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func (s *Say) Hello(c *gin.Context) {
	log.Print("Received Say.Hello API request")

	name := c.Param("name")

	response, err := cl.Hello(context.TODO(), &proto.HelloRequest{
		Name: name,
	})

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, response)
}

func main() {
	// Create service 这里需要注意使用的web.NewService 而不是micro.NewService 后文会有解释
	service := web.NewService(
		web.Name("go.micro.api.greeter"),
	)

	service.Init()

	// setup Greeter Server Client
	//cl = proto.NewGreeterService("go.micro.srv.greeter", client.DefaultClient)
	cl = proto.NewGreeterService("greeter", client.DefaultClient)

	// Create RESTful handler (using Gin)
	say := new(Say)
	router := gin.Default()
	router.GET("/greeter", say.Anything)
	router.GET("/greeter/:name", say.Hello)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
