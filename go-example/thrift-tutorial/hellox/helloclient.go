package main

import (
	"fmt"
	"hello"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func handleClient(client *hello.HelloClient) error {
	str, err := client.HelloString("123")
	if err != nil {
		fmt.Printf("request failed: %s\n", err.Error())
		return nil
	} else {
		fmt.Printf("HelloString():%s\n", str)
	}

	n, err := client.HelloInt(1)
	if err != nil {
		fmt.Printf("request failed: %s\n", err.Error())
	} else {
		fmt.Printf("HelloInt(): %d\n", n)
	}

	b, err := client.HelloBoolean(true)
	if err != nil {
		fmt.Printf("request failed: %s\n", err.Error())
	} else {
		fmt.Printf("HelloBoolean(): %t\n", b)
	}

	err = client.HelloVoid()
	if err != nil {
		fmt.Printf("request failed: %s\n", err.Error())
	} else {
		fmt.Printf("Hellovoid()\n")
	}

	str, err = client.HelloNull()
	if err != nil {
		fmt.Printf("request failed: %s\n", err.Error())
	} else {
		fmt.Printf("HelloNull(): %s\n", str)
	}

	return nil
}

func runClient(transportFactory thrift.TTransportFactory,
	protocolFactory thrift.TProtocolFactory,
	addr string) error {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error NewTSocket failed:", err)
		return err
	}
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println("Error GetTransport failed:", err)
		return err
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		fmt.Println("Error Open failed:", err)
		return err
	}
	return handleClient(hello.NewHelloClientFactory(transport, protocolFactory))
}

func main() {
	//transportFactory := thrift.NewTTransportFactory()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	runClient(transportFactory, protocolFactory, "localhost:9090")
}
