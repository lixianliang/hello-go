package main

import (
	"context"
	"fmt"
	"time"
)

func stop1() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出， 停止")
				return
			default:
				fmt.Println("g监控中")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以，通知监控停止")
	stop <- true
	time.Sleep(5 * time.Second)
}

func stop2() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出， 停止")
				return
			default:
				fmt.Println("g监控中")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以，通知监控停止")
	cancel()
	time.Sleep(5 * time.Second)
}

func main() {
	//stop1()
	stop2()
}
