// Package scrabble computes the scrabble score for a given word
package scrabble

import "unicode"

// the test version is 4 for scrabble scoring
const testVersion = 4

// the lookup map for letters to scrabble scores
var scrabbleScores = map[rune]int{
	'a': 1,
	'b': 3,
	'c': 3,
	'd': 2,
	'e': 1,
	'f': 4,
	'g': 2,
	'h': 4,
	'i': 1,
	'j': 8,
	'k': 5,
	'l': 1,
	'm': 3,
	'n': 1,
	'o': 1,
	'p': 3,
	'q': 10,
	'r': 1,
	's': 1,
	't': 1,
	'u': 1,
	'v': 4,
	'w': 4,
	'x': 8,
	'y': 4,
	'z': 10,
}

//Score computes the scrabble point value of a given string
func Score(input string) (score int) {
	for _, r := range input {
		score += scrabbleScores[unicode.ToLower(r)]
	}
	return
}
