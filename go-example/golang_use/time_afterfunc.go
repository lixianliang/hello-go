package main

import (
    "fmt"
    "time"
)

func main() {
    var t *time.Timer

    f := func() {
        fmt.Printf("Expiration timer: %v\n", time.Now())
        fmt.Printf("c len: %d\n", len(t.C))
    }

    t = time.AfterFunc(1 * time.Second, f)
    time.Sleep(2*time.Second)
}
