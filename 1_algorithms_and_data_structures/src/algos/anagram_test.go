package algos_test

import (
	"sort"
	"strings"
	"testing"

	"algos"
)

func TestAnagramCheck(t *testing.T) {
	v := algos.AnagramCheck("hello", "helol")

	if v != true {
		t.Error("Expected true, got", v)
	}

	v = algos.AnagramCheck("bye", "Hello")

	if v != false {
		t.Error("Expected false, got", v)
	}

	firstWord := "radiometeorograph"

	wordArray := strings.Split(firstWord, "")
	sort.Strings(wordArray)

	newWord := strings.Join(wordArray, "")

	v = algos.AnagramCheck(firstWord, newWord)

	if v != true {
		t.Error("Expected true, got", v)
	}
}
