package algos_test

import (
	"testing"

	"algos"
)

func TestDecimalToBinary(t *testing.T) {
	v := algos.DecimalToBinary(233)

	if (v != "11101001") {
		t.Error("Expected 11101001, got", v)
	}

	v = algos.DecimalToBinary(0)

	if (v != "0") {
		t.Error("Expected 0, got", v)
	}
}