package stack

import "testing"

func assertLength(t *testing.T, stack Stack, expected int) {
	t.Helper()
	actual := stack.Size()
	if actual != expected {
		t.Errorf("expected length %d, but got: %d", expected, actual)
	}
}

func TestPush(t *testing.T) {
	t.Run("should push an element onto a stack", func(t *testing.T) {
		stack := NewLinkedListStack()

		stack.Push(1)

		assertLength(t, stack, 1)
	})
}
