package main

import (
	"fmt"
	"strings"
	"sort"

	"sum"
	"datastructures"
	"algos"
)

func main() {

	fmt.Println("Hello World")


	// Create a stack and play with it
	var newStack datastructures.Stack
	
	fmt.Println("Stack is empty?" , newStack.IsEmpty())

	newStack.Push("Arun")
	
	fmt.Println("Stack is empty?" , newStack.IsEmpty())

	newStack.Push("Zack")

	fmt.Println("Stack size is?" , newStack.Size())
	fmt.Println("item at the top is?" , newStack.Pop())
	fmt.Println("item at the top is?" , newStack.Pop())
	fmt.Println("Stack is empty?" , newStack.IsEmpty())	

}
