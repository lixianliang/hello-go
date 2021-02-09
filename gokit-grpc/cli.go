package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	"gokit-grpc/pb"
	"gokit-grpc/service"
)

func main() {
	var (
		grpcAddr = flag.String("addr", ":9001", "gRPC address")
	)
	flag.Parse()

	ctx := context.Background()

	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		fmt.Println("gRPC dial err:", err)
	}
	defer conn.Close()

	svr := NewClient(conn)
	result, err := svr.Calculate(ctx, "Add", 10, 2)
	if err != nil {
		fmt.Println("calculate error", err.Error())

	}

	fmt.Println("result=", result)
}

func NewClient(conn *grpc.ClientConn) service.Service {
	var ep = grpctransport.NewClient(conn,
		"pb.ArithmeticService",
		"Calculate",
		service.EncodeGRPCArithmeticRequest,
		service.DecodeGRPCArithmeticResponse,
		pb.ArithmeticResponse{},
	).Endpoint()

	arithmeticEp := service.ArithmeticEndpoints{
		CalculateEndpoint: ep,
	}
	return arithmeticEp
}
