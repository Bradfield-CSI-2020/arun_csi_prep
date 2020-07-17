package datastructures_test

import (
	"testing"

	"datastructures"
)

func TestQueue(t *testing.T) {
	var newQueue datastructures.Queue

	if newQueue.IsEmpty() != true {
		t.Error("expected true, but got ", newQueue.IsEmpty())
	}

	newQueue.Enqueue("Arun")

	if newQueue.IsEmpty() != false {
		t.Error("expected false, but got ", newQueue.IsEmpty())
	}

	newQueue.Enqueue("Zack")

	if newQueue.Size() != 2 {
		t.Error("expected 2, but got ", newQueue.Size())
	}

	item, _ := newQueue.Dequeue()

	if item != "Arun" {
		t.Error("expected Arun, but got ", item)
	}

	item, _ = newQueue.Dequeue()

	if item != "Zack" {
		t.Error("expected Zack, but got ", item)
	}

	if newQueue.IsEmpty() != true {
		t.Error("expected true, but got ", newQueue.IsEmpty())
	}

	item, err := newQueue.Dequeue()

	if item != "" {
		t.Error("expected \"\", but got ", item)
	}

	if err == nil {
		t.Error("expected failure, but got ", err)
	}

}
