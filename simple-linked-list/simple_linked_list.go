package linkedlist

import (
	"errors"
	"slices"
)

// Define the List and Element types here.
type Node struct {
	Value int
	next  *Node
}

type List struct {
	head *Node
}

func New(elements []int) *List {
	list := &List{
		head: nil,
	}
	for _, v := range elements {
		list.Push(v)
	}
	return list
}

// We could've kept a `size` field in the `List`, but
// calculating size of a singly linked list is an O(n) operation
// in functional languages, so, we decided to be consistent.
func (list *List) Size() int {
	node := list.head
	var count int
	for node != nil {
		count++
		node = node.next
	}
	return count
}

// Tests expect push to append an element at the end.
// Instead of traversing the list for every push operation,
// we prepend the element at the front.
// If the list is converted to an array, we reverse the array
// to match the expectation in the tests.
func (list *List) Push(element int) {
	node := &Node{Value: element}
	node.next = list.head
	list.head = node
}

// Since push happens at the front, we must pop from the front too.
func (list *List) Pop() (int, error) {
	if list.head == nil {
		return 0, errors.New("pop from empty list")
	}
	node := list.head
	list.head = list.head.next
	return node.Value, nil
}

// See the comment on `Push` about the order of elements in the array.
func (list *List) Array() []int {
	node := list.head
	var arr []int
	for node != nil {
		arr = append(arr, node.Value)
		node = node.next
	}
	slices.Reverse(arr)
	return arr
}

func (list *List) Reverse() *List {
	var prev *Node
	node := list.head
	for node != nil {
		next := node.next
		node.next = prev
		prev = node
		node = next
	}
	return &List{
		head: prev,
	}
}
