package main

import (
    "fmt"
    "time"
)

func main() {
    var ticker *time.Ticker = time.NewTicker(100 * time.Millisecond)

    num := 2
    c := make(chan int, num)
    go func() {
        for t := range ticker.C {
            c <- 1
            fmt.Println("Tick at", t)
        }
    }()

    time.Sleep(time.Millisecond * 1500)
    ticker.Stop()
    fmt.Println("Ticker stopped")
}
