package main

import (
	"sync"
)

var once sync.Once
var a string
var done bool

func setup() {
	a = "hello world"
	done = true
}

func doprint() {
	if done {
		once.Do(setup)
	}
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}

func main() {
	twoprint()
	//select {}
}
