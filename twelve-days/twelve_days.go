// Package twelve gives users the lyrics to the "Twelve Days of Christmas" song
package twelve

import (
	"fmt"
	"strings"
)

// ordinals is a listing of ordinal numbers
var ordinals = []string{
	"zeroth",
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

// cardinals is a listing of cardinal numbers
var cardinals = [13]string{
	"zero",
	"a",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
}

// gifts is a listing of the gifts for the holiday
var gifts = [13]string{
	"",
	"Partridge in a Pear Tree.",
	"Turtle Doves",
	"French Hens",
	"Calling Birds",
	"Gold Rings",
	"Geese-a-Laying",
	"Swans-a-Swimming",
	"Maids-a-Milking",
	"Ladies Dancing",
	"Lords-a-Leaping",
	"Pipers Piping",
	"Drummers Drumming",
}

// introFormat is the introductory part of the verse
const introFormat = "On the %s day of Christmas my true love gave to me"

// testVersion is only 1
const testVersion = 1

// Verse gives the user back the Nth line of the song "Twelve Days of Christmas"
func Verse(n int) string {
	//bounds check in put
	if n <= 0 || n > 12 {
		return ""
	}

	//is our input one?
	isOne := n == 1

	verseParts := make([]string, 0, n+2)
	verseParts = append(verseParts, fmt.Sprintf(introFormat, ordinals[n]))
	//count down to 1
	for i := n; i >= 1; i-- {
		phrase := fmt.Sprintf("%s %s", cardinals[i], gifts[i])
		if !isOne && i == 1 {
			verseParts = append(verseParts, "and "+phrase)
		} else {
			verseParts = append(verseParts, phrase)
		}
	}

	return strings.Join(verseParts, ", ")

}

// Song gives you back the full lyrics of the song
func Song() string {
	songParts := make([]string, 0, 12)
	for i := 1; i <= 12; i++ {
		songParts = append(songParts, Verse(i))
	}

	return strings.Join(songParts, "\n") + "\n"
}
