package main

import (
	"fmt"
)

type ICommand interface {
	Execute() error
}

type StartCommand struct{}

func NewStartCommand() *StartCommand {
	return &StartCommand{}
}

func (c *StartCommand) Execute() error {
	fmt.Println("game strat")
	return nil
}

type ArchiveCommand struct{}

func NewArchiveCommand() *ArchiveCommand {
	return &ArchiveCommand{}
}

func (c *ArchiveCommand) Execute() error {
	fmt.Println("game archive")
	return nil
}

func main() {
}
