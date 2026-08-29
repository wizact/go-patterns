package generics

import (
	"testing"
)

func TestTypedList(t *testing.T) {
	list := &TypedList[int]{}

	list.Append(1)
	list.Append(2)
	list.Prepend(0)

	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}

	if list.head.Value != 0 {
		t.Errorf("Expected head Value 0, got %d", list.head.Value)
	}

	if list.tail.Value != 2 {
		t.Errorf("Expected tail Value 2, got %d", list.tail.Value)
	}
}

func TestTypedListFirstInt(t *testing.T) {
	list := &TypedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	node := &Node[int]{Value: 2}
	foundNode := list.First(node)

	if foundNode == nil || foundNode.Value != 2 {
		t.Errorf("Expected to find node with Value 2, got %v", foundNode)
	}

	nodeNotFound := &Node[int]{Value: 4}
	notFound := list.First(nodeNotFound)

	if notFound != nil {
		t.Errorf("Expected to not find node with Value 4, got %v", notFound)
	}
}

func TestTypedListFirstString(t *testing.T) {
	list := &TypedList[string]{}
	list.Append("a")
	list.Append("b")
	list.Append("c")

	node := &Node[string]{Value: "b"}
	foundNode := list.First(node)

	if foundNode == nil || foundNode.Value != "b" {
		t.Errorf("Expected to find node with Value 'b', got %v", foundNode)
	}

	nodeNotFound := &Node[string]{Value: "d"}
	notFound := list.First(nodeNotFound)

	if notFound != nil {
		t.Errorf("Expected to not find node with Value 'd', got %v", notFound)
	}
}

func TestTypedListLast_DuplicateValue_ReturnsLastMatchingNode(t *testing.T) {
	list := &TypedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(2)
	list.Append(3)
	expected := list.head.next.next

	actual := list.Last(2)

	if actual != expected {
		t.Errorf("expected last matching node %p, got %p", expected, actual)
	}
}

func TestTypedListLast_TailValue_ReturnsTail(t *testing.T) {
	list := &TypedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	actual := list.Last(3)

	if actual != list.tail {
		t.Errorf("expected tail node %p, got %p", list.tail, actual)
	}
}

func TestTypedListLast_MissingValue_ReturnsNil(t *testing.T) {
	list := &TypedList[int]{}
	list.Append(1)
	list.Append(2)

	actual := list.Last(3)

	if actual != nil {
		t.Errorf("expected nil, got %p", actual)
	}
}

func TestTypedListLast_EmptyList_ReturnsNil(t *testing.T) {
	list := &TypedList[int]{}

	actual := list.Last(1)

	if actual != nil {
		t.Errorf("expected nil, got %p", actual)
	}
}

func TestTypedListRemove(t *testing.T) {
	list := &TypedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	nodeToRemove := &Node[int]{Value: 2}
	list = list.Remove(nodeToRemove)

	if list.Size() != 2 {
		t.Errorf("Expected size 2 after removal, got %d", list.Size())
	}

	notFound := list.First(nodeToRemove)
	if notFound != nil {
		t.Errorf("Expected to not find node with Value 2 after removal, got %v", notFound)
	}
}

func TestTypedList_Map(t *testing.T) {
	list := &TypedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	f := func(x int) int {
		return x * 2
	}

	mappedList := list.Map(f)

	if mappedList.Size() != 3 {
		t.Errorf("Expected mapped list size 3, got %d", mappedList.Size())
	}

	expectedValues := []int{2, 4, 6}
	current := mappedList.head
	for i, expected := range expectedValues {
		if current == nil || current.Value != expected {
			t.Errorf("Expected value %d at index %d, got %v", expected, i, current)
		}
		current = current.next
	}
}
