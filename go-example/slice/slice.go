package main

import (
    "fmt"
)

func main() {
    var iBuffer [10]int
    slice := iBuffer[0:0]
    for i := 0; i < 20; i++ {   // 超过10会报错，最底层的数组并不是伸缩的
        slice = Extend(slice, i)
        fmt.Println(slice)
    }
}

func Extend(slice []int, element int) []int {
    n := len(slice)
    slice = slice[0 : n+1]
    slice[n] = element
    return slice
}

func Insert(slice []int, index, value int) []int {
    slice = slice[0 : len(slice)+1]
    copy(slice[index+1:], slice[index:])
    slice[index] = value
    return slice
}
