package algos

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	word := "asdffdsa"

	v := IsPalindrome(word)

	if v != true {
		t.Error("expected true, but got ", v)
	}

	word = "asdffdsas"

	v = IsPalindrome(word)

	if v != false {
		t.Error("expected false, but got ", v)
	}

	word = "asdfffdsa"
	v = IsPalindrome(word)

	if v != true {
		t.Error("expected true, but got ", v)
	}
}
