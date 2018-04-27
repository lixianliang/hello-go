package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"hello"
)

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TServerTransport
	var err error
	transport, err = thrift.NewTServerSocket(addr)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", transport)
	handler := NewHelloHandler()
	processor := hello.NewHelloProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Println("Starting the simple server... on ", addr)
	return server.Serve()
}

func main() {
	//transportFactory := thrift.NewTTransportFactory()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	runServer(transportFactory, protocolFactory, "localhost:9090")
}
