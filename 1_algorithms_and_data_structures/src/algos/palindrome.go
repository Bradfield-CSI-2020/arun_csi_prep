package algos

import (
	"fmt"

	"datastructures"
)

// IsPalindrome returns true if the word is a palindrome
func IsPalindrome(word string) bool {

	var dequeue datastructures.Dequeue

	for _, value:= range word {
		dequeue.Enqueue(string(value))
	}

	for dequeue.Size() > 0 {

		fmt.Println("size", dequeue.Size())
		first, _ := dequeue.Dequeue()
		second, _ := dequeue.DequeueEnd()


		if (first != second) {
			return false
		}
	}

	return true
}