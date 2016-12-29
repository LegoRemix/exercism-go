// Package anagram finds whether or not there any anagrams of an input from a set of candidates
package anagram

import (
	"reflect"
	"strings"
	"unicode"
)

// prune takes a list of strings and returns the list minus duplicates and normalized into lower case
func prune(original string, dictionary []string) []string {
	wordSet := make(map[string]bool)
	for _, word := range dictionary {
		wordSet[strings.ToLower(word)] = true
	}

	lower := strings.ToLower(original)
	var pruned []string
	for k, _ := range wordSet {
		if k == lower {
			continue
		}
		pruned = append(pruned, k)
	}
	return pruned
}

// letterFrequency converts a word into a frequency map of runes
func letterFrequency(input string) map[rune]int {
	freqMap := make(map[rune]int)
	for _, r := range input {
		freqMap[unicode.ToLower(r)]++
	}
	return freqMap
}

// Detect finds any anagram of the original word in the list of candidates
func Detect(original string, candidates []string) []string {
	//prune our list of candidates
	prunedCandidates := prune(original, candidates)
	//get our expected letter frequency
	expected := letterFrequency(original)

	//anagrams is the list of anagrams we've found
	var anagrams []string

	for _, candid := range prunedCandidates {
		if reflect.DeepEqual(expected, letterFrequency(candid)) {
			anagrams = append(anagrams, candid)
		}
	}

	return anagrams
}
