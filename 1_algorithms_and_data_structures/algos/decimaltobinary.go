package algos

import (
	"strconv"

	"github.com/Bradfield-CSI-2020/arun_csi_prep/1_algorithms_and_data_structures/datastructures"
)

// DecimalToBinary provides a binary string rep of a decimal number
func DecimalToBinary(decimalNum int) string {

	var reminderStack datastructures.Stack

	for decimalNum > 0 {
		reminder := decimalNum % 2
		decimalNum = decimalNum / 2

		reminderStack.Push(strconv.Itoa(reminder))
	}

	var binaryRep = ""

	for reminderStack.Size() > 0 {
		binaryRep = binaryRep + reminderStack.Pop()
	}

	if binaryRep == "" {
		return "0"
	}

	return binaryRep
}
