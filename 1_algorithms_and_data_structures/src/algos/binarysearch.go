package algos

// BinarySearch is better (sometimes) than Search
func BinarySearch(list []int, value int) int {

	if (len(list) == 0) {
		return -1
	}

	mid := len(list) / 2
	
	if (list[mid] == value) {
		return 1
	}

	if (list[mid] < value) {
		list = list[(mid+1):]
	} else {
		list = list[:mid]
	}

	if (len(list) == 0) {
		return -1
	}

	return BinarySearch(list, value)
}
