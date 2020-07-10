package algos

import (
	"datastructures"
)

// CheckbalancedParans checks a parans string to see if its balanced
func CheckbalancedParans(parans string) bool {

	paranStack := new(datastructures.Stack)

	for _, vRune := range parans {
		value := string(vRune)

		if (value == "{") {
			paranStack.Push("{")
		} 

		if (value == "}") {
			paranStack.Pop()
		}
	}

	return paranStack.Size() == 0
}
