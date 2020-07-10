package sum_test

import (
	"testing"

	"sum"
)


func TestOfFirstN(t *testing.T) {
	v := sum.OfFirstN(10)

	if (v != 55) {
		t.Error("Expected 55, got", v)
	}

	v = sum.OfFirstN(100)

	if (v != 5050) {
		t.Error("Expected 5050, got", v)
	}
}