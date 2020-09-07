package main

import (
	"context"
	"flag"
	//	"fmt"
	"log"
	//"strconv"
	"time"

	//	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	//	"github.com/lixianliang/hello-go/gokit-playground/lorem-grpc"
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

	log.Println("client connect ok")
	cli := pb.NewLoremClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	resp, err := cli.Lorem(ctx, &pb.LoremRequest{RequestType: "Sentence", Min: 4, Max: 6})
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(resp)
	}
}
