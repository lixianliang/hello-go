package main

import (
    "fmt"
    //"time"
)

func main() {
    ch1 := make (chan int, 1)
    ch2 := make (chan int, 1)

    select {
    case x1 := <-ch1:
        fmt.Println("ch1 pop one element %v", x1)
    case x2 := <-ch2:
        fmt.Println("ch2 pop one element %v", x2)
    default:
        fmt.Println("default")
    }
}
