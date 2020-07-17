package datastructures_test

import (
	"testing"

	"datastructures"
)

func TestStack(t *testing.T) {
	var newStack datastructures.Stack

	if newStack.IsEmpty() != true {
		t.Error("expected true, but got ", newStack.IsEmpty())
	}

	newStack.Push("Arun")

	if newStack.IsEmpty() != false {
		t.Error("expected false, but got ", newStack.IsEmpty())
	}

	newStack.Push("Zack")

	if newStack.Size() != 2 {
		t.Error("expected 2, but got ", newStack.Size())
	}

	topItem := newStack.Pop()

	if topItem != "Zack" {
		t.Error("expected Zack, but got ", topItem)
	}

	topItem = newStack.Pop()

	if topItem != "Arun" {
		t.Error("expected Arun, but got ", topItem)
	}

	if newStack.IsEmpty() != true {
		t.Error("expected true, but got ", newStack.IsEmpty())
	}
}
