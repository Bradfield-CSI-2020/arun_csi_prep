package algos_test

import (
	"testing"

	"algos"
)

func TestHotpotato(t *testing.T) {
	v := algos.Hotpotato([]string{"Bill", "David", "Susan", "Jane", "Kent", "Brad"}, 9)

	if (v != "David") {
		t.Error("Expected David, got", v)
	}
}