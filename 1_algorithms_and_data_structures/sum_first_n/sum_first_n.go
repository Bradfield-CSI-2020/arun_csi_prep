package main

import (
	"fmt"
)

func main() {
	fmt.Printf("sumOfFirstN for 10: %v\n", sumOfFirstN(10))
	fmt.Printf("sumOfFirstN for 10: %v\n", sumOfFirstN(100))
}

func sumOfFirstN(n int) int {
	return n * (n + 1) / 2
}