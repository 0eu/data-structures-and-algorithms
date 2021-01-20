package stack

import "container/list"

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
	return s.length >= s.capacity
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.length == 0
}

func (s *LinkedListStack) Size() int {
	return s.length
}

func (s *LinkedListStack) Pop() (interface{}, error) {
	if lastElement, err := s.peek(); err != nil {
		return nil, err
	} else {
		s.container.Remove(lastElement)
		s.length--
		return lastElement.Value, nil
	}
}

func (s *LinkedListStack) peek() (*list.Element, error) {
	if s.IsEmpty() {
		return nil, ErrorEmptyStack
	}
	return s.container.Back(), nil
}

func (s *LinkedListStack) Peek() (interface{}, error) {
	if element, err := s.peek(); err != nil {
		return nil, err
	} else {
		return element.Value, nil
	}
}
