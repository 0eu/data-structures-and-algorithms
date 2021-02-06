package LinkedList

import "testing"

func assertLength(t *testing.T, list *LinkedList, expected int) {
	t.Helper()
	actual := list.Size()
	if actual != expected {
		t.Errorf("expected length %d, but got: %d", expected, actual)
	}
}

func assertEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected error %s, but got: %s", expected, actual)
	}
}

func TestNewLinkedList(t *testing.T) {
	t.Run("Build a linked list", func(t *testing.T) {
		list := NewLinkedList()

		assertLength(t, list, 0)
	})
}

func TestLinkedList_Append(t *testing.T) {
	t.Run("Append adds elements to the end of a list", func(t *testing.T) {
		list := NewLinkedList()

		list.Append(10)
		list.Append(12)
		list.Append(13)
		first, _ := list.PeekFirst()
		last, _ := list.PeekLast()

		assertLength(t, list, 3)
		assertEqual(t, first, 10)
		assertEqual(t, last, 13)
	})
}

func TestLinkedList_Prepend(t *testing.T) {
	t.Run("Prepend adds elements to the front of a list", func(t *testing.T) {
		list := NewLinkedList()

		list.Prepend(10)
		list.Append(12)
		list.Prepend(13)
		first, _ := list.PeekFirst()
		last, _ := list.PeekLast()

		assertLength(t, list, 3)
		assertEqual(t, first, 13)
		assertEqual(t, last, 12)
	})
}

func TestLinkedList_Count(t *testing.T) {
	list := NewLinkedList()
	list.Append(10)
	list.Append(12)
	list.Append(12)
	list.Append(12)
	list.Append(12)
	list.Append(13)

	count := list.Count(12)

	assertEqual(t, count, 4)
}
