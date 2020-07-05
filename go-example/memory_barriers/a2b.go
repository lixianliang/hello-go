package main

import (
	"fmt"
	//"time"
	"sync"
)

var c chan int
var sum int
var wg sync.WaitGroup

func A() {
	defer wg.Done()
	sum += 1
	fmt.Println("A: ", sum)
	c <- 0
}

func B() {
	defer wg.Done()
	<-c
	sum += 1
	fmt.Println("B: ", sum)
}

func main() {
	c = make(chan int)
	// c2 := make(chan int)
	sum = 0
	//wg = sync.WaitGroup{}
	//go func() {
	for i := 0; i < 20; i += 2 {
		wg.Add(2)
		go A()
		go B()
	}
	wg.Wait()
	//c2 <- 0
	//}()
	//time.Sleep(time.Second * 1)
	//<-c2
	fmt.Println(sum)
}
