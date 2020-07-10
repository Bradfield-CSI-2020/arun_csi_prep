package sumton

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Printf("sumOfFirstN for 10: %v\n", sumOfFirstN(10))
// 	fmt.Printf("sumOfFirstN for 10: %v\n", sumOfFirstN(100))
// }

// SumOfFirstN returns the sum of the first n natural numbers
func SumOfFirstN(n int) int {
	return n * (n + 1) / 2
}