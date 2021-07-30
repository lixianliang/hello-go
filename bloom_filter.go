package main

import (
	"fmt"
)

const (
	//	Seeds        = []int32{7, 11, 13, 31, 37, 61, 73, 97}
	KDefaultsize = 16 * 100000000
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

type SimpleHash struct {
	seed     int32
	capacity int32
}

func (hash SimpleHash) Hash(s string) int32 {
	var res int32 = 0
	for _, c := range s {
		res = res*hash.seed + c
	}
	return (hash.capacity - 1) & res // 求位1计算
}

type BloomFilter struct {
	//	seeds  []int
	hashs  []SimpleHash
	bitmap *Bitmap
}

func NewBloomFilter(seeds []int32) *BloomFilter {
	bitmap := NewBitmap(KDefaultsize)
	hashs := make([]SimpleHash, 0, len(seeds))
	for _, seed := range seeds {
		hashs = append(hashs, SimpleHash{seed: seed, capacity: KDefaultsize})
	}
	return &BloomFilter{hashs: hashs, bitmap: bitmap}
}

func (bf *BloomFilter) Add(s string) {
	for _, h := range bf.hashs {
		bf.bitmap.SetBit(int(h.Hash(s)))
	}
}

func (bf *BloomFilter) Contain(s string) bool {
	re := false
	for _, h := range bf.hashs {
		re = bf.bitmap.GetBit(int(h.Hash(s)))
		if re == false {
			return false
		}
	}
	return true
}

func main() {
	seeds := []int32{7, 11, 13, 31, 37, 61, 73, 97}
	bf := NewBloomFilter(seeds)
	bf.Add("lixianliang")
	fmt.Println(bf.Contain("lixianliang"))
	fmt.Println(bf.Contain("lixianliangx"))
}
