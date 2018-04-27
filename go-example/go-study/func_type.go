package main

import (
    "fmt"
)

func square(x int) int {
    return x * x
}

func main() {
    fmt.Println(square(3))

    f := square
    fmt.Println(f(3))
    fmt.Printf("%T", f)

}
