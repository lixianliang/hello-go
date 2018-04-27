package main

import (
    "fmt"
)

func main() {
    slice1 := make([]int, 6)
    slice2 := slice1[2:4]
    //delete(slice2[0])
    //x := []int{3, 4, 5}
    slice2 = append(slice2, 1, 2)
    slice2 = append(slice2, 1, 2, 3)
    fmt.Printf("%p %#v %d %d \n", slice1, slice1, len(slice1), cap(slice1))
    fmt.Printf("%p %#v %d %d \n", slice2, slice2, len(slice2), cap(slice2))
}
