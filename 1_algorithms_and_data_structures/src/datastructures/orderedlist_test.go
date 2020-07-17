package datastructures_test

import (
	"testing"

	"datastructures"
)

func TestOrderedList(t *testing.T) {

	var newList datastructures.OrderedList

	isEmpty := newList.IsEmpty()

	if isEmpty != true {
		t.Error("expected true, got ", isEmpty)
	}

	size := newList.Size()

	if size != 0 {
		t.Error("expected size to be 0, got ", size)
	}

	items := []int{5, 2, 3, 1, 4}

	for _, value := range items {
		newList.Add(value)
	}

	size = newList.Size()

	if size != 5 {
		t.Error("expected size to be 5, got ", size)
	}

	isEmpty = newList.IsEmpty()

	if isEmpty != false {
		t.Error("expected false, got ", isEmpty)
	}

	cur := newList.Head
	expectedValue := 1

	for cur != nil {

		if cur.Value != expectedValue {
			t.Error("expected", expectedValue, "got", cur.Value)
		}

		expectedValue++
		cur = cur.Next
	}

	found := newList.Search(0)

	if found != -1 {
		t.Error("expected 0, got ", found)
	}

	found = newList.Search(1)

	if found != 0 {
		t.Error("expected 0, got ", found)
	}

	found = newList.Search(6)

	if found != -1 {
		t.Error("expected -1, got ", found)
	}

	removed := newList.Remove(4)

	if removed != true {
		t.Error("expected true, got ", removed)
	}

	if newList.Size() != 4 {
		t.Error("expected 4, got ", newList.Size())
	}

}
