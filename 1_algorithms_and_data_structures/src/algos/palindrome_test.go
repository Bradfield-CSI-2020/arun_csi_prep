package algos_test

import (
	"testing"

	"algos"
)

func TestIsPalindrome(t *testing.T) {

	word := "asdffdsa"

	v := algos.IsPalindrome(word)

	if v != true {
		t.Error("expected true, but got ", v)
	}

	word = "asdffdsas"

	v = algos.IsPalindrome(word)

	if v != false {
		t.Error("expected false, but got ", v)
	}

	word = "asdfffdsa"
	v = algos.IsPalindrome(word)

	if v != true {
		t.Error("expected true, but got ", v)
	}
}
