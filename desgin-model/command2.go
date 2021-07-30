package main

import (
	"fmt"
	"time"
)

type Command func() error

func StartCommandFunc() Command {
	return func() error {
		fmt.Println("game start")
		return nil
	}
}

func ArchiveCommandFunc() Command {
	return func() error {
		fmt.Println("game archive")
		return nil
	}
}

func main() {
	// 用于测试，模拟来自客户端事件
	eventChan := make(chan string)
	go func() {
		events := []string{"start", "archive", "start", "archive", "start", "start"}
		for _, e := range events {
			eventChan <- e
		}
	}()
	defer close(eventChan)

	// 使用命名队列缓存命令
	commands := make(chan Command, 1000)
	defer close(commands)

	go func() {
		for {
			event, ok := <-eventChan
			if !ok {
				return
			}

			var command Command
			switch event {
			case "start":
				command = StartCommandFunc()
			case "archive":
				command = ArchiveCommandFunc()
			}

			commands <- command
		}
	}()

	for {
		select {
		case c := <-commands:
			c()
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1s")
			return
		}
	}
}
