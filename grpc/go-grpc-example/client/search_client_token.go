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

type Auth struct {
	AppKey    string
	AppSecret string
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"app_key": a.AppKey, "app_secret": a.AppSecret}, nil
}

func (a *Auth) RequireTransportSecurity() bool {
	return false
}

func main() {
	auth := Auth{
		AppKey:    "eddycjy",
		AppSecret: "20181005",
	}
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
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
