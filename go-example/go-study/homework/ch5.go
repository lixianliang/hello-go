package main

import (
    "errors"
    "fmt"
)

func max(vals ...int) (int, error) {
    max := 0
    if len(vals) < 1 {
        return 0, errors.New("params must equal zero")
    }
    for i, val := range vals {
        if i == 0 {
            max = val
        }
        if val > max {
            max = val
        }
    }
    return max, nil
}

func min(vals ...int) (int, error) {
    min := 0
    if len(vals) < 1 {
        return 0, errors.New("params must equal zero")
    }
    for i, val := range vals {
        if i == 0 {
            min = val
        }
        if val < min {
            min = val
        }
    }
    return min, nil
}


func main() {
    maxx, _ := max(4, 9 , 6, 5)
    fmt.Println(maxx)
    min, _ := min(4, 9, 6, 5)
    fmt.Println(min)
    //s := []int{0, 1}
    s := []int{}
    x, err := max(s...)
    fmt.Printf("%d %v", x, err)
}
