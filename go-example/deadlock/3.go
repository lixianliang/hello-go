package main

import (
    "fmt"
)

func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2

    for v := range ch {
        fmt.Println(v)

    }
    // 也会deadlock， range不等到信道关闭是不会结束读取的; 如果缓冲信道没有数据，那么range就会阻塞当前goroutine，所以会死锁
    // 解决方式
        // 1. if len(ch) <= 0 break
        // 1. 显示关闭信道 close(ch)
}
