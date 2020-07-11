package algos

import (
	"datastructures"
)

// Hotpotato tells us who survies a game of hotpotato
func Hotpotato(names []string, num int) string {

	var newQueue datastructures.Queue

	for _, value := range names {
		newQueue.Enqueue(value)
	}

	for newQueue.Size() > 1 {

		for i:= 0; i < num; i++ {
			name := newQueue.Dequeue()
			newQueue.Enqueue(name)
		}

		newQueue.Dequeue() // dead
	}

	return newQueue.Dequeue()
}
