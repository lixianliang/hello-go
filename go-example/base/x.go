package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3}
	s1 := make([]int, 0, 4)
	s1 = append(s1, 5)
	s1 = append(s1, s...)
	fmt.Println(s1)

	s2 := make([]int, 4, 5)
	s2[0] = 5
	copy(s2[1:], s)
	fmt.Println(s2)
}
