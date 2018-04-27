package main

import (
	"fmt"
)

type HelloHandler struct {
}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h *HelloHandler) HelloString(para string) (string, error) {
	fmt.Printf("HelloString accept params:%s\n", para)
	return "hello, world", nil
}

func (h *HelloHandler) HelloBoolean(para bool) (r bool, err error) {
	fmt.Printf("HelloBoolean accept params:%t\n", para)
	return para, nil
}

func (h *HelloHandler) HelloInt(para int32) (r int32, err error) {
	fmt.Printf("HelloInt accept params:%d\n", para)
	return para, nil
}

func (h *HelloHandler) HelloVoid() (err error) {
	fmt.Printf("HelloVoid\n")
	return nil
}

func (h *HelloHandler) HelloNull() (r string, err error) {
	fmt.Printf("HelloNull\n")
	return "hello null", nil
}
