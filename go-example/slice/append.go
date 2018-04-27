package main

import (
    "fmt"
)

func main() {
    slice := make([]int, 0, 5)
    for i := 0; i < 10; i++ {
        slice = Extend(slice, i)
        fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
        fmt.Println("address of 0th element:", &slice[0])
    }

    slice = []int{0, 1, 2, 3, 4}
    fmt.Println(slice)
    slice = Append(slice, 5, 6, 7, 8)
    fmt.Println(slice)

    slice = []int{0, 1, 2, 3, 4}
    fmt.Println(slice)
    slice = Append2(slice, 5, 6, 7, 8)
    fmt.Println(slice)
}

func Append(slice []int, items ...int) []int {
    for _, item := range items {
        slice = Extend(slice, item)
    }
    return slice
}

func Extend(slice []int, element int) []int {
    n := len(slice)
    if n == cap(slice) {
        newSlice := make([]int, len(slice), 2*len(slice)+1)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0 : n+1]
    slice[n] = element
    return slice
}

func Append2(slice []int, elements ...int) []int {
    n := len(slice)
    total := len(slice) + len(elements)
    if total > cap(slice) {
        newSize := total*3/2 + 1
        newSlice := make([]int, total, newSize)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[:total]
    copy(slice[n:], elements)
    return slice
}
