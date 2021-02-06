package LinkedList

import "errors"

var (
	ErrorEmptyList = errors.New("a list is empty")
)

// Node is a unit in a linked list, that stores data and a pointer to the next element.
type Node struct {
	Value interface{}
	Next  *Node
}

// LinkedList is a struct that has pointers to the first and last nodes and also size of a list.
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// NewLinkedList builds an empty linked list and returns its pointer.
func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Size returns the number of nodes in a list.
func (l *LinkedList) Size() int {
	return l.size
}

// Append adds an element to the end of a list.
func (l *LinkedList) Append(value interface{}) {
	newNode := &Node{Value: value, Next: nil}
	if l.head != nil {
		l.tail.Next, l.tail = newNode, newNode
	} else {
		l.head, l.tail = newNode, newNode
	}
	l.size++
}

func (l *LinkedList) Prepend(value interface{}) {
	newNode := &Node{Value: value, Next: l.head}
	if l.head != nil {
		l.head = newNode
	} else {
		l.head, l.tail = newNode, newNode
	}
	l.size++
}

func (l *LinkedList) PeekFirst() (interface{}, error) {
	if l.head != nil {
		return l.head.Value, nil
	}
	return nil, ErrorEmptyList
}

func (l *LinkedList) PeekLast() (interface{}, error) {
	if l.tail != nil {
		return l.tail.Value, nil
	}
	return nil, ErrorEmptyList
}

// Count returns a number of times a given value is occurred in a list.
func (l *LinkedList) Count(value interface{}) int {
	current, count := l.head, 0
	for current != nil {
		if current.Value == value {
			count++
		}
		current = current.Next
	}
	return count
}
