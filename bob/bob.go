// Package bob emulates an annoying teen named bob
package bob

import (
	"unicode"
)

const testVersion = 2 // same as targetTestVersion

// Hey is a function which emulates bob's annoying speech patterns
func Hey(query string) string {

	//we need to keep track of whitespace
	isWhiteSpace := true
	//as well as if we are all caps
	isAllCaps := true
	//check if we have any letters
	hasLetters := false
	//and if we have an ending question mark
	hasEndingQMark := false

	for _, r := range query {
		hasLetters = hasLetters || unicode.IsLetter(r)

		if r == '?' || (unicode.IsSpace(r) && hasEndingQMark) {
			hasEndingQMark = true
		} else {
			hasEndingQMark = false
		}

		isAllCaps = isAllCaps && (!unicode.IsLetter(r) || unicode.IsUpper(r))
		isWhiteSpace = isWhiteSpace && unicode.IsSpace(r)

	}

	//if we yell, bob says to chill
	if isAllCaps && hasLetters {
		return "Whoa, chill out!"

		// if we have only white space, get angry
	} else if isWhiteSpace {
		return "Fine. Be that way!"

		// if we have a question, say sure
	} else if hasEndingQMark {
		return "Sure."

	} else {
		return "Whatever."
	}

}
