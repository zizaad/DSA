package main

import (
	"fmt"

	dll "github.com/zizaad/DSA/DatStruct/lists/doublyLinkedList"
	sll "github.com/zizaad/DSA/DatStruct/lists/singlyLinkedList"
	"github.com/zizaad/DSA/DatStruct/queue"
	"github.com/zizaad/DSA/DatStruct/stack"
)

func main() {
	// checkStack()
	// checkQueue()
	// checkSinglyLinkedList()
	checkDoublyLinkedList()
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

func checkSinglyLinkedList() {
	ls := sll.SinglyLinkedList{}
	ls.Show()
	ls.InsertAtEnd(1)
	ls.InsertAtEnd(2)
	ls.Show()
	ls.InsertAtHead(3)
	ls.Show()
	tmp := ls.Search(2)
	fmt.Printf("Search(2) = %d\n", tmp.Val)
	tmp = ls.Search(4)
	if tmp == nil {
		fmt.Printf("Search(4) = nil\n")
	}
	ls.Delete(2)
	ls.Show()
	ls.DeleteAtHead()
	ls.Show()
	ls.DeleteAtHead()
	ls.Show()
	ok := ls.DeleteAtHead()
	if !ok {
		fmt.Printf("List is empty. You can't delete anything.\n")
	}
}

func checkDoublyLinkedList() {
	ls := dll.DoublyLinkedList{}
	ls.Show()
	ls.InsertAtEnd(1)
	ls.InsertAtEnd(2)
	ls.Show()
	ls.InsertAtHead(3)
	ls.Show()
	tmp := ls.Search(2)
	fmt.Printf("Search(2) = %d\n", tmp.Val)
	tmp = ls.Search(4)
	if tmp == nil {
		fmt.Printf("Search(4) = nil\n")
	}
	ls.Delete(2)
	ls.Show()
	ls.DeleteAtHead()
	ls.Show()
	ls.DeleteAtHead()
	ls.Show()
	ok := ls.DeleteAtHead()
	if !ok {
		fmt.Printf("List is empty. You can't delete anything.\n")
	}
}
