package main

import (
    "fmt"
)

func main() {
    str := "中国共产党123"
    for i, s := range str {
        fmt.Printf("%d %q %d\n", i, s, s)
    }

    for i := 0; i < len(str); i++ {
        fmt.Printf("%d %q %d\n", i, str[i], str[i])
    }
}
