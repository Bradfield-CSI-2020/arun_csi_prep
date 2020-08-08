package algos

import (
	"testing"
)

func TestNumberToAnyBase(t *testing.T) {
	result := NumberToAnyBase(1000, 10)

	if result != "1000" {
		t.Error("expected 1000, got ", result)
	}

	// fullstack academy class number
	result = NumberToAnyBase(5, 2)

	if result != "101" {
		t.Error("expected 101, got ", result)
	}

	result = NumberToAnyBase(1453, 16)

	if result != "5ad" {
		t.Error("expected 5ad, got ", result)
	}
}
