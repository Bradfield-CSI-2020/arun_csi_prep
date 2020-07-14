package datastructures

// import (
// 	"fmt"
// )

// OrderedListNode is awesome
// type OrderedListNode struct {
// 	Value int
// 	Next *OrderedListNode
// }

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

	cur := ol.Head

	if (cur == nil) {
		return false
	}

	if (cur.Value == value) {
		ol.Head = cur.Next
		return true
	}

	if (cur.Value > value) {
		return false
	}

	prev := cur
	cur = cur.Next

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