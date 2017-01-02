// Package romannumerals converts an integer into roman numerals
package romannumerals

import (
	"bytes"
	"errors"
)

const testVersion = 3

// Error we through when the integer is out of our conversion range [0, 4000]
var ErrOutOfBounds = errors.New("Submitted value is out of bounds")

// intToNumeral maps numbers to roman numbers
var intToNumeral = []struct {
	value   int
	numeral string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeral converts an integer to a roman numeral
func ToRomanNumeral(number int) (string, error) {
	if number <= 0 || number >= 4000 {
		return "", ErrOutOfBounds
	}

	var buf bytes.Buffer
	for _, numericInfo := range intToNumeral {
		limit := number / numericInfo.value
		for i := 0; i < limit; i++ {
			buf.WriteString(numericInfo.numeral)
			number -= numericInfo.value
		}
	}

	return buf.String(), nil
}
