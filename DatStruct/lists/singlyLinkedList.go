package singlyLinkedList

import "fmt"

type SinglyLinkedList struct {
	head *Node
}

type Node struct {
	val  int
	next *Node
}

func (l *SinglyLinkedList) InsertAtEnd(value int) {
	if l.isEmpty() {
		l.head = &Node{value, nil}
		return
	}
	tmp := l.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = &Node{value, nil}
}

func (l *SinglyLinkedList) InsertAtHead(value int) {
	if l.isEmpty() {
		l.head = &Node{value, nil}
		return
	}
	tmp := l.head
	l.head = &Node{value, tmp}
}

func (l *SinglyLinkedList) Delete(value int) bool {
	if l.isEmpty() {
		return false
	}
	if l.head.val == value {
		l.head = l.head.next
		return true
	}
	tmp := l.head
	for tmp.next != nil {
		if tmp.next.val == value {
			tmp.next = tmp.next.next
			return true
		}
		tmp = tmp.next
	}
	return false
}

func (l *SinglyLinkedList) DeleteAtHead() bool {
	if l.isEmpty() {
		return false
	}
	l.head = l.head.next
	return true
}

func (l *SinglyLinkedList) Search(value int) *Node {
	if l.isEmpty() {
		return nil
	}
	if l.head.val == value {
		return l.head
	}
	tmp := l.head
	for tmp.next != nil {
		tmp = tmp.next
		if tmp.val == value {
			return tmp
		}
	}
	return nil
}

func (l *SinglyLinkedList) isEmpty() bool {
	return l.head == nil
}

func (l *SinglyLinkedList) Show() {
	if l.isEmpty() {
		fmt.Println("List is empty")
		return
	}
	fmt.Printf("%d", l.head.val)
	tmp := l.head
	for tmp.next != nil {
		tmp = tmp.next
		fmt.Printf(" -> %d", tmp.val)
	}
	fmt.Printf("\n")
}
