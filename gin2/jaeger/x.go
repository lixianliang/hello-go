package main

import "fmt"
import "github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"

func main() {
	fmt.Println("vim-go")

	//	var tracer opentracing.Tracer = ...
	//client
	//... // other options
	conn, err := grpc.Dial(
		address,
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(tracer)),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(tracer)))

	// ... // other options
	// server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(tracer)),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(tracer)))
}
