package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{7, 11, 15, 1, 2, 4}
	mid := FindMidNumber2(arr)
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

func FindMidNumber2(arr []int) int {
	// arr数组被分割成两段有序的数组，可以找到对应的临界点
	split := 0 // 分割索引默认为0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			split = i + 1
			break
		}
	}

	n := len(arr)
	if n%2 == 1 {
		// 基数个数
		index := (n/2 + split) % n
		return arr[index]
	} else {
		// 偶数个数
		index1 := (n/2 - 1 + split) % n
		index2 := (n/2 + split) % n
		n1 := arr[index1]
		n2 := arr[index2]
		return (n1 + n2) / 2
	}
}
