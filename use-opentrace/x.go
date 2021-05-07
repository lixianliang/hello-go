package main

import (
	"fmt"
	"net"
)

func main() {
	host := "www.baidu.com"
	//host := "tracer-agent.istio-system:6831"
	addrs, err := net.LookupHost(host)
	if err != nil {
		panic(fmt.Sprintf("aaa: %v\n", err))
	}

	fmt.Println(addrs)
}
