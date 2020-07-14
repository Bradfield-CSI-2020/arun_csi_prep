package datastructures

// import (
// 	"fmt"
// )

// OrderedList is awesome
type OrderedList struct {
	List
}

// Add item to the head of the List
func (ol *OrderedList) Add(value int) {
	var node ListNode
	node.Value = value

	prev := ol.Head

	if (prev == nil) {
		ol.Head = &node
		return
	}

	if (prev.Value >= value) {
		node.Next = prev
		ol.Head = &node
		return
	}

	for prev.Value < value {
		cur := prev.Next

		if (cur == nil) {
			prev.Next = &node
			return
		}

		if (cur.Value < value) {
			prev = cur
			cur = cur.Next
			continue
		}

		if (cur.Value > value) {
			prev.Next = &node
			node.Next = cur
			return
		}
	}

}

// Search if an item is in the list
func (ol *OrderedList) Search(value int) int {
	curNode := ol.Head
	var count int

	for curNode != nil {

		if (curNode.Value > value) {
			return -1
		}

		if (curNode.Value == value) {
			return count
		}
		curNode = curNode.Next
		count++
	}

	return -1
}

// Remove the first occurrence of a value from the list
func (ol *OrderedList) Remove(value int) bool {

	prev := ol.Head

	if (prev == nil) {
		return false
	}

	if (prev.Value == value) {
		ol.Head = prev.Next
		return true
	}

	if (prev.Value > value) {
		return false
	}

	cur := prev.Next

	for cur != nil {
		if (cur.Value == value) {
			prev.Next = cur.Next
			return true
		}

		prev = cur
		cur = cur.Next
	}

	return false
}