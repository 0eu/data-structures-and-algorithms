package queue

import (
	"testing"
)

func assertLength(t *testing.T, queue Queue, expected int) {
	t.Helper()
	actual := queue.Size()
	if actual != expected {
		t.Errorf("expected length %d, but got: %d", expected, actual)
	}
}

func assertError(t *testing.T, actual, expected error) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected error %s, but got: %s", expected, actual)
	}
}

func assertEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected error %s, but got: %s", expected, actual)
	}
}

func TestNewArrayQueue(t *testing.T) {
	t.Run("Init queue with incorrect capacity", func(t *testing.T) {
		_, err := NewArrayQueue(-1)
		assertError(t, err, ErrorWrongCapacity)

		_, err = NewArrayQueue(0)
		assertError(t, err, ErrorWrongCapacity)

		_, err = NewArrayQueue(9999)
		assertError(t, err, ErrorWrongCapacity)
	})

	t.Run("Init queue with correct capacity", func(t *testing.T) {
		_, err := NewArrayQueue(1)
		assertError(t, err, nil)
	})
}

func TestEnqueue(t *testing.T) {
	t.Run("Enqueue puts an element to the end of a queue", func(t *testing.T) {
		queue, _ := NewArrayQueue(2)

		assertLength(t, queue, 0)

		_ = queue.Enqueue(15)
		_ = queue.Enqueue(2)
		el, _ := queue.Peek()

		assertLength(t, queue, 2)
		assertEqual(t, el, 15)
	})

	t.Run("Enqueue and dequeue bunch of elements", func(t *testing.T) {
		queue, _ := NewArrayQueue(2)

		assertLength(t, queue, 0)

		for index := 0; index < 10; index++ {
			_ = queue.Enqueue(index)
			_ = queue.Enqueue(index + 1)

			assertLength(t, queue, 2)

			el, _ := queue.Dequeue()
			assertEqual(t, el, index)

			el, _ = queue.Dequeue()
			assertEqual(t, el, index+1)

			assertLength(t, queue, 0)
		}
		assertLength(t, queue, 0)
	})

	t.Run("Enqueue more elements than a queue can store", func(t *testing.T) {
		queue, _ := NewArrayQueue(2)

		assertLength(t, queue, 0)

		_ = queue.Enqueue(15)
		_ = queue.Enqueue(2)
		err := queue.Enqueue(9)

		assertError(t, err, ErrorExceededCapacity)
		assertLength(t, queue, 2)
	})
}

func TestDequeue(t *testing.T) {
	t.Run("Dequeue takes an element from a front of a queue", func(t *testing.T) {
		queue, _ := NewArrayQueue(2)
		_ = queue.Enqueue(111)
		_ = queue.Enqueue(1337)

		el, _ := queue.Dequeue()
		assertEqual(t, el, 111)

		el, _ = queue.Dequeue()
		assertEqual(t, el, 1337)
	})

	t.Run("Dequeue from empty queue", func(t *testing.T) {
		queue, _ := NewArrayQueue(2)

		_, err := queue.Dequeue()

		assertError(t, err, ErrorEmptyQueue)
	})
}
