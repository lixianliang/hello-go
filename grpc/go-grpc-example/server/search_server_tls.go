package main

import (
	"context"
	"log"
	"net"
	//"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	pb "go-grpc-example/proto"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "searchService.Search canceled")
	}
	return &pb.SearchResponse{Response: r.Request + " Server"}, nil
}

const PORT = "9001"

func main() {
	// credentials.NewServerTLSFromFile：根据服务端输入的证书文件和密钥构造 TLS 凭证
	c, err := credentials.NewServerTLSFromFile("conf/server.pem", "conf/server.key")
	if err != nil {
		log.Fatalf("NewServerTLSFromFile err: %v", err)
	}

	// grpc.Creds()：返回一个 ServerOption，用于设置服务器连接的凭据。用于 grpc.NewServer(opt ...ServerOption) 为 gRPC Server 设置连接选项
	server := grpc.NewServer(grpc.Creds(c))
	pb.RegisterSearchServiceServer(server, &SearchService{})
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
