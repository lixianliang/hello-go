package main

import (
	"context"
	"log"
	"net"
	"time"
	//"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "go-grpc-example/proto"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	for i := 0; i < 5; i++ {
		// ctx.Err() == context.Canceled 添加sleep进行模拟场景
		if ctx.Err() == context.Canceled {
			log.Println("xxx: canceled")
			return nil, status.Errorf(codes.Canceled, "searchService.Search canceled")
		}
		time.Sleep(1 * time.Second)
	}

	return &pb.SearchResponse{Response: r.Request + " Server"}, nil
}

const PORT = "9001"

func main() {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &SearchService{})
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
