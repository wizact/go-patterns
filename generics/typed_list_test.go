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
