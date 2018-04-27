package main

import (
    "fmt"
    "unsafe"
)

func main() {
    i := int(1)
    fmt.Println(unsafe.Sizeof(i)) // 8
    j := 1
    fmt.Println(unsafe.Sizeof(j)) // 8
    u := uint(1)
    fmt.Println(unsafe.Sizeof(u)) // 8

    s := "中国"
    fmt.Println(len(s)) // 6
}
