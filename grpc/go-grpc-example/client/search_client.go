package main

import (
	"context"
	//"io"
	"log"

	"google.golang.org/grpc"

	pb "go-grpc-example/proto"
)

const (
	PORT = "9001"
)

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
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
