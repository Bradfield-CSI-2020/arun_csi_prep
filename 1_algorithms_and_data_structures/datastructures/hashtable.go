package datastructures

import "math"

const length = 100

// HashTable is ok
type HashTable struct {
	data [length]*GenericList
}

// ListData is ok
func hash(s string) int {
	h := 0
	for pos, char := range s {
		h += int(char) * int(math.Pow(31, float64(len(s)-pos+1)))
	}
	return h
}

func index(hash int) int {
	return hash % length
}

// Set is needed
func (h *HashTable) Set(key string, value interface{}) bool {
	index := index(hash(key))

	if h.data[index] == nil {
		var newList GenericList
		newList.Append(key, value)

		h.data[index] = &newList
	} else {
		h.data[index].Append(key, value)
	}

	return true
}

// Get gets
func (h *HashTable) Get(key string) interface{} {
	index := index(hash(key))

	if h.data[index] == nil {
		return nil
	}
	return h.data[index].Get(key)

}
