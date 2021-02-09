package main

import (
	"context"
	"fmt"

	// 引用上面生成的proto文件
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/selector"
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

type firstNodeSelector struct {
	opts selector.Options
}

// 初始化选择器
func (n *firstNodeSelector) Init(opts ...selector.Option) error {
	for _, o := range opts {
		o(&n.opts)
	}
	return nil
}

// select 返回options
func (n *firstNodeSelector) Options() selector.Options {
	return n.opts
}

// 对从服务发现取出来的服务实例进行选择
func (n *firstNodeSelector) Select(service string, opts ...selector.SelectOtption) (selector.Next, error) {
	services, err := n.opts.Resgistry.GetService(service)
	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return nil, selector.ErrNotFound
	}

	var sopts selector.SelectOtptions
	for _, opt := range opts {
		opt(&sopts)
	}

	for _, filter := range sopts.Filter {
		services = filter(services)
	}
	if len(services) == 0 {
		return nil, selector.ErrNotFound
	}
	if len(services[0].Nodes) == 0 {
		return nil, selector.ErrNotFound
	}

	return func() (*registry.Node, error) {
		return services[0].Nodes[0], nil
	}, nil
}

func (n *firstNodeSelector) Mark(services string, node *registry.Node, err error) {
	return
}

func (n *firstNodeSelector) Reset(services string) {
	return
}

func (n *firstNodeSelector) Close() error {
	return nil
}

// 返回selector的命名
func (n *firstNodeSelector) String() string {
	return "first"
}

func main() {
	// new一个服务
	// service := micro.NewService()
	service := micro.NewService(micro.Name("greeter.client"), micro.WrapClient(logWrap))

	// 解析命令行flag
	service.Init()

	// 使用proto创建一个客户端
	xcli := client.NewClient(client.Selector(FirstNodeSelector))
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
