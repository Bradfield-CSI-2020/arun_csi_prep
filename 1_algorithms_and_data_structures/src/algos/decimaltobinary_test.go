package algos

import (
	"testing"
)

func TestDecimalToBinary(t *testing.T) {
	v := DecimalToBinary(233)

	if v != "11101001" {
		t.Error("Expected 11101001, got", v)
	}

	v = DecimalToBinary(0)

	if v != "0" {
		t.Error("Expected 0, got", v)
	}

	v = DecimalToBinary(42)

	if v != "101010" {
		t.Error("Expected 101010, got", v)
	}
}
