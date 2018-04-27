package main

import (
    "fmt"
    "time"
)

func main() {
    var ticker *time.Ticker = time.NewTicker(1 * time.Second)

    go func() {
        for t := range ticker.C {
            fmt.Println("Ticker at ", t)
        }
    }()

    time.Sleep(time.Second*5)
    ticker.Stop()
    fmt.Println("Ticker stopped")
}
