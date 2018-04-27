package main

import (
	"crypto/tls"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"tutorial"
)

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	var transport thrift.TServerTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		if cert, err := tls.LoadX509KeyPair("server.crt", "server.key"); err == nil {
			cfg.Certificates = append(cfg.Certificates, cert)
		} else {
			return err
		}
		transport, err = thrift.NewTSSLServerSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTServerSocket(addr)
	}

	if err != nil {
		return err
	}
	fmt.Println("%T\n", transport)
	handler := NewCalculatorHandler()
	processor := tutorial.NewCalculatorProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on", addr)
	return server.Serve()
}

func main() {
	transportFactory := thrift.NewTBufferedTransportFactory(8800)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	err := runServer(transportFactory, protocolFactory, ":8800", false)
	//err := runServer(transportFactory, protocolFactory, ":8800", true)
	if err != nil {
		fmt.Printf("runServer failed: %s\n", err.Error())
	}
}
