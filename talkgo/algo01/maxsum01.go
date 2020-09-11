package main

import (
	"fmt"
	"log"
)

func main() {
	arr := []int{1, 2, 4, 7, 11, 15}
	a, b, err := FindMaxSum(arr, 15)
	if err != nil {
		fmt.Println("not find")
	} else {
		fmt.Printf("arr:%v {%d, %d}\n", arr, a, b)
	}

	arr2 := []int{1, 2, 4, 7, 11, 16}
	a, b, err = FindMaxSum(arr2, 17)
	if err != nil {
		fmt.Println("not find")
	} else {
		fmt.Printf("arr2:%v {%d, %d}\n", arr2, a, b)
	}

	arr3 := []int{1, 2, 4, 7, 11, 16}
	a, b, err = FindMaxSum(arr3, 10)
	if err != nil {
		fmt.Println("not find")
	} else {
		fmt.Printf("arr3:%v {%d, %d}\n", arr3, a, b)
	}
}

func FindMaxSum(arr []int, s int) (int, int, error) {
	isFind := false
	a, b := 0, 0
	if len(arr) < 2 {
		return a, b, fmt.Errorf("arr must > 2")
	}
	for i, v := range arr {
		x := s - v //  定义x为目标值
		if x < v { // 要查的目标值已经小于当前的值
			continue
		}

		if i == len(arr)-1 {
			break
		}

		// 当前值后面查找目标值
		index, err := BinarySearch(arr[i+1:], x)
		if err != nil {
			continue
		}

		log.Printf("find match arr[%d]:%d arr[%d]:%d", i, v, i+1+index, x)
		if isFind {
			if v*x > a*b {
				a = v
				b = x
			}
		} else {
			a = v
			b = x
			isFind = true
		}
	}

	if !isFind {
		return a, b, fmt.Errorf("not found")
	}

	return a, b, nil
}

// 二分查找
func BinarySearch(arr []int, k int) (int, error) {
	low, high := 0, len(arr)-1
	for low <= high {
		//mid := (low + high) / 2
		mid := low + (high-low)/2 // 防止整数溢出
		if k < arr[mid] {
			high = mid - 1
		} else if k > arr[mid] {
			low = mid + 1
		} else {
			return mid, nil
		}
	}

	return 0, fmt.Errorf("BinarySearch not found")
}
