// Circular Singly Linked List

package circularlinkedlist

import "fmt"

type CircularLinkedList struct {
	head    *Node
	counter int
}

type Node struct {
	Val   int
	left  *Node
	right *Node
}

func (l *CircularLinkedList) isEmpty() bool {
	return l.head == nil
}

func (l *CircularLinkedList) InsertAtEnd(value int) {
	if l.isEmpty() {
		l.head = &Node{value, nil, nil}
		l.head.left = l.head
		l.head.right = l.head
		l.counter++
		return
	}
	tmp := l.head.left
	l.head.left = &Node{value, tmp, l.head}
	tmp.right = l.head.left
	l.counter++
}

func (l *CircularLinkedList) InsertAtHead(value int) {
	if l.isEmpty() {
		l.head = &Node{value, nil, nil}
		l.head.left = l.head
		l.head.right = l.head
		l.counter++
		return
	}
	tmp := l.head
	l.head = &Node{value, tmp.left, tmp}
	tmp.left.right = l.head
	tmp.left = l.head
	l.counter++
}

func (l *CircularLinkedList) Delete(value int) bool {
	if l.isEmpty() {
		return false
	}
	if l.counter == 1 {
		if l.head.Val == value {
			l.head = nil
			l.counter--
			return true
		}
		return false
	}
	if l.head.Val == value {
		tmp := l.head
		l.head = l.head.right
		l.head.left = tmp.left
		tmp.left.right = l.head
		l.counter--
		return true
	}
	el := l.head.right
	for i := 1; i < l.counter; i++ {
		if el.Val == value {
			tmp := el
			el = el.right
			el.left = tmp.left
			tmp.left.right = el
			l.counter--
			return true
		}
		el = el.right
	}
	return false
}

func (l *CircularLinkedList) DeleteAtHead() bool {
	if l.isEmpty() {
		return false
	}
	if l.counter == 1 {
		l.head = nil
		l.counter--
		return true
	}
	tmp := l.head
	l.head = l.head.right
	l.head.left = tmp.left
	tmp.left.right = l.head
	l.counter--
	return true
}

func (l *CircularLinkedList) Search(value int) *Node {
	if l.isEmpty() {
		return nil
	}
	tmp := l.head
	for i := 0; i < l.counter; i++ {
		if tmp.Val == value {
			return tmp
		}
		tmp = tmp.right
	}
	return nil
}

func (l *CircularLinkedList) Show() {
	if l.isEmpty() {
		fmt.Println("List is empty")
		return
	}
	fmt.Printf("%d", l.head.Val)
	tmp := l.head.right
	for i := 1; i < l.counter; i++ {
		fmt.Printf(" <-> %d", tmp.Val)
		tmp = tmp.right
	}
	fmt.Printf("\n")
}
