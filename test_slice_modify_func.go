package main

import (
	"fmt"
	"sync"
	"time"
)

func xappend(s []int, wg *sync.WaitGroup) {
	defer wg.Done()
	s = append(s, 3)
	//s = append(s, 4)
	//fmt.Println(s)
	s = append(s, 4)
	fmt.Println(s)
	time.Sleep(2 * time.Second)

	fmt.Println(s)
}

func xmod(s []int) {
	s[1] = 3
}

func main() {
	var wg sync.WaitGroup
	//s := []int{1, 2, 2}
	s := make([]int, 3, 10)
	s[0] = 1
	s[1] = 2
	s[2] = 2
	s1 := s[:2]
	//xmod(s)
	wg.Add(1)
	xappend(s1, &wg)
	// go xappend(s1, &wg) // 通过race检测，对slice里面共享数组部分会有读写竞争的情况，但串行的方式运行时没有的
	time.Sleep(1 * time.Second)
	s[2] = 5
	wg.Wait()
	fmt.Println(s)
	fmt.Println(s1)
}
