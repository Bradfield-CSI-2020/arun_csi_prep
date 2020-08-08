package algos

import (
	"sort"
	"strings"
	"testing"
)

func TestAnagramCheck(t *testing.T) {
	v := AnagramCheck("hello", "helol")

	if v != true {
		t.Error("Expected true, got", v)
	}

	v = AnagramCheck("bye", "Hello")

	if v != false {
		t.Error("Expected false, got", v)
	}

	firstWord := "radiometeorograph"

	wordArray := strings.Split(firstWord, "")
	sort.Strings(wordArray)

	newWord := strings.Join(wordArray, "")

	v = AnagramCheck(firstWord, newWord)

	if v != true {
		t.Error("Expected true, got", v)
	}
}
