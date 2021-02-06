package queue

const MAX_CAPACITY = 1 << 10

type ArrayQueue struct {
	front     int
	end       int
	capacity  int
	container []interface{}
}

func (q *ArrayQueue) Enqueue(element interface{}) error {
	if q.Size() == q.capacity-1 {
		return ErrorExceededCapacity
	}
	q.container[q.end] = element
	q.end++
	if q.end == q.capacity {
		q.end = 0
	}
	return nil
}

func (q *ArrayQueue) Dequeue() (interface{}, error) {
	element, err := q.Peek()
	if err != nil {
		return nil, err
	}
	q.front++
	if q.front == q.capacity {
		q.front = 0
	}
	return element, nil
}

func (q *ArrayQueue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrorEmptyQueue
	}
	return q.container[q.front], nil
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.front == q.end
}

func (q *ArrayQueue) Size() int {
	if q.front > q.end {
		return q.end + q.capacity - q.front
	}
	return q.end - q.front
}

func NewArrayQueue(capacity int) (*ArrayQueue, error) {
	if capacity <= 0 || capacity > MAX_CAPACITY {
		return nil, ErrorWrongCapacity
	}
	return &ArrayQueue{
		front:     0,
		end:       0,
		capacity:  capacity + 1,
		container: make([]interface{}, capacity+1),
	}, nil
}
