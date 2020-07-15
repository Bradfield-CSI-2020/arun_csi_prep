package algos_test


import (
	"testing"

	"algos"
)

func TestNumberToAnyBase(t *testing.T) {
	result := algos.NumberToAnyBase(1000, 10)

	if (result != "1000") {
		t.Error("expected 1000, got ", result)
	}

	// fullstack academy class number
	result = algos.NumberToAnyBase(5, 2)

	if (result != "101") {
		t.Error("expected 101, got ", result)
	}

	result = algos.NumberToAnyBase(1453, 16)

	if (result != "5ad") {
		t.Error("expected 5ad, got ", result)
	}
}