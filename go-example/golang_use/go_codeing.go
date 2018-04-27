package main

import "fmt"

func main() {
    str :=  "李显良"

    fmt.Printf("len:%d str:%s\n", len(str), str)
    b := []byte(str)
    fmt.Printf("len:%d str:%b\n", len(b), b)

}
