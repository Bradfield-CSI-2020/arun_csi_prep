package datastructures_test

import (
	"testing"

	"datastructures"
)

func TestDequeue(t *testing.T) {
	var newDeQueue datastructures.Dequeue

	if (newDeQueue.IsEmpty() != true) {
		t.Error("expected true, but got ", newDeQueue.IsEmpty())
	}

	newDeQueue.Enqueue("Arun")

	if (newDeQueue.IsEmpty() != false) {
		t.Error("expected false, but got ", newDeQueue.IsEmpty())
	}

	newDeQueue.Enqueue("Zack")

	if (newDeQueue.Size() != 2) {
		t.Error("expected 2, but got ", newDeQueue.Size())
	}

	item := newDeQueue.Dequeue()

	if (item != "Arun") {
		t.Error("expected Arun, but got ", item)
	}

	item = newDeQueue.Dequeue()

	if (item != "Zack") {
		t.Error("expected Zack, but got ", item)
	}

	if (newDeQueue.IsEmpty() != true) {
		t.Error("expected true, but got ", newDeQueue.IsEmpty())
	}

	newDeQueue.EnqueueFront("ArunAgain")	

	newDeQueue.EnqueueFront("ZackAgain")


	item = newDeQueue.Dequeue()

	if (item != "ZackAgain") {
		t.Error("expected ZackAgain, but got ", item)
	}

	newDeQueue.EnqueueFront("ZackAgain")

	item = newDeQueue.DequeueEnd()

	if (item != "ArunAgain") {
		t.Error("expected ArunAgain, but got ", item)
	}

}
