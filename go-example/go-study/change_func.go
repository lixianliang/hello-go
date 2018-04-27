package main

import (
    "fmt"
)

func sum(vals ...int) int {
    total := 0
    for _, val := range vals {
        total += val
    }
    return total
}

func main() {
    var arr = [...]int{1, 2, 3}
    fmt.Println(sum(arr[:]...))
    //fmt.Println(sum(arr...))
}
