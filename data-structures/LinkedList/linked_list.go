package LinkedList

// Node is a unit in a linked list, that stores data and a pointer to the next element.
type Node struct {
	Value interface{}
	Next  *Node
}

// LinkedList is a struct that has pointers to the first and last nodes and also size of a list.
type LinkedList struct {
	Head *Node
	Tail *Node
	size int
}

// NewLinkedList builds an empty linked list and returns its pointer.
func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
		size: 0,
	}
}

// Size returns the number of nodes in a list.
func (l *LinkedList) Size() int {
	return l.size
}

// Append adds an element to the end of a list.
func (l *LinkedList) Append(value interface{}) {
	newNode := &Node{Value: value, Next: nil}
	if l.Head != nil {
		l.Tail.Next, l.Tail = newNode, newNode
	} else {
		l.Head, l.Tail = newNode, newNode
	}
	l.size++
}
