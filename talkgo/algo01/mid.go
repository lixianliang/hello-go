package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 4, 7, 11, 15}
	mid := FindMidNumber(arr)
	fmt.Printf("mid:%d\n", mid)
}

func FindMidNumber(arr []int) int {
	sort.Ints(arr)
	n := len(arr)
	if n%2 == 1 {
		// 基数个数
		return arr[n/2]
	} else {
		// 偶数个数
		n1 := arr[(n/2)-1]
		n2 := arr[n/2]
		return (n1 + n2) / 2
	}
}
