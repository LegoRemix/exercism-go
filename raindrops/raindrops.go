// Package raindrops handles converting to raindrop language from numbers
package raindrops

import "strconv"

const testVersion = 2

// Convert handles converting a number to raindrop language
func Convert(input int) string {
	out := ""

	//check if this number is divisible by 3 (Pling)
	if input%3 == 0 {
		out = "Pling"
	}

	//check if the number is divisible by 5 (Plang)
	if input%5 == 0 {
		out = out + "Plang"
	}

	//check if the number is divisible by 7 (Plong)
	if input%7 == 0 {
		out = out + "Plong"
	}

	//if we go into this branch, we have no factors in {3,5,7}
	if len(out) == 0 {
		return strconv.Itoa(input)
	}

	return out
}
