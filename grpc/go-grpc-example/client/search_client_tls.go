package main

import (
	"context"
	//"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "go-grpc-example/proto"
)

const (
	PORT = "9001"
)

func main() {
	c, err := credentials.NewClientTLSFromFile("conf/server.pem", "go-grpc-example")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	}

	// NameConstraintsWithoutSANs
	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial failed: %v", err)
	}

	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)

	resp, err := client.Search(context.Background(), &pb.SearchRequest{Request: "abc"})
	if err != nil {
		log.Fatalf("Search failed: %v", err)
	}

	log.Printf("res: %s", resp.Response)
}
