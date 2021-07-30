package main

import (
	"fmt"
)

/*
	可以对slice 为nil的切片进行range操作
*/

func main() {
	var x []string
	for _, a := range x {
		fmt.Println(a)
	}
}
