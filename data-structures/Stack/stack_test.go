package stack

import "testing"

func assertLength(t *testing.T, stack Stack, expected int) {
	t.Helper()
	actual := stack.Size()
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

func TestPush(t *testing.T) {
	t.Run("Push should put an element onto a stack", func(t *testing.T) {
		stack := NewLinkedListStack(1)

		_ = stack.Push(1)

		assertLength(t, stack, 1)
	})

	t.Run("Push should exceeded a stack's capacity", func(t *testing.T) {
		stack := NewLinkedListStack(1)

		err := stack.Push(1)

		assertError(t, err, nil)

		err = stack.Push(1)

		assertError(t, err, ErrorExceededCapacity)
		assertLength(t, stack, 1)
	})
}

func TestPeek(t *testing.T) {
	t.Run("Peek should return first element", func(t *testing.T) {
		stack := NewLinkedListStack(5)
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		actual, err := stack.Peek()

		assertError(t, err, nil)
		assertEqual(t, actual, 3)
	})

	t.Run("Peek should throw an error on empty stack", func(t *testing.T) {
		stack := NewLinkedListStack(5)

		_, err := stack.Peek()

		assertError(t, err, ErrorEmptyStack)
	})
}

func TestPop(t *testing.T) {
	t.Run("Pop should return first element", func(t *testing.T) {
		stack := NewLinkedListStack(5)
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		actual, err := stack.Pop()

		assertError(t, err, nil)
		assertEqual(t, actual, 3)

		actual, err = stack.Pop()

		assertError(t, err, nil)
		assertEqual(t, actual, 2)
	})

	t.Run("Pop should throw an error on empty stack", func(t *testing.T) {
		stack := NewLinkedListStack(5)

		_, err := stack.Pop()

		assertError(t, err, ErrorEmptyStack)
	})
}
