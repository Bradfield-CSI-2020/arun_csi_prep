package stack

// import (
// 	"fmt"
// )

// func main() {

// 	var randomStack Stack

// 	fmt.Println("Stack is empty?" , randomStack.IsEmpty())

// 	randomStack.Push("Arun")
// 	fmt.Println("Stack is empty?" , randomStack.IsEmpty())

// 	randomStack.Push("Zack")
// 	fmt.Println("Stack size is?" , randomStack.Size())

// 	fmt.Println("item at the top is?" , randomStack.Pop())
// 	fmt.Println("item at the top is?" , randomStack.Pop())

// 	fmt.Println("Stack is empty?" , randomStack.IsEmpty())	

// }

// Stack is amazing
type Stack []string

// IsEmpty returns true only if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Size returns the size of the stack
func (s *Stack) Size() int {
	return len(*s)
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item string) {
	*s = append(*s, item)
}

// Pop returns from the top of the stack
func (s *Stack) Pop() string {
	item := (*s)[len(*s) - 1] 
	*s =  (*s)[:(len(*s) - 1)]
	return item
}