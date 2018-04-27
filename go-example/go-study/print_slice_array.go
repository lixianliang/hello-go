package main

import (
    "fmt"
    "unsafe"

    "mypackage"
)

func main() {
    var s []int
    fmt.Printf("%p\n", &s)
    s = append(s, 1)
    fmt.Printf("%p\n", &s)
    s = append(s, ([]int{1, 2, 3})...)
    fmt.Printf("%p\n", &s)
    //fmt.Printf("%d %d\n", s.len, s.cap)

    a := mypackage.Xxx{A: 3, B: 4}
    fmt.Println(a.A)
    fmt.Printf("%p\n", &a)
    fmt.Printf("%d\n", unsafe.Offsetof(a.A))
    fmt.Printf("%d\n", unsafe.Offsetof(a.B))
    fmt.Printf("0x%0x\n", unsafe.Pointer(&a))
    fmt.Printf("%d\n", *(*int64)(unsafe.Pointer(&a)))
    fmt.Printf("0x%0x\n", *(*int64)(unsafe.Pointer(&s)))
}
