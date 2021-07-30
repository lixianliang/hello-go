package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ch <- 1
		ch <- 2
		defer wg.Done()
	}()
	wg.Wait()
	close(ch)
	go func() {
		for x := range ch {
			fmt.Println(x)
		}
	}()
	time.Sleep(10 * time.Second)
	// 可以从close的ch里面range出来数据
	// 如果不关闭ch，可以range出来但会报deadlock， 因为range 阻塞到ch上面了
}
