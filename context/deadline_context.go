package main

import (
	"context"
	"fmt"
	"time"
)

func dl2(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("dl2 :", n)
			n++
			time.Sleep(time.Second)
		}
	}
}

func dl1(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("dl1: ", n)
			n++
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	d := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	go dl1(ctx)
	go dl2(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("over", ctx.Err())
			return
		}
	}
}
