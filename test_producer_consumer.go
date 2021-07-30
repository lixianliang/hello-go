package main

import (
	"fmt"
	"sync"
)

func consumer(wg *sync.WaitGroup, ch <-chan int) {
	for {
		select {
		case x, ok := <-ch:
			if ok {
				fmt.Println(x, ok)
				//wg.Done()
			} else {
				// not ok
				fmt.Println("xx")
			}
		}
	}
}

func main() {
	//ch = make(chan int, 10)
	ch := make(chan int, 10)
	//ch = make(chan int)
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		x := i
		go func(y int) {
			ch <- y
			defer wg.Done()
		}(x)
	}

	wg.Wait() // 保证生产的数据，能够消费完
	close(ch) // 后关闭
	consumer(&wg, ch)

	// 可以通过done channal来通知和关闭所有的g, 这里的g没有主动关闭，通过程序退出关闭的
	// 关闭一个ch，所有监听的g都可以获取到这个通知

	fmt.Printf("Hello World!\n")
}
