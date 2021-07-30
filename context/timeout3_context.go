package main

import (
	"context"
	"fmt"
	"time"
)

func to1(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("to1 is over")
			return
		default:
			fmt.Println("to1: ", n)
			n++
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	go to1(ctx)
	n := 1
	for {
		select {
		case <-time.Tick(2 * time.Second):
			if n == 9 {
				return
			}
			fmt.Println("number: ", n)
			n++
		}
	}
}
