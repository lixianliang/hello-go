package main

import "fmt"

type Human interface {
	Say()
}

type Man struct {
}

type Woman struct {
}

func (m Man) Say() {
	fmt.Println("I'm a man")
}

func (w *Woman) Say() {
	fmt.Println("I'm a woman")
}

/*func (w Woman) Say() {
	fmt.Println("I'm a woman")
}*/

/*
	接口方法调用需要是recvier为指针，可寻址
*/

func main() {
	humans := []Human{Man{}, Woman{}}
	for _, human := range humans {
		human.Say()
	}
}
