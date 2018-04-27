package main

import (
    "runtime/pprof"
    "fmt"
    "log"
    "os"
)

func main() {
    f, err := os.Create("cpu_profile")
    if err != nil {
        log.Fatal(err)
    }
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()

    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }

    f2, err := os.Create("mem_profile")
    pprof.WriteHeapProfile(f2)
    f2.Close()
}
