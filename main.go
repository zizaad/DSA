package main

import (
	"DatStruct/stack"
	"fmt"
)

func main() {
	stack := stack.Stack{nil}
	stack.Show()
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Show()
	fmt.Println("Top: %d\n", stack.Top())
	stack.Show()
	fmt.Println("Pop: %d\n", stack.Pop())
}
