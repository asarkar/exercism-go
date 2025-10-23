package linkedlist

import (
	"errors"
)

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
	Value any
	next  *Node
	prev  *Node
}

type List struct {
	head, tail *Node
}

var errEmpty = errors.New("empty list")

func NewList(elements ...any) *List {
	if len(elements) == 0 {
		return &List{}
	}
	head := &Node{}
	prev := head
	for _, v := range elements {
		if v == nil {
			panic("nil value is not permitted")
		}
		node := &Node{Value: v}
		prev.next = node
		node.prev = prev
		prev = node
	}
	if head.next != nil {
		head.next.prev = nil
	}
	return &List{
		head: head.next,
		tail: prev,
	}
}

func (node *Node) Next() *Node {
	if node == nil {
		return nil
	}
	return node.next
}

func (node *Node) Prev() *Node {
	if node == nil {
		return nil
	}
	return node.prev
}

func (list *List) Unshift(v any) {
	if list == nil {
		panic("empty list")
	}
	node := &Node{Value: v}
	node.next = list.head
	if list.head == nil {
		list.tail = node
	} else {
		list.head.prev = node
	}

	list.head = node
}

func (list *List) Shift() (any, error) {
	if list == nil || list.head == nil {
		return nil, errEmpty
	}
	node := list.head
	if list.head.next == nil {
		list.head = nil
		list.tail = nil
	} else {
		list.head.next.prev = nil
		list.head = list.head.next
	}
	node.prev = nil
	node.next = nil
	return node.Value, nil
}

func (list *List) Push(v any) {
	if list == nil {
		panic("empty list")
	}
	node := &Node{Value: v}
	node.prev = list.tail
	if list.tail == nil {
		list.head = node
	} else {
		list.tail.next = node
	}

	list.tail = node
}

func (list *List) Pop() (any, error) {
	if list == nil || list.head == nil {
		return nil, errEmpty
	}

	node := list.tail
	if list.tail.prev == nil {
		list.head = nil
		list.tail = nil
	} else {
		list.tail = list.tail.prev
		list.tail.next = nil
	}

	node.prev = nil
	node.next = nil
	return node.Value, nil
}

func (list *List) Reverse() {
	if list == nil || list.head == nil {
		return
	}
	prev := list.head
	node := list.head.next
	for node != nil {
		next := node.next
		node.next = prev
		prev.prev = node
		prev = node
		node = next
	}
	list.head.next = nil
	list.tail.prev = nil
	list.tail = list.head
	list.head = prev
}

func (list *List) First() *Node {
	return list.head
}

func (list *List) Last() *Node {
	return list.tail
}
