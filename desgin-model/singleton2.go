package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	a int
}

var (
	lazySingleton *Singleton
	once          sync.Once
)

func GetInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{a: 3}
		})
	}
	return lazySingleton
}

func main() {
	fmt.Println(GetInstance())
}
