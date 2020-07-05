package main

import (
	"fmt"
)

var c chan int
var a string

func f() {
	a = "hello world"
	<-c
}

func main() {
	//c = make(chan int, 1)
	c = make(chan int)
	go f()
	c <- 0
	fmt.Println(a)
}
