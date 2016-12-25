// Package luhn allows a user to compute validity of a set of numbers
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// reverse takes a string and inverts the order of the runes
func reverse(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// preprocess takes a string, trims out whitespace, and reverses it
func preprocess(input string) string {
	return reverse(strings.Trim(input, " "))
}

// luhnSum computes the sum of all digits according to the luhn formula
func luhnSum(input string) (int, error) {

	var digitCount int
	var sum int
	for _, r := range input {
		if unicode.IsDigit(r) {
			//convert this digit to a number
			v, err := strconv.Atoi(string(r))
			//handle any conversion errors
			if err != nil {
				return 0, err
			}

			//if we are on even numbered digit, double it
			if digitCount%2 == 1 {
				v *= 2
			}

			// subtract off 9 if we go 10 or more
			if v > 9 {
				v -= 9
			}
			sum += v
			digitCount++
		}
	}

	return sum, nil
}

// Valid computes whether or not a given set of numbers has a valid checksum
func Valid(input string) bool {
	normalized := preprocess(input)
	if len(normalized) == 0 {
		return false
	}

	v, ok := luhnSum(normalized)
	if ok != nil {
		return false
	}

	return v%10 == 0
}

// AddCheck computes the needed checksum digit for a given set of digits
func AddCheck(test string) string {

	v, _ := luhnSum("0" + preprocess(test))

	res := (10 - v%10) % 10

	return test + strconv.Itoa(res)
}
