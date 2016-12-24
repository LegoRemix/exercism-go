// Package house allows a user to generate the lyrics to the song "The House That Jack Built"
package house

import (
	"fmt"
	"strings"
)

//lyrics to the song "The House that Jack Built".
var lyrics = []string{
	"the malt\nthat lay in",
	"the rat\nthat ate",
	"the cat\nthat killed",
	"the dog\nthat worried",
	"the cow with the crumpled horn\nthat tossed",
	"the maiden all forlorn\nthat milked",
	"the man all tattered and torn\nthat kissed",
	"the priest all shaven and shorn\nthat married",
	"the rooster that crowed in the morn\nthat woke",
	"the farmer sowing his corn\nthat kept",
	"the horse and the hound and the horn\nthat belonged to",
}

// Embed takes a relative phrase and embeds it in a noun phrase
func Embed(relPhrase, nounPhrase string) string {
	return fmt.Sprintf("%s %s", relPhrase, nounPhrase)
}

// Verse generates a verse of a song with relative clause given a recursive structure
func Verse(subject string, relPhrases []string, nounPhrase string) string {
	if len(relPhrases) <= 0 {
		return Embed(subject, nounPhrase)
	}
	return Verse(Embed(subject, relPhrases[0]), relPhrases[1:], nounPhrase)
}

// Song returns the full text of the song
func Song() string {
	firstLine := "This is"
	lastLine := "the house that Jack built."
	var verses = make([]string, 0)
	for v := 0; v <= len(lyrics); v++ {
		verse := Verse(firstLine, reverse(lyrics[:v]), lastLine)
		verses = append(verses, verse)

	}
	return strings.Join(verses, "\n\n")

}

// reverse copies and reverses the contents of our array
func reverse(items []string) []string {
	l := len(items)
	var reversed = make([]string, l)
	for i, v := range items {
		reversed[l-i-1] = v

	}
	return reversed

}
