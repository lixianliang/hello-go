package main

import (
	"fmt"
)

type IOranization interface {
	Count() int
}

type Employee struct {
	Name string
}

func (Employee) Count() int {
	return 1
}

type Department struct {
	Name string

	SubOrganizations []IOranization
}

func (d Department) Count() int {
	c := 0
	for _, org := range d.SubOrganizations {
		c += org.Count()
	}
	return c
}

func (d *Department) AddSub(org IOranization) {
	d.SubOrganizations = append(d.SubOrganizations, org)
}

func NewOrganization() IOranization {
	root := &Department{Name: "root"}
	for i := 0; i < 10; i++ {
		root.AddSub(&Employee{})
		root.AddSub(&Department{Name: "sub", SubOrganizations: []IOranization{&Employee{}}})
	}

	return root
}

func main() {
	got := NewOrganization().Count()
	fmt.Printf("%v\n", got)
}
