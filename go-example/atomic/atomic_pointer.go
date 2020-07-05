package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

func main() {
	var p unsafe.Pointer
	newp := 42
	atomic.CompareAndSwapPointer(&p, nil, unsafe.Pointer(&newp))

	v := (*int)(p)
	fmt.Println(*v)
}
