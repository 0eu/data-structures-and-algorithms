package stack

type ArrayStack struct {
	container []interface{}
	length    int
	capacity  int
}

func NewArrayStack(capacity int) *ArrayStack {
	return &ArrayStack{
		container: make([]interface{}, capacity),
		length:    0,
		capacity:  capacity,
	}
}

func (s *ArrayStack) Push(element interface{}) error {
	if s.IsFull() {
		return ErrorExceededCapacity
	}
	s.container[s.length] = element
	s.length++
	return nil
}

func (s *ArrayStack) IsFull() bool {
	return s.length >= s.capacity
}

func (s *ArrayStack) IsEmpty() bool {
	return s.length == 0
}

func (s *ArrayStack) Size() int {
	return s.length
}

func (s *ArrayStack) Pop() (interface{}, error) {
	value, err := s.Peek()
	if err != nil {
		return nil, err
	}
	s.container[s.length-1] = nil
	s.length--
	return value, nil
}

func (s *ArrayStack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrorEmptyStack
	}
	return s.container[s.length-1], nil
}
