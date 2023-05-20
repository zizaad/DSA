package main

import (
	"fmt"

	"github.com/zizaad/DSA/DatStruct/queue"
	"github.com/zizaad/DSA/DatStruct/stack"
)

func main() {
	// checkStack()
	checkQueue()
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
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Top: %d\n", tmp)

	stack.Show()
	tmp, err = stack.Pop()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Pop: %d\n", tmp)

	stack.Show()
}

func checkQueue() {
	q := queue.Queue{}
	q.Show()
	q.Enqueue(5)
	q.Enqueue(15)
	q.Enqueue(45)
	q.Show()
	tmp, err := q.Dequeue()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Dequeue: %d\n", tmp)
	q.Show()
	_, err = q.Dequeue()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	_, err = q.Dequeue()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	q.Show()
	_, err = q.Dequeue()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
