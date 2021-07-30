package main

import (
	"fmt"
)

type Iterator interface {
	HasNext() bool
	Next()
	CurrentItem() interface{}
}

type ArrayInt []int

func (a ArrayInt) Iterator() Iterator {
	return &ArrayIntIterator{
		arrayInt: a,
		index:    0,
	}
}

type ArrayIntIterator struct {
	arrayInt ArrayInt
	index    int
}

func (iter *ArrayIntIterator) HasNext() bool {
	return iter.index < len(iter.arrayInt)-1
}

func (iter *ArrayIntIterator) Next() {
	iter.index++
}

func (iter *ArrayIntIterator) CurrentItem() interface{} {
	return iter.arrayInt[iter.index]
}

func main() {
	data := ArrayInt{1, 3, 5, 7, 8}
	iterator := data.Iterator()
	i := 0
	for iterator.HasNext() {
		fmt.Println(iterator.CurrentItem())
		iterator.Next()
		i++
	}
}
