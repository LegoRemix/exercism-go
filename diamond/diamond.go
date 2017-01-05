// Package diamond implements a fancy diamond printing routine
package diamond

import (
	"bytes"
	"errors"
)

// we're on the first iteration of the test suite
const testVersion = 1

// ErrOutOfRange signifies that the target byte is not a letter
var ErrOutOfRange = errors.New("diamond : argument provided is out of range")

// genLine generates a specific line of the diamond
func genLine(curr, limit byte) bytes.Buffer {
	spaces := int(limit - curr)
	isA := curr == 'A'
	middleSpaces := 1 + 2*int(curr-'B')
	var b bytes.Buffer

	// print our spaces and first character
	for i := 0; i < spaces; i++ {
		b.WriteByte(' ')
	}
	b.WriteByte(curr)

	//print middle spaces
	if !isA {
		for i := 0; i < middleSpaces; i++ {
			b.WriteByte(' ')
		}
		b.WriteByte(curr)
	}

	//print our final spaces
	for i := 0; i < spaces; i++ {
		b.WriteByte(' ')
	}

	b.WriteByte('\n')
	return b
}

// Gen generates a diamond of a from letter [A, limit]
func Gen(limit byte) (string, error) {
	// bounds check our function
	if limit < 'A' || limit > 'Z' {
		return "", ErrOutOfRange
	}

	savedLines := make(map[byte]bytes.Buffer)

	var b bytes.Buffer
	for c := byte('A'); c <= limit; c++ {
		line := genLine(c, limit)
		b.Write(line.Bytes())
		savedLines[c] = line
	}

	for c := limit - 1; c >= 'A'; c-- {
		line := savedLines[c]
		b.Write(line.Bytes())
	}

	return b.String(), nil

}
