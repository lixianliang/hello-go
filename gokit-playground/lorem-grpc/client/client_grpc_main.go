package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	"github.com/lixianliang/hello-go/gokit-playground/lorem-grpc"
	//grpcClient "github.com/lixianliang/hello-go/gokit-playground/lorem-grpc/client"
	"github.com/ru-rocker/gokit-playground/lorem-grpc/pb"
)

func main() {
	var (
		gRPCAddr = flag.String("addr", ":8081", "gRPC address")
	)
	flag.Parse()
	ctx := context.Background()
	conn, err := grpc.Dial(*gRPCAddr, grpc.WithInsecure(),
		grpc.WithTimeout(1*time.Second))

	if err != nil {
		log.Fatalln("gRPC dial: ", err)
	}
	defer conn.Close()

	loremService := New(conn)
	args := flag.Args()
	var cmd string
	cmd, args = pop(args)

	switch cmd {
	case "lorem":
		var requestType, minStr, maxStr string

		requestType, args = pop(args)
		minStr, args = pop(args)
		maxStr, args = pop(args)

		min, _ := strconv.Atoi(minStr)
		max, _ := strconv.Atoi(maxStr)
		lorem(ctx, loremService, requestType, min, max)
	default:
		log.Fatalln("unknown command", cmd)
	}
}

// parse command line args one by one
func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}

// call lorem service
func lorem(ctx context.Context, service lorem_grpc.Service, requestType string, min, max int) {
	log.Printf("lorem %s %d %d", requestType, min, max)
	mesg, err := service.Lorem(ctx, requestType, min, max)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(mesg)
}

// 这种grpc请求方式有问题
func New(conn *grpc.ClientConn) lorem_grpc.Service {
	// /pb.Lorem/Lorem
	// 都不行
	// 出现Unimplemented desc = unknown service Lorem错误
	var loremEndpoint = grpctransport.NewClient(
		//conn, "pb.Lorem", "Lorem",
		conn, "Lorem", "Lorem",
		lorem_grpc.EncodeGRPCLoremRequest,
		lorem_grpc.DecodeGRPCLoremResponse,
		pb.LoremResponse{},
	).Endpoint()

	return lorem_grpc.Endpoints{
		LoremEndpoint: loremEndpoint,
	}
}
