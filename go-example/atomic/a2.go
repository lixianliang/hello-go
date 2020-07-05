package main

import (
	"fmt"
	"sync/atomic"
)

var value int32

func main() {
	fmt.Println("======old value======")
	fmt.Println(value)
	fmt.Println("======cas value======")
	addValue(3)
	fmt.Println(value)
}

func addValue(delta int32) {
	for {
		// 在读取value的操作中，其它对此值的读写操作都是可以被同时进行，那么这个操作可能会读取到一个被修改了一半数据(如果是在64位机器对齐的情况下指针、整形都是原子行的？)
		// 使用Load
		// v := value
		v := atomic.LoadInt32(&value)
		if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
			break
		}
		// 操作失败的原因可能会是value的old不与v相等了
		// cas操作虽然不会让某个goroutine阻塞在某条语句上，但可能会使用流程执行暂停一下，不过时间极其短暂
	}

	// fmt.Println()
	// atomic.StoreInt32(&value, 10)
	// 在原子的存储某个值的过程中，任何cpu都不会进行针对同一个值读或者写操作
	// 原子的值存储操作总会成功，因为它并不关心被操作值的旧值是什么和cas操作有着明显的区别
}
