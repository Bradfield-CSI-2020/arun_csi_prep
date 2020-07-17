package algos

const charIntMap = "0123456789abcdef"

// NumberToAnyBase converts a base decimal number to anybase string representation
func NumberToAnyBase(num int, base int) string {

	if num < base {
		return string(charIntMap[num])
	}

	return NumberToAnyBase(num/base, base) + string(charIntMap[num%base])

}
