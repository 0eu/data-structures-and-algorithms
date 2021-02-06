package queue

type Queue interface {
	Enqueue(element interface{}) error
	Dequeue() (interface{}, error)
	IsFull() bool
	IsEmpty() bool
	Size() int
}
