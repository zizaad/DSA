package doublyLinkedList

import "fmt"

type DoublyLinkedList struct {
	head *Node
}

type Node struct {
	Val   int
	left  *Node
	right *Node
}

func (l *DoublyLinkedList) InsertAtEnd(value int) {
	if l.isEmpty() {
		l.head = &Node{value, nil, nil}
		return
	}
	tmp := l.head
	for tmp.right != nil {
		tmp = tmp.right
	}
	tmp.right = &Node{value, tmp, nil}
}

func (l *DoublyLinkedList) InsertAtHead(value int) {
	if l.isEmpty() {
		l.head = &Node{value, nil, nil}
		return
	}
	tmp := l.head
	l.head = &Node{value, nil, tmp}
	tmp.left = l.head
}

func (l *DoublyLinkedList) Delete(value int) bool {
	if l.isEmpty() {
		return false
	}
	if l.head.Val == value {
		if l.head.right != nil {
			l.head.right.left = nil
		}
		l.head = l.head.right
		return true
	}
	tmp := l.head
	for tmp.right != nil {
		if tmp.right.Val == value {
			tmp.right.left = tmp
			tmp.right = tmp.right.right
			return true
		}
		tmp = tmp.right
	}
	return false
}

func (l *DoublyLinkedList) DeleteAtHead() bool {
	if l.isEmpty() {
		return false
	}
	if l.head.right != nil {
		l.head.right.left = nil
	}
	l.head = l.head.right
	return true
}

func (l *DoublyLinkedList) Search(value int) *Node {
	if l.isEmpty() {
		return nil
	}
	if l.head.Val == value {
		return l.head
	}
	tmp := l.head
	for tmp.right != nil {
		tmp = tmp.right
		if tmp.Val == value {
			return tmp
		}
	}
	return nil
}

func (l *DoublyLinkedList) isEmpty() bool {
	return l.head == nil
}

func (l *DoublyLinkedList) Show() {
	if l.isEmpty() {
		fmt.Println("List is empty")
		return
	}
	fmt.Printf("%d", l.head.Val)
	tmp := l.head
	for tmp.right != nil {
		tmp = tmp.right
		fmt.Printf(" <-> %d", tmp.Val)
	}
	fmt.Printf("\n")
}
