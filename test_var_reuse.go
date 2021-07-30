package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	go匿名函数复用上层for循环的变量，一般运行的块都是使用最后的变量，因为内存地址都是一样的，如果通过sleep的方式，可以看到时批量一致的
*/

func main() {
	wg := sync.WaitGroup{}
	//ch := make(chan int)

	for i := 1; i < 2000; i++ {
		wg.Add(1)
		if i%100 == 0 {
			time.Sleep(100 * time.Millisecond)
		}
		go func() {
			fmt.Printf("%d ", i)
			defer wg.Done()
		}()
	}
	wg.Wait()

	//close(ch)

	//<-ch
}
