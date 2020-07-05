package main

import (
	"fmt"
	"time"
)

type T struct {
	msg string
}

var g *T

func setup() {
	t := new(T)
	t.msg = "hello, world"

	g = t
}

// 一直在循环run?
func main() {
	go setup()
	for {
		if g != nil { // g一直判断不了,多个G的同步不安全
			break
		}
		// 加上sleep就可以了
		time.Sleep(time.Second * 1)
	}
	fmt.Println(g.msg)
}
