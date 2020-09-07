package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
	// 引用上面生成的proto文件
	proto "helloworld/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	log.Printf("msg: %s\n", rsp.Greeting)
	return nil
}

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Printf("%v server request: %s\n", time.Now(), req.Endpoint())
		return fn(ctx, req, rsp)
	}
}

func main() {
	// new一个微服务出来
	service := micro.NewService(
		micro.Name("greeter"),
		//micro.Name("go.micro.srv.greeter"),
		micro.Version("latest"),
		micro.WrapHandler(logWrapper),
	)

	// 可选 解析命令行
	service.Init()

	// 注册 handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
