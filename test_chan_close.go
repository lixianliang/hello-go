package main

import "fmt"

func main() {
	// 创建一个整型的通道
	ch := make(chan int)
	// 关闭通道
	close(ch)
	// 打印通道的指针, 容量和长度
	fmt.Printf("ptr:%p cap:%d len:%d\n", ch, cap(ch), len(ch))
	// 给关闭的通道发送数据
	ch <- 1
}
