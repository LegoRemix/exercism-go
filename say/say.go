// Package say converts a number into an english language phrase
package say

import "strings"

// numsToName maps integers to their names
var numsToName = map[uint64]string{
	0:                   "zero",
	1:                   "one",
	2:                   "two",
	3:                   "three",
	4:                   "four",
	5:                   "five",
	6:                   "six",
	7:                   "seven",
	8:                   "eight",
	9:                   "nine",
	10:                  "ten",
	11:                  "eleven",
	12:                  "twelve",
	13:                  "thirteen",
	14:                  "fourteen",
	15:                  "fifteen",
	16:                  "sixteen",
	17:                  "seventeen",
	18:                  "eighteen",
	19:                  "nineteen",
	20:                  "twenty",
	30:                  "thirty",
	40:                  "forty",
	50:                  "fifty",
	60:                  "sixty",
	70:                  "seventy",
	80:                  "eighty",
	90:                  "ninety",
	100:                 "hundred",
	1000:                "thousand",
	1000000:             "million",
	1000000000:          "billion",
	1000000000000:       "trillion",
	1000000000000000:    "quadrillion",
	1000000000000000000: "quintillion",
}

const (
	// cut off to switch to regular double-digit naming
	regularDDNaming uint64 = 20
	// cut off for segmenting into chunks
	chunkSize uint64 = 1000
	//place values
	hundred     uint64 = 100
	thousand    uint64 = 1000
	million     uint64 = 1000 * 1000
	billion     uint64 = 1000 * 1000 * 1000
	trillion    uint64 = 1000 * 1000 * 1000 * 1000
	quadrillion uint64 = 1000 * 1000 * 1000 * 1000 * 1000
	quintillion uint64 = 1000 * 1000 * 1000 * 1000 * 1000 * 1000
)

var placeValues = []struct {
	value uint64
	name  string
}{
	{quintillion, numsToName[quintillion]},
	{quadrillion, numsToName[quadrillion]},
	{trillion, numsToName[trillion]},
	{billion, numsToName[billion]},
	{million, numsToName[million]},
	{thousand, numsToName[thousand]},
	{1, ""},
}

// doubleDigitName returns the name of a number < 100, false if nothing is needed to represent it
func doubleDigitName(n uint64) (string, bool) {
	if n == 0 {
		return "", false
	} else if n > regularDDNaming {
		low := n % 10
		high := n - low
		return numsToName[high] + "-" + numsToName[low], true
	}

	return numsToName[n], true
}

// tripleDigitName returns the name of a number < 1000, false if we don't need any thing to represent it
func tripleDigitName(n uint64) (string, bool) {
	low := n % hundred
	high := n - low
	ddName, ok := doubleDigitName(low)
	triName := ""
	if digit := high / hundred; high != 0 {
		triName = numsToName[digit] + " " + numsToName[hundred]
	}

	switch {
	case !ok && triName == "":
		return "", false
	case !ok && triName != "":
		return triName, true
	case ok && triName == "":
		return ddName, true
	}

	return triName + " " + ddName, true

}

// Say returns the english language rendering of a number
func Say(num uint64) string {
	if num == 0 {
		return numsToName[num]
	}

	var strs []string
	for _, pv := range placeValues {
		low := num % pv.value
		high := num / pv.value
		num = low
		name, ok := tripleDigitName(high)
		if ok {
			if pv.name != "" {
				name = name + " " + pv.name
			}
			strs = append(strs, name)
		}
	}

	return strings.Join(strs, " ")
}
