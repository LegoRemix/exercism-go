// Package pangram helps users determine if an input sentence is a pangram
package pangram

import (
	"unicode"
)

const testVersion = 1
const LETTERS string = "abcdefghijklmnopqrstuvwxyz"

// IsPangram checks if a sentence is a pangram
func IsPangram(word string) bool {
	alphaMap := make(map[rune]bool)

	for _, ch := range word {
		if unicode.IsLetter(ch) {
			alphaMap[unicode.ToLower(rune(ch))] = true
		}
	}

	for _, ch := range LETTERS {
		if !alphaMap[ch] {
			return false
		}
	}

	return true
}
