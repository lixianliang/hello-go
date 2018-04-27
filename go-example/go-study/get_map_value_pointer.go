package main

import (
    "fmt"
)

//type xxx []string

func main() {
    //kvs1 := map[string]string{"a":"1", "b":"2"}
    //p1 := &(kvs1["a"])
    //fmt.Println(p1)
    //kvs2 := map[string][]string{"a":{"1", "2"}, "b":{"3", "4"}}
    //p2 := &(kvs2["a"])
    //fmt.Println(p2)

    s := []string{"1", "2"}
    fmt.Println(s)
    x := &(s[1])
    fmt.Println(x)
    fmt.Println(*x)
    *x = "5"
    fmt.Println(*x)
    fmt.Println(s)
}
