package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan int, 1)
    ch2 := make(chan int, 1)

    ch1 <- 1

    select {
    case e1 := <-ch1:
        fmt.Printf("1th case is select, e1=%v", e1)
    case e2 := <-ch2:
        fmt.Printf("2th case is select, e2=%v", e2)
    case <- time.After(2*time.Second):
        fmt.Println("Time out")
    }
}
