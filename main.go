package main

import (
	"fmt"

	"github.com/zizaad/DSA/DatStruct/stack"
)

func main() {
	checkStack()
}

func checkStack() {
	stack := stack.Stack{}
	stack.Show()
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Show()
	tmp, err := stack.Top()
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	fmt.Printf("Top: %d\n", tmp)

	stack.Show()
	tmp, err = stack.Pop()
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	fmt.Printf("Pop: %d\n", tmp)

	stack.Show()
}
