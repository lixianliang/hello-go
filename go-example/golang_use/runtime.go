package main

import (
    "fmt"
    "runtime"
    "time"
)

func main() {
    runtime.GOMAXPROCS(1)
    var m runtime.MemStats


    go func() {
        time.Sleep(time.Second)
    }()

    runtime.ReadMemStats(&m)
    fmt.Printf("size: %dkb, number cpu: %d, number goroutine:%d \n", m.Alloc/1014, runtime.NumCPU(), runtime.NumGoroutine())
    for {
        time.Sleep(time.Second)
    }
}
