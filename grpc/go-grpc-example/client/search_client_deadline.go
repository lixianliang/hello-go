package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "go-grpc-example/proto"
)

const (
	PORT = "9001"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(4*time.Second)))
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(6*time.Second)))
	defer cancel()

	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial failed: %v", err)
	}

	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(ctx, &pb.SearchRequest{Request: "gRPC"})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Fatalf("client.Search err: DeadlineExceeded")
				return
			}
		}
		log.Fatalf("Search failed: %v", err)
	}

	log.Printf("res: %s", resp.Response)
}
