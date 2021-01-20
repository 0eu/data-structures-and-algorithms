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
