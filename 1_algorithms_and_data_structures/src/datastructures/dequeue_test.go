package datastructures

import (
	"testing"
)

func TestDequeue(t *testing.T) {
	var newDeQueue Dequeue

	if newDeQueue.IsEmpty() != true {
		t.Error("expected true, but got ", newDeQueue.IsEmpty())
	}

	newDeQueue.Enqueue("Arun")

	if newDeQueue.IsEmpty() != false {
		t.Error("expected false, but got ", newDeQueue.IsEmpty())
	}

	newDeQueue.Enqueue("Zack")

	if newDeQueue.Size() != 2 {
		t.Error("expected 2, but got ", newDeQueue.Size())
	}

	item, _ := newDeQueue.Dequeue()

	if item != "Arun" {
		t.Error("expected Arun, but got ", item)
	}

	item, _ = newDeQueue.Dequeue()

	if item != "Zack" {
		t.Error("expected Zack, but got ", item)
	}

	if newDeQueue.IsEmpty() != true {
		t.Error("expected true, but got ", newDeQueue.IsEmpty())
	}

	newDeQueue.EnqueueFront("ArunAgain")

	newDeQueue.EnqueueFront("ZackAgain")

	item, _ = newDeQueue.Dequeue()

	if item != "ZackAgain" {
		t.Error("expected ZackAgain, but got ", item)
	}

	newDeQueue.EnqueueFront("ZackAgain")

	item, _ = newDeQueue.DequeueEnd()

	if item != "ArunAgain" {
		t.Error("expected ArunAgain, but got ", item)
	}

	item, _ = newDeQueue.DequeueEnd()

	if item != "ZackAgain" {
		t.Error("expected ZackAgain, but got ", item)
	}

	if newDeQueue.Size() != 0 {
		t.Error("expected 0, but got ", newDeQueue.Size())
	}

	item, err := newDeQueue.DequeueEnd()

	if item != "" {
		t.Error("expected \"\", but got ", item)
	}

	if err == nil {
		t.Error("expected failure, but got ", err)
	}

}
