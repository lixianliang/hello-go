package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3}
	s1 := make([]int, 1, 5)
	s1[0] = 5
	//s1 = append(s1, 5)
	s1 = append(s1, s...)
	fmt.Println(s1)
}
