package main

import (
	"fmt"
)

type sss struct {
	a, b, c int
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13}
	//x := int{}
	add := []sss{}
	sub := []sss{}
	xx1 := []sss{}
	xx2 := []sss{}
	for i, x := range arr {
		for _, y := range arr[i+1:] {
			if x+y <= arr[len(arr)-1] {
				fmt.Printf("i:%d j:%d add:%d\n", x, y, x+y)
				add = append(add, sss{x, y, x + y})
			}
		}
	}

	for i, x := range arr {
		for _, y := range arr[i+1:] {
			xx := y - x
			if xx != x && xx != y {
				fmt.Printf("i:%d j:%d sub:%d\n", x, y, y-x)
				sub = append(sub, sss{x, y, y - x})
			}
		}
	}

	for i, x := range arr {
		for _, y := range arr[i+1:] {
			xx := x * y
			if xx <= arr[len(arr)-1] && xx != x && xx != y {
				fmt.Printf("i:%d j:%d xx1:%d\n", x, y, x*y)
				s := sss{x, y, x * y}
				xx1 = append(xx1, s)
			}
		}
	}

	for i, x := range arr {
		for _, y := range arr[i+1:] {
			xx := y % x
			yy := y / x
			if xx == 0 && yy != y && yy != x {
				//if y%
				fmt.Printf("i:%d j:%d xx2:%d\n", x, y, y/x)
				xx2 = append(xx2, sss{x, y, y / x})
			}
		}
	}

	for _, x := range xx2 {
		for _, y := range xx1 {
			for _, z := range sub {
				for _, zz := range add {
					if x.a == y.a || x.a == y.b || x.a == y.c ||
						x.b == y.a || x.b == y.b || x.b == y.c ||
						x.c == y.a || x.c == y.b || x.c == y.c {
						continue
					}
					if x.a == z.a || x.a == z.b || x.a == z.c ||
						x.b == z.a || x.b == z.b || x.b == z.c ||
						x.c == z.a || x.c == z.b || x.c == z.c {
						continue
					}
					if x.a == zz.a || x.a == zz.b || x.a == zz.c ||
						x.b == zz.a || x.b == zz.b || x.b == zz.c ||
						x.c == zz.a || x.c == zz.b || x.c == zz.c {
						continue
					}

					if y.a == z.a || y.a == z.b || y.a == z.c ||
						y.b == z.a || y.b == z.b || y.b == z.c ||
						y.c == z.a || y.c == z.b || y.c == z.c {
						continue
					}
					if y.a == zz.a || y.a == zz.b || y.a == zz.c ||
						y.b == zz.a || y.b == zz.b || y.b == zz.c ||
						y.c == zz.a || y.c == zz.b || y.c == zz.c {
						continue
					}

					if z.a == zz.a || z.a == zz.b || z.a == zz.c ||
						z.b == zz.a || z.b == zz.b || z.b == zz.c ||
						z.c == zz.a || z.c == zz.b || z.c == zz.c {
						continue
					}

					fmt.Printf("%v %v %v %v\n", x, y, z, zz)
				}
			}
		}
	}
}
