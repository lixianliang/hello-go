package main

import (
	"context"
	"sync"

	//"github.com/pkg/errors"
	"errors"
)

func rpc(ctx context.Context, url string) error {
	result := make(chan int)
	err := make(chan error)

	go func() {
		isSuccess := true
		if isSuccess {
			result <- 1
		} else {
			err <- errors.New("some error happen")
		}
	}()

	select {
	case <-result:
		// 本rpc成功
		return nil
	case e := <-err:
		// 本rpc失败
		return e
	case <-ctx.Done():
		// 其它rpc失败
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	err := rpc(ctx, "http://rpc_url1")
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := rpc(ctx, "http://rpc_url2")
		if err != nil {
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := rpc(ctx, "http://rpc_url3")
		if err != nil {
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := rpc(ctx, "http://rpc_url4")
		if err != nil {
			cancel()
		}
	}()

	wg.Wait()
	//	defer cancel()
}
