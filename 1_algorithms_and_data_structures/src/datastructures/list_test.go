package datastructures_test

import (
	"testing"

	"datastructures"
)

func TestListNode(t *testing.T)  {

	var newList datastructures.List

	isEmpty := newList.IsEmpty()
	
	if (isEmpty != true) {
		t.Error("expected true, got ", isEmpty)
	}

	size := newList.Size()

	if (size != 0) {
		t.Error("expected size to be 0, got ", size)
	}


	items := []int{1,2,3,4,5}

	for _, value := range items {
		newList.Add(value)
	}

	size = newList.Size()

	if (size != 5) {
		t.Error("expected size to be 5, got ", size)
	}

	isEmpty = newList.IsEmpty()
	
	if (isEmpty != false) {
		t.Error("expected false, got ", isEmpty)
	}

	cur := newList.Head
	expectedValue := 5

	for cur != nil {
		
		if (cur.Value != expectedValue) {
			t.Error("expected", expectedValue, "got", cur.Value)
		}

		expectedValue--
		cur = cur.Next
	}


	found := newList.Search(1)

	if (found != 4) {
		t.Error("expected 4, got ", found)
	}

	found = newList.Search(6)

	if (found != -1) {
		t.Error("expected -1, got ", found)
	}

	removed := newList.Remove(5)

	if (removed != true) {
		t.Error("expected true, got ", removed)
	}

	if (newList.Size() != 4) {
		t.Error("expected 4, got ", newList.Size())
	}

	removed = newList.Remove(5)

	if (removed != false) {
		t.Error("expected false, got ", removed)
	}


	added:= newList.Insert(7, 3)

	if (added != true) {
		t.Error("expected true, got ", added)
	}


	found = newList.Search(7)

	if (found != 3) {
		t.Error("expected 3, got ", found)
	}


	newList.Insert(8, 0)

	found = newList.Search(8)

	if (found != 0) {
		t.Error("expected 0, got ", found)
	}
	
	added = newList.Insert(8, 10)

	if (added != false) {
		t.Error("expected false, got ", added)
	}

}