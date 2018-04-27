package main

import (
    "fmt"

    "go-example/go-study/homework/intset"
)

func main() {
    var s intset.IntSet
    fmt.Println(s.Len())
    s.Add(1)
    s.Add(2)
    s.Add(18)
    s.Add(30)
    s.Add(1000)
    fmt.Println(s.Len())

    ss := s.Copy()
    ss.Remove(2)
    fmt.Println(ss.Len())
    fmt.Println(ss.String())
}
