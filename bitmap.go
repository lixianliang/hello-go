package main

import (
	"fmt"
)

type Bitmap struct {
	size int
	bits []int32
}

func NewBitmap(size int) *Bitmap {
	len := size / 32
	return &Bitmap{
		size: len,
		bits: make([]int32, len, len),
	}
}

func (bit *Bitmap) SetBit(val int) {
	index := val / 32
	offset := val % 32
	bit.bits[index] |= (1 << offset)
}

func (bit *Bitmap) GetBit(val int) bool {
	index := val / 32
	offset := val % 32
	x := bit.bits[index] & (1 << offset)
	if x > 0 { // val > 0 is contain 1 bit
		return true
	} else {
		return false
	}
}

func main() {
	bitmap := NewBitmap(10000000)
	bitmap.SetBit(9999)
	fmt.Println("test", bitmap.GetBit(9999))
}
