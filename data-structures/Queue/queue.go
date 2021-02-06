package queue

import "errors"

var (
	ErrorExceededCapacity = errors.New("capacity is exceeded")
	ErrorWrongCapacity    = errors.New("capacity should be in range [1, 1024]")
	ErrorEmptyQueue       = errors.New("the operation can't be permitted on an empty queue")
)

type Queue interface {
	Enqueue(element interface{}) error
	Dequeue() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
	Size() int
}
