package main

import (
	"fmt"
)

// 饮料接口
type Beverage interface {
	getDescription() string
	cost() int
}

// 咖啡实现
type Coffee struct {
	description string
}

func (this Coffee) getDescription() string {
	return this.description
}

func (this Coffee) cost() int {
	return 1
}

// mocha实现
type Mocha struct {
	beverage    Beverage
	description string
}

func (this Mocha) getDescription() string {
	return fmt.Sprintf("%s, %s", this.beverage.getDescription(), this.description)
}

func (this Mocha) cost() int {
	return this.beverage.cost() + 1
}

// Whip实现
type Whip struct {
	beverage    Beverage
	description string
}

func (this Whip) getDescription() string {
	return fmt.Sprintf("%s, %s", this.beverage.getDescription(), this.description)
}

func (this Whip) cost() int {
	return this.beverage.cost() + 1
}

func main() {
	var beverage Beverage
	beverage = Coffee{description: "houseBlend"}
	// 给咖啡加上Mocha
	beverage = Mocha{beverage: beverage, description: "Mocha"}
	// 给咖啡加上Whip
	beverage = Whip{beverage: beverage, description: "Whip"}
	// 最后计算Coffee价格
	fmt.Println(beverage.getDescription(), ", cost is ", beverage.cost())
}
