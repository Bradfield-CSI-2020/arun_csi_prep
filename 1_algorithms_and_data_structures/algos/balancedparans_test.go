package algos

import (
	"testing"
)

func TestCheckbalancedParans(t *testing.T) {
	v := CheckbalancedParans("{{}}")

	if v != true {
		t.Error("Expected true, got", v)
	}

	v = CheckbalancedParans("[[[]]][")

	if v != false {
		t.Error("Expected false, got", v)
	}

	v = CheckbalancedParans("[[{{(())}}]]")

	if v != true {
		t.Error("Expected true, got", v)
	}

	v = CheckbalancedParans("[[{{((())}}]]")

	if v != false {
		t.Error("Expected false, got", v)
	}
}
