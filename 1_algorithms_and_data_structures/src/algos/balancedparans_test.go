package algos_test

import (
	"testing"

	"algos"
)

func TestCheckbalancedParans(t *testing.T) {
	v := algos.CheckbalancedParans("{{}}")

	if (v != true) {
		t.Error("Expected true, got", v)
	}

	v = algos.CheckbalancedParans("[[[]]][")

	if (v != false) {
		t.Error("Expected false, got", v)
	}
	
	v = algos.CheckbalancedParans("[[{{(())}}]]")

	if (v != true) {
		t.Error("Expected true, got", v)
	}

	v = algos.CheckbalancedParans("[[{{((())}}]]")

	if (v != true) {
		t.Error("Expected true, got", v)
	}
}