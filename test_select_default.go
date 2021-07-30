package main

import "fmt"

func main() {
	ch := make(chan string)
	select {
	case <-ch:
		fmt.Println("ch")
	default:
		fmt.Println("default case executed")
	}
}
