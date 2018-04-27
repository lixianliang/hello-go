package main

import (
    "fmt"
)

func main() {
    s := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

    s1 := s[0:3]
    s2 := s[4:]

    fmt.Println(len(s1), cap(s1))
    fmt.Println(len(s2), cap(s2))
}

