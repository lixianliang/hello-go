package main

import (
    "fmt"
)

func main() {
    fmt.Println("please enter your firstname and lastname")
    var a1, a2 string
    fmt.Scan(&a1, &a2)
    fmt.Println("hello", a1, "and", a2)
}
