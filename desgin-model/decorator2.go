package main

import (
	"fmt"
	"time"
)

// 为函数类型设置别名提高代码可读性
type MultiPlyFunc func(int, int) int

// 乘法运算函数
func multiply1(a, b int) int {
	return a * b
}

// 乘法运算函数2
func multiply2(a, b int) int {
	return a << b
}

// 通过高阶函数在不侵入原有函数实现的前提下计算乘法函数执行时间
func execTime(f MultiPlyFunc) MultiPlyFunc {
	return func(a, b int) int {
		start := time.Now()
		c := f(a, b)
		end := time.Since(start)
		fmt.Printf("执行耗时：%v\n", end)
		return c
	}
}

func main() {
	x := multiply1(2, 8)
	fmt.Println("x: ", x)
	a := 2
	b := 8
	fmt.Println("算术运算： ")
	decorator1 := execTime(multiply1)
	c := decorator1(a, b)
	fmt.Printf("%d x %d = %d\n", a, b, c)

	fmt.Println("位运算： ")
	decorator2 := execTime(multiply2)
	a = 1
	b = 4
	c = decorator2(a, b)
	fmt.Printf("%d << %d = %d\n", a, b, c)
}
