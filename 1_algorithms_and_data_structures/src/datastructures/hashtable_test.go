package datastructures

import "testing"

func TestHashTable(t *testing.T) {
	var newHashTable HashTable

	got := newHashTable.Get("a")

	if got != nil {
		t.Error("expected nil, got ", got)
	}

	doneSet := newHashTable.Set("a", 1)

	if doneSet != true {
		t.Error("expected true, got ", doneSet)
	}

	got = newHashTable.Get("a")

	if got != 1 {
		t.Error("expected 1, got ", got)
	}

	newHashTable.Set("b", 2)
	newHashTable.Set("c", 3)
	newHashTable.Set("d", 4)
	newHashTable.Set("e", 5)

	got = newHashTable.Get("b")

	if got != 2 {
		t.Error("expected 2, got ", got)
	}

	got = newHashTable.Get("c")

	if got != 3 {
		t.Error("expected 3, got ", got)
	}

	got = newHashTable.Get("d")

	if got != 4 {
		t.Error("expected 4, got ", got)
	}

	got = newHashTable.Get("e")

	if got != 5 {
		t.Error("expected 5, got ", got)
	}

	got = newHashTable.Get("f")

	if got != nil {
		t.Error("expected nil, got ", got)
	}

}

// todo: set should overwrite an existing key
// OR todo: remove should remove all instances of a key
