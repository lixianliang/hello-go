package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	valCtx := context.WithValue(ctx, key, "监控1")
	go watch(valCtx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以， 通知监控停止")
	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), "监控退出，停止")
			return
		default:
			fmt.Println(ctx.Value(key), "g监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
