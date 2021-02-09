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

	pb "helloworld/proto"
)

func main() {
	var (
		gRPCAddr = flag.String("addr", ":32845", "gRPC address")
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
	cli := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	resp, err := cli.Hello(ctx, &pb.HelloRequest{Name: "John"})
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(resp)
	}
}
