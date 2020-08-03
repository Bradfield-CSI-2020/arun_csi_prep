package algos

import (
	"testing"
)

func TestSumOfFirstN(t *testing.T) {
	v := SumOfFirstN(10)

	if v != 55 {
		t.Error("Expected 55, got", v)
	}

	v = SumOfFirstN(100)

	if v != 5050 {
		t.Error("Expected 5050, got", v)
	}
}
