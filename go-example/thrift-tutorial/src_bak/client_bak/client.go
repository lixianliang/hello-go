package main

import (
    "crypto/tls"
    "fmt"
    "git.apache.org/thrift.git/lib/go/thrift"
    "tutorial"
)

func handleClient(client *tutorial.CalculatorClient) (err error) {
    client.Ping()
    fmt.Println("ping()")

    sum, _ := client.Add(1, 1)
    fmt.Print("1+1=", sum, "\n")

    work := tutorial.NewWork()
    work.Op = tutorial.Operation_DIVIDE
    work.Num1 = 1
    work.Num2 = 0
    quotient, err := client.Calculate(1, work)
    if err != nil {
        switch v := err.(type) {
        case *tutorial.InvalidOperation:
            fmt.Println("Invalid operation:", err)
        default:
            fmt.Println("Error during operation:", err)
        }
        return err
    } else {
        fmt.Println("Whoa we can divide by 0 with new value:", quotient)
    }

    work.Op = tutorial.Operation_SUBTRACT
    work.Num1 = 15
    work.Num2 = 10
    diff, err := client.Calculate(1, work)
    if err != nil {
        switch v := err.(type) {
        case *tutorial.InvalidOperation:
            fmt.Println("Invalid operation:", v)
        default:
            fmt.Println("Error during operation:", err)
        }
    } else {
        fmt.Print("15-10=", diff, "\n")
    }

    log, err := client.GetStruct(1)
    if err != nil {
        fmt.Println("Unable to get struct:", err)
    } else {
        fmt.Println("Check log:", log.Value)
    }
    return err
}

func runClient(transportFacetory thrift.TTransportFactory, protocolFactory thrift.TProtocolFacetory, addr string, secure bool) error {
    var transport thrift.TTransport
    var err error
    if secure {
        cfg := new(tls.Config)
        cfg.InsecureSkipVerify = true
        transport, err = thrift.NewTSSLSocket(addr, cfg)
    } else {
        transport, err = thrift.NewTSocket(addr)
    }
    if err != nil {
        fmt.Println("Error opening socket:", err)
        return err
    }
    transport, err = transportFacetory.GetTransport(transport)
    if err != nil {
        return err
    }
    defer transport.Close()
    if err := transport.Open(); err != nil {
        return err
    }
    return handleClient(tutorial.NewCalculatorClientFactory(transport, protocolFactory))
}
