package main

import (
    "fmt"
)

func abc1(s []string) {
    fmt.Printf("%p\n", s)
    s = append(s, "xx")
    fmt.Println("abc1: ", s)
    fmt.Printf("%p\n", s) // slice地址未改变
}

func abc2(s *[]string) {
    fmt.Printf("%p\n", s)
    *s = append(*s, "xx")
    fmt.Println("abc2: ", *s)
    fmt.Printf("%p\n", s)
}

func main() {
    //s1 := []string{"1", "2"}
    s1 := make([]string, 0, 4)
    fmt.Printf("%p\n", s1)
    abc1(s1)
    fmt.Printf("%p %v\n", &s1, s1) // slice地址改变
    //s2 := []string{"1", "2"}
    //s2 := make([]string, 0, 4)
    s2 := make([]string, 0)
    fmt.Printf("%p\n", &s2)
    abc2(&s2)
    fmt.Printf("%p %v\n", &s2, s2) // 地址都不会变
}
