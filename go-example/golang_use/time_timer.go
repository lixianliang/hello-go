package main

import (
    "fmt"
    "time"
)

func main() {
    chl1 := make(chan int, 20)
    sign := make(chan byte, 1)

    for i := 0; i < 20; i++ {
        chl1 <- i
    }

    go func() {
        var e int
        ok := true
        var timer *time.Timer
        for {
            select {
            case e = <- chl1:
                fmt.Printf("chl1 -> %d\n", e)
            case <- func() <-chan time.Time {
                if timer == nil {
                    fmt.Println("x")
                    timer = time.NewTimer(time.Millisecond)
                } else {
                    fmt.Println("xx")
                    timer.Reset(time.Millisecond)
                }
                fmt.Println("xxx")
                return timer.C
            }():
            fmt.Println("Timeout")
            ok = false
            break
        }
        if !ok {
            sign <- 0
            break
        }
        }
    }()

    <- sign
}
