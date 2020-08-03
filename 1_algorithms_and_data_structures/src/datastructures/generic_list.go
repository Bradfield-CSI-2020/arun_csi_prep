package datastructures

// GenericList is awesome
type GenericList struct {
	Head *Node
}

type ListData struct {
	key   string
	value interface{}
}

// Node is awesome
type Node struct {
	Next *Node
	Data ListData
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
func (ll *GenericList) Append(key string, value interface{}) bool {
	nextNode := &Node{
		Next: ll.Head,
		Data: ListData{key, value},
	}

	ll.Head = nextNode

	return true
}

// Get gets
func (ll *GenericList) Get(key string) interface{} {
	cur := ll.Head

	for cur != nil {
		if cur.Data.key == key {
			return cur.Data.value
		}
		cur = cur.Next
	}

	return nil
}

// Remove is awesome
func (ll *GenericList) Remove(key string) bool {
	cur := ll.Head

	if cur == nil {
		return false
	}

	if cur.Data.key == key {
		ll.Head = cur.Next
		return true
	}

	for cur.Next != nil {
		if cur.Next.Data.key == key {
			cur.Next = cur.Next.Next
			return true
		}
		cur = cur.Next
	}

	return false
}
