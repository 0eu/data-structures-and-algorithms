package stack

import (
	list "container/list"
	"errors"
)

var (
	ErrorExceededCapacity = errors.New("capacity is exceeded")
	ErrorEmptyStack       = errors.New("can't perform pop, peek on empty stack")
)

// Stack is an ADT.
type Stack interface {
	Push(element interface{}) error
	Peek() (interface{}, error)
	Pop() (interface{}, error)
	IsFull() bool
	IsEmpty() bool
	Size() int
}

type LinkedListStack struct {
	container *list.List
	length    int
	capacity  int
}

func NewLinkedListStack(capacity int) *LinkedListStack {
	return &LinkedListStack{
		container: list.New(),
		length:    0,
		capacity:  capacity,
	}
}

func (s *LinkedListStack) Push(element interface{}) error {
	if s.IsFull() {
		return ErrorExceededCapacity
	}
	s.container.PushBack(element)
	s.length++
	return nil
}

func (s *LinkedListStack) IsFull() bool {
	if s.length < s.capacity {
		return false
	}
	return true
}

func (s *LinkedListStack) IsEmpty() bool {
	if s.length == 0 {
		return true
	}
	return false
}

func (s *LinkedListStack) Size() int {
	return s.length
}

func (s *LinkedListStack) Pop() (interface{}, error) {
	return 0, nil
}

func (s *LinkedListStack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrorEmptyStack
	}
	return s.container.Back().Value, nil
}
