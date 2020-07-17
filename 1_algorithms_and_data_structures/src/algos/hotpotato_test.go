package algos

import (
	"testing"
)

func TestHotpotato(t *testing.T) {
	v := Hotpotato([]string{"Bill", "David", "Susan", "Jane", "Kent", "Brad"}, 9)

	if v != "David" {
		t.Error("Expected David, got", v)
	}
}
