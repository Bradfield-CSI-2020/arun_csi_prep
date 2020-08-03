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

	removed := newList.Remove("")

	if removed != false {
		t.Error("expected false, got ", removed)
	}

	values := []int{1, 2, 3, 4, 5}
	keys := []string{"a", "b", "c", "d", "e"}

	for index, key := range keys {
		newList.Append(key, values[index])
	}

	size = newList.Size()

	if size != 5 {
		t.Error("expected size to be 5, got ", size)
	}

	removed = newList.Remove("d")

	if removed != true {
		t.Error("expected true, got ", removed)
	}

	removed = newList.Remove("e")

	if removed != true {
		t.Error("expected true, got ", removed)
	}

	removed = newList.Remove("e")

	if removed != false {
		t.Error("expected false, got ", removed)
	}

	got := newList.Get("a")

	if got != 1 {
		t.Error("expected 1, got ", got)
	}

	got = newList.Get("e")

	if got != nil {
		t.Error("expected nil, got ", got)
	}

}
