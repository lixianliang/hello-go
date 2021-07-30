package main

import (
	"fmt"
)

type Singleton struct {
	a int
}

var singleton *Singleton

func init() {
	singleton = &Singleton{a: 3}
}

func GetInstance() *Singleton {
	return singleton
}

func main() {
	fmt.Println(GetInstance())
}
