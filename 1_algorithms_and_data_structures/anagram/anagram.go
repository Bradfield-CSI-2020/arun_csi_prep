package main

import (
	"fmt"
	"strings"
	"sort"
)

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239}
const alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	fmt.Printf("anagramCheck: %v\n", anagramCheck("hello", "helol"))
	fmt.Printf("anagramCheck: %v\n", anagramCheck("bye", "Hello"))

	// could not run this word in node, the resulting number was too big for nodejs
	firstWord := "radiometeorograph"
	wordArray := strings.Split(firstWord, "")
	sort.Strings(wordArray)
	newWord := strings.Join(wordArray, "")
	
	fmt.Printf("anagramCheck: %v\n", anagramCheck(firstWord, newWord))

}

func anagramCheck(first string, second string) bool {
	if (len(first) != len(second)) {
		return false
	}

	var  primeValueFirst = 1
	var  primeValueSecond = 1

	primeValues := generateLetterPrimeValueMap()

	for _, value := range first {
		primeValueFirst = primeValueFirst * primeValues[string(value)]
	}

	for _, value := range second {
		primeValueSecond = primeValueSecond * primeValues[string(value)]
	}

	return primeValueFirst == primeValueSecond
}

func generateLetterPrimeValueMap() map[string]int {
	var valueMap = make(map[string]int)

	for i, value := range alphabets {
		valueMap[string(value)] = primes[i]
	}

	return valueMap
}