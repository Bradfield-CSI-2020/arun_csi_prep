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

	// Sum of first n numbers
	fmt.Println("Hello World")
	fmt.Printf("sumOfFirstN for 10: %v\n", sum.OfFirstN(10))
	fmt.Printf("sumOfFirstN for 10: %v\n", sum.OfFirstN(100))


	// Check 2 strings to see if they are anagrams of each other
	fmt.Printf("anagramCheck: %v\n", algos.AnagramCheck("hello", "helol"))
	fmt.Printf("anagramCheck: %v\n", algos.AnagramCheck("bye", "Hello"))
	// could not run this word in node, the resulting number was too big for nodejs
	firstWord := "radiometeorograph"
	wordArray := strings.Split(firstWord, "")
	sort.Strings(wordArray)
	newWord := strings.Join(wordArray, "")
	
	fmt.Printf("anagramCheck: %v\n", algos.AnagramCheck(firstWord, newWord))


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


	fmt.Println("Hello?", algos.CheckbalancedParans("{{}}"))

}