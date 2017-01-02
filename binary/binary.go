// Package binary converts strings representing binary numbers into integers
package binary

import "errors"

const testVersion = 2

// ErrInvalidCharacter is thrown if we have a non-binary character in the string
var ErrInvalidCharacter = errors.New("Invalid character found in binary string")

// reverse takes a string and reverses it rune-wise
func reverse(txt string) string {
	runes := []rune(txt)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ParseBinary takes a binary string and tries to convert it to an int
func ParseBinary(binary string) (int, error) {
	var value int
	for p, bit := range reverse(binary) {
		sgn := 0
		if bit != '0' && bit != '1' {
			return 0, ErrInvalidCharacter
		} else if bit == '1' {
			sgn = 1
		}

		value += sgn * (1 << uint(p))
	}
	return value, nil
}
