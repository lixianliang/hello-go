package main

import (
    "fmt"
)

func main() {
    valueParamter()
    valueParamter2()
    pointerParamter()
}

func valueParamter() {
    var buffer [256]byte
    slice := buffer[10:20]
    for i := 0; i < len(slice); i++ {
        slice[i] = byte(i)
    }
    fmt.Println("before0", slice)
    AddOneToEachElement(slice)
    fmt.Println("after0", slice)
}

func AddOneToEachElement(slice []byte) {
    for i := range slice {
        slice[i]++
    }
}

func valueParamter2() {
    var buffer [256]byte
    slice := buffer[10:20]
    fmt.Println("beofre1: len(slice) =", len(slice))
    newSlice := SubtractOneFromLength(slice)
    fmt.Println("after1: len(slice) =", len(slice))
    fmt.Println("after1: len(newSlice) =", len(newSlice))
}

func SubtractOneFromLength(slice []byte) []byte {
    slice = slice[0 : len(slice)-1]
    return slice
}

func pointerParamter() {
    var buffer [256]byte
    slice := buffer[10:20]
    fmt.Println("pointer before: len(slice)=", len(slice))
    PtrSubtractOneFromLength(&slice)
    fmt.Println("pointer after: len(slice)=", len(slice))
}

func PtrSubtractOneFromLength(slicePtr *[]byte) {
    slice := *slicePtr
    *slicePtr = slice[0 : len(slice)-1]
}
