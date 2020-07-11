package datastructures

// import ()

// Dequeue is amazing
type Dequeue struct {
	Queue
}

// EnqueueFront does something
func (d *Dequeue) EnqueueFront(item string) {
	var name Queue

	name.Enqueue(item)

	newQueue := append(name, d.Queue...)

	d.Queue = newQueue
}

// DequeueEnd does something
func (d *Dequeue) DequeueEnd() string {
	item := (d.Queue)[len(d.Queue) - 1]
	d.Queue = d.Queue[len(d.Queue) - 2:]

	return item
}