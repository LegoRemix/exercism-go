// Package cryptosquare converts a blob of text into a cryptographic square
package cryptosquare

import (
	"bytes"
	"math"
	"unicode"
)

const testVersion = 2

// Encode string to square code
func Encode(plaintxt string) string {

	//filter out all non-alphanumeric characters
	var b1 bytes.Buffer
	for _, r := range plaintxt {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b1.WriteRune(unicode.ToLower(r))
		}
	}

	//if we have no characters, just bail now
	if b1.Len() == 0 {
		return ""
	}

	filtered := b1.String()
	flen := float64(len(filtered))
	ilen := int(flen)
	cols := int(math.Ceil(math.Sqrt(flen)))
	result, rows := bytes.Buffer{}, ilen/cols

	// For each column ...
	for i := 0; i < cols; i++ {
		// For each row ...
		for j := 0; j < rows; j++ {
			result.WriteByte(filtered[i+j*cols])

		}
		//if we're beyond our square, add an additional byte
		if i+rows*cols < ilen {
			result.WriteByte(filtered[i+rows*cols])
		}

		//add the space at the end of each word
		result.WriteByte(' ')

	}

	out := result.String()
	return out[:len(out)-1]
}
