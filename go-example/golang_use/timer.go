package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.NewTimer(2*time.Second)
    now := time.Now()
    fmt.Printf("Now time: %v\n", now)
    expire := <- t.C
    fmt.Printf("Expiration time: %v\n", expire)
}
