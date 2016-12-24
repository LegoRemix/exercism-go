// Package acronym handles abbreviating a set of words
package acronym

import "strings"
import "unicode"

// We have test version 1
const testVersion = 1

// processWord takes a single world and outputs the letters used to represent it in an acronym
func processWord(word string) string {
	//if a word contains a dash, we need to split it and process those parts
	if strings.Contains(word, "-") {
		letters := ""
		for _, w := range strings.Split(word, "-") {
			letters += processWord(w)
		}
		return letters
	}

	//if a word is all caps, then just use the first letter
	if strings.ToUpper(word) == word {
		return string(unicode.ToUpper(rune(word[0])))
	}

	//otherwise return all letters that are capitalized or the first one
	letters := ""
	for i, r := range word {
		if unicode.IsLetter(r) && (i == 0 || unicode.IsUpper(r)) {
			letters += string(unicode.ToUpper(rune(word[i])))
		}
	}

	return letters
}

// abbreviate takes a series of words and forms an abbreivation
func abbreviate(input string) string {
	var split []string = strings.Split(input, " ")

	acronym := ""
	for _, fragment := range split {
		acronym += processWord(fragment)
	}

	return acronym
}
