package stack

import (
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
