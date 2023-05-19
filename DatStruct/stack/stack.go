package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	head *node
}

type node struct {
	val  int
	next *node
}

func (s *Stack) Push(value int) {
	if s.IsEmpty() {
		s.head = &node{value, nil}
		return
	}
	newNode := node{value, s.head}
	s.head = &newNode
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	tmp := s.head.val
	s.head = s.head.next
	return tmp, nil
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	return s.head.val, nil
}

func (s *Stack) Show() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Printf("%d", s.head.val)
	tmp := s.head
	for tmp.next != nil {
		tmp = tmp.next
		fmt.Printf(" -> %d", tmp.val)
	}
	fmt.Printf("\n")
}
