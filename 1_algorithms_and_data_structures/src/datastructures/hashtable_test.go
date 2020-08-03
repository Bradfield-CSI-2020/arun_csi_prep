package datastructures

import "testing"

func TestHashTable(t *testing.T) {
	var newHashTable HashTable

	doneSet := newHashTable.Set("a", 1)

	if doneSet != true {
		t.Error("expected true, got ", doneSet)
	}

	got := newHashTable.Get("a")

	newHashTable.Set("b", 2)
	newHashTable.Set("c", 3)
	newHashTable.Set("d", 4)
	newHashTable.Set("e", 5)

	got = newHashTable.Get("d")

	if got != 4 {
		t.Error("expected 4, got ", got)
	}
}

// todo: set should overwrite an existing key
// OR todo: remove should remove all instances of a key
