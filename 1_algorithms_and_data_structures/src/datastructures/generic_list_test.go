package datastructures

import (
	"testing"
)

func TestGenericList(t *testing.T) {

	var newList GenericList

	isEmpty := newList.IsEmpty()

	if isEmpty != true {
		t.Error("expected true, got ", isEmpty)
	}

	size := newList.Size()

	if size != 0 {
		t.Error("expected size to be 0, got ", size)
	}

	removed := newList.Remove(5)

	if removed != false {
		t.Error("expected false, got ", removed)
	}

	items := []int{1, 2, 3, 4, 5}

	for _, value := range items {
		newList.Append(value)
	}

	size = newList.Size()

	if size != 5 {
		t.Error("expected size to be 5, got ", size)
	}

	removed = newList.Remove(4)

	if removed != true {
		t.Error("expected true, got ", removed)
	}

	removed = newList.Remove(5)

	if removed != true {
		t.Error("expected true, got ", removed)
	}

	removed = newList.Remove(5)

	if removed != false {
		t.Error("expected false, got ", removed)
	}

}
