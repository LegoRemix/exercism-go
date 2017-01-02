// Package strand converts DNA nucleotides into RNA
package strand

import "bytes"

const testVersion = 3

// dnaToRNA maps nucleotides from DNA to RNA nucleotides
var dnaToRNA = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA converts a dna strand into rna
func ToRNA(dna string) string {
	var buf bytes.Buffer
	for _, n := range dna {
		buf.WriteRune(dnaToRNA[n])
	}

	return buf.String()
}
