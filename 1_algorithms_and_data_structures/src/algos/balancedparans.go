package algos

import (
	"datastructures"
)

// CheckbalancedParans checks a parans string to see if its balanced
func CheckbalancedParans(parans string) bool {

	paranStack := new(datastructures.Stack)

	frontParansMap := make(map[string]string)

	reverseParansMap := make(map[string]string)

	frontParansMap["["] = "]"
	frontParansMap["{"] = "}"
	frontParansMap["("] = ")"

	reverseParansMap["]"] = "["
	reverseParansMap["}"] = "{"
	reverseParansMap[")"] = "("



	for _, vRune := range parans {
		value := string(vRune)

		if (frontParansMap[value] != "") {
			paranStack.Push(value)
		} else if (reverseParansMap[value] != "") {
			openParan := paranStack.Pop()
			if (openParan != reverseParansMap[value]) {
				return false
			}
		}
	}

	return paranStack.Size() == 0
}
