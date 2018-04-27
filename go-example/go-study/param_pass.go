package main

import (
    "fmt"
)

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    dump(arr)
    dump1(&arr)
}

func dump(arr [5]int) {
    fmt.Println(arr)
}

func dump1(arr *[5]int) {
    fmt.Println(arr)
}
