package datastructures

// GenericList is awesome
type GenericList struct {
	Head *Node
}

// Node is awesome
type Node struct {
	Next *Node
	Data interface{}
}

// IsEmpty returns true when empty
func (ll *GenericList) IsEmpty() bool {
	return ll.Head == nil
}

// Size returns the number of items in the list
func (ll *GenericList) Size() int {

	curNode := ll.Head
	count := 0

	for curNode != nil {
		count++
		curNode = curNode.Next
	}

	return count
}

// Append is awesome
func (ll *GenericList) Append(d interface{}) *GenericList {
	nextNode := &Node{
		Next: ll.Head,
		Data: d,
	}

	ll.Head = nextNode

	return ll
}

// Remove is awesome
func (ll *GenericList) Remove(v interface{}) bool {
	cur := ll.Head

	if cur == nil {
		return false
	}

	if cur.Data == v {
		ll.Head = cur.Next
		return true
	}

	for cur.Next != nil {
		if cur.Next.Data == v {
			cur.Next = cur.Next.Next
			return true
		}
		cur = cur.Next
	}

	return false
}
