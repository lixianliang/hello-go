package main

import (
	"fmt"
	//	"log"
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
	if len(arr) < 2 {
		return 0, 0, fmt.Errorf("arr must > 2")
	}

	isFind := false
	a, b := 0, 0
	// 将数组value和index当作key value放入map中
	//arrMap := map[int]int{}
	arrMap := make(map[int]int, len(arr))
	for i, val := range arr {
		arrMap[val] = i
	}

	for _, val := range arr {
		x := s - val
		if _, ok := arrMap[x]; ok {
			if isFind {
				if val*x > a*b {
					a = val
					b = x
				}
			} else {
				isFind = true
				a = val
				b = x
			}
		}
	}

	if !isFind {
		return a, b, fmt.Errorf("arr not find")
	}

	return a, b, nil
}
