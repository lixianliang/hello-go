package main

import (
	"context"
	"fmt"

	// 引用上面生成的proto文件
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	proto "helloworld/proto"
)

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Printf("wrapper client request to service: %s method %s\n", req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

// 实现client.Wrapper充当日志包装器
func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}

func main() {
	// new一个服务
	// service := micro.NewService()
	service := micro.NewService(micro.Name("greeter.client"), micro.WrapClient(logWrap))

	// 解析命令行flag
	service.Init()

	// 使用proto创建一个客户端
	cl := proto.NewGreeterService("greeter", service.Client())

	// 发出请求
	rsp, err := cl.Hello(context.Background(), &proto.HelloRequest{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}
