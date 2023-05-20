package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	head *node
}

type node struct {
	val  int
	next *node
}

func (q *Queue) Enqueue(value int) {
	if q.isEmpty() {
		q.head = &node{value, nil}
		return
	}
	tmp := q.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = &node{value, nil}
}

func (q *Queue) Dequeue() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("Queue is empty")
	}
	tmp := q.head.val
	q.head = q.head.next
	return tmp, nil
}

func (q *Queue) isEmpty() bool {
	return q.head == nil
}

func (q *Queue) Top() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("Queue is empty")
	}
	return q.head.val, nil
}

func (q *Queue) Show() {
	if q.isEmpty() {
		fmt.Println("Queue is empty")
		return
	}
	fmt.Printf("%d", q.head.val)
	tmp := q.head
	for tmp.next != nil {
		fmt.Printf(" <- %d", tmp.val)
	}
	fmt.Printf("\n")
}
