package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var sum uint32 = 0
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//sum +=1
			atomic.AddUint32(&sum, 1)
		}()
	}
	wg.Wait()

	fmt.Println(sum)
}
