// Package wordcount finds the frequency of words in a given string
package wordcount

import (
	"regexp"
	"strings"
)

// We are on the second version of the test suite
const testVersion = 2

// Frequency is a map from strings to their frequency of appearance in the input
type Frequency map[string]int

// WordCount counts the number of instances of a word in a given phrase
func WordCount(phrase string) Frequency {

	freqs := make(Frequency)
	//a word is an extent of alphabetic characters
	re := regexp.MustCompile("[[:alnum:]]+")
	//for each word we find ...
	for _, word := range re.FindAllString(phrase, -1) {
		//increment the frequency of the word (words are all lowercased)
		freqs[strings.ToLower(word)]++
	}
	return freqs
}
