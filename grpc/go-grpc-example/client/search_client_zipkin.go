package main

import (
	"context"
	//"io"
	"log"

	//"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"

	pb "go-grpc-example/proto"
)

const (
	PORT = "9001"

	SERVICE_NAME              = "search_server_zipkin"
	ZIPKIN_HTTP_ENDPOINT      = "http://127.0.0.1:9411/api/v2/spans"
	ZIPKIN_RECORDER_HOST_PORT = "127.0.0.1:9000"
)

func main() {
	reporter := zipkinhttp.NewReporter(ZIPKIN_HTTP_ENDPOINT)
	defer reporter.Close()
	endpoint, err := zipkin.NewEndpoint(SERVICE_NAME, ZIPKIN_RECORDER_HOST_PORT)
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}
	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
	}
	tracer := zipkinot.Wrap(nativeTracer)

	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure(), grpc.WithUnaryInterceptor(
		otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads()),
	))
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
