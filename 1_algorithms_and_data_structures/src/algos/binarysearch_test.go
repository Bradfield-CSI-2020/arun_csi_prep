package algos_test

import (
	"testing"

	"algos"
)


func TestBinarySearch(t *testing.T)  {

	newList := []int{1,2,3,4,5,6,7,8,9}

	found := algos.BinarySearch(newList, 10)

	if (found != -1) {
		t.Error("expected -1, got ", found)
	}

	found = algos.BinarySearch(newList, 2)

	if (found != 1) {
		t.Error("expected 1, got ", found)
	}

	found = algos.BinarySearch(newList, -1)

	if (found != -1) {
		t.Error("expected -1, got ", found)
	}

	anotherNewList := []int{}

	found = algos.BinarySearch(anotherNewList, 1)

	if (found != -1) {
		t.Error("expected -1, got ", found)
	}
}