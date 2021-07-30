package main

import "fmt"
import "context"

type traceKey struct{}

var k = traceKey{}

func main() {
	//	fmt.Println("vim-go")
	ctx := context.Background()
	x := ctx.Done()
	fmt.Printf("%v \n", x)
	fmt.Printf("%v %s\n", k, k)
}
