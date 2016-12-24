//Package isogram allows users to test whether or not a word is an isogram (doesn't have duplicate letters)
package isogram

import "unicode"

// we are on the 1st iteration of the testing suite
const testVersion = 1

//IsIsogram returns true iff the function has no repeated letters
func IsIsogram(text string) bool {
	letters := make(map[rune]bool)

	for _, r := range text {
		r = unicode.ToLower(r)
		if unicode.IsLetter(r) && letters[r] {
			return false
		}
		letters[r] = true
	}

	return true
}
