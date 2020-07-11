package datastructures

// Queue is amazing
type Queue []string

// IsEmpty returns true only if queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Size returns the size of the queue
func (q *Queue) Size() int {
	return len(*q)
}

// Enqueue adds an item to the end of the queue
func (q *Queue) Enqueue(item string) {
	*q = append(*q, item)
}

// Dequeue returns from front of the queue
func (q *Queue) Dequeue() string {
	item := (*q)[0] 
	*q =  (*q)[1:]
	return item
}
