package main

import (
    "fmt"
)

func main() {
    fmt.Println("please enter your names")
    var b1, b2 string
    fmt.Scanf("%s, %s", &b1, &b2)
    fmt.Println("hello,", b1, "and", b2)
}

