package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//  ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	//  defer cancel()
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Millisecond)

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslep")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
