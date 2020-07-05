package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.RWMutex
	arr := []int{1, 2, 3}
	go func() {
		fmt.Println("try to lock writing operation")
		mutex.Lock()
		fmt.Println("writing operation is locked")

		arr = append(arr, 4)

		fmt.Println("try to unlock writing operation")
		mutex.Unlock()
		fmt.Println("writing operation is unlock")
	}()

	go func() {
		fmt.Println("try to lock reading operation")
		mutex.RLock()
		fmt.Println("thr reading operation is locked")

		fmt.Println("the len of arr is : ", len(arr))

		fmt.Println("try to unlock reading operation")
		mutex.RUnlock()
		fmt.Println("the reading operation is unlock")
	}()

	time.Sleep(time.Second * 2)
}
