package generics

// TypedList is a generic linked list that can hold values of any comparable type.
type TypedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// Node represents a node in the linked list, holding a value of type T and a pointer to the next node.
type Node[T comparable] struct {
	Value T
	next  *Node[T]
}

// Head returns the first node in the list.
func (l *TypedList[T]) Head() *Node[T] {
	return l.head
}

// Tail returns the last node in the list.
func (l *TypedList[T]) Tail() *Node[T] {
	return l.tail
}

// Size returns the number of nodes in the list.
func (l *TypedList[T]) Size() int {
	return l.size
}

// Append adds a new node with the given value to the end of the list.
func (l *TypedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

// Prepend adds a new node with the given value to the beginning of the list.
func (l *TypedList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
	l.size++
}

// First returns the first node in the list that matches the given node's value.
func (l *TypedList[T]) First(n *Node[T]) *Node[T] {
	if l.head == nil || n == nil {
		return nil
	}

	if l.head.Value == n.Value {
		return l.head
	}
	current := l.head
	for current.next != nil {
		if current.next.Value == n.Value {
			return current.next
		}
		current = current.next
	}
	return nil
}

// Last returns the last node in the list whose value matches the supplied value.
func (l *TypedList[T]) Last(v T) *Node[T] {
	if l.tail == nil {
		return nil
	}

	if l.tail.Value == v {
		return l.tail
	}

	var last *Node[T]

	for current := l.head; current != nil; current = current.next {
		if current.Value == v {
			last = current
		}
	}
	return last
}

// Remove removes the first occurrence of the given node from the list.
func (l *TypedList[T]) Remove(n *Node[T]) *TypedList[T] {
	if l.head == nil || n == nil {
		return l
	}

	if l.head.Value == n.Value {
		l.head = l.head.next
		if l.head == nil {
			l.tail = nil
		}
		l.size--
		return l
	}

	current := l.head
	for current.next != nil {
		if current.next.Value == n.Value {
			current.next = current.next.next
			if current.next == nil {
				l.tail = current
			}
			l.size--
			return l
		}
		current = current.next
	}
	return l
}
