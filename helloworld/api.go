package main

import (
	"encoding/json"
	"log"
	"strings"

	// 引用上面生成的proto文件
	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/v2"
	proto "helloworld/proto"

	"context"
)

type Say struct {
	Client proto.GreeterService
}

func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Say.Hello API request")

	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.greeter", "Name cannot be blank")
	}
	log.Println("name: %s", name.Values)

	response, err := s.Client.Hello(ctx, &proto.HelloRequest{
		Name: strings.Join(name.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Greeting,
	})
	rsp.Body = string(b)

	return nil
}

func main() {
	// new一个微服务出来 资源类型设置为api
	service := micro.NewService(
		micro.Name("go.micro.api.greeter"),
	)

	// 可选 解析命令行
	service.Init()

	// 注册handler
	service.Server().Handle(
		service.Server().NewHandler(
			//&Say{Client: proto.NewGreeterService("go.micro.srv.greeter", service.Client())},
			&Say{Client: proto.NewGreeterService("greeter", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
