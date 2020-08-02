package datastructures

import (
	"fmt"
)

// ListNode is amazing
type ListNode struct {
	Value int
	Next  *ListNode
}

// List is amazing
type List struct {
	Head *ListNode
}

// IsEmpty returns true when empty
func (l *List) IsEmpty() bool {
	return l.Head == nil
}

// Add item to the head of the List
func (l *List) Add(value int) {
	node := &ListNode{
		Value: value,
		Next:  l.Head,
	}
	l.Head = node
}

// Size returns the size of the list
func (l *List) Size() int {

	curNode := l.Head
	count := 0

	for curNode != nil {
		count++
		curNode = curNode.Next
	}

	return count
}

// Search if an item is in the list
func (l *List) Search(value int) int {

	curNode := l.Head
	count := 0

	for curNode != nil {
		if curNode.Value == value {
			return count
		}
		curNode = curNode.Next
		count++
	}

	return -1
}

// Remove the first occurrence of a value from the list
func (l *List) Remove(value int) bool {

	cur := l.Head

	if cur == nil {
		return false
	}

	if cur.Value == value {
		l.Head = cur.Next
		return true
	}

	for cur.Next != nil {
		if cur.Next.Value == value {
			cur.Next = cur.Next.Next
			return true
		}
		cur = cur.Next
	}

	return false
}

// Insert item at the provided
func (l *List) Insert(value int, index int) bool {

	if index == 0 {
		l.Add(value)
		return true
	}

	var node ListNode
	node.Value = value
	prev := l.Head
	cur := prev.Next

	count := 1

	for cur != nil {
		if count == index {
			fmt.Println("Adding")
			node.Next = cur
			prev.Next = &node
			return true
		}
		prev = cur
		cur = prev.Next
		count++
	}

	return false
}
