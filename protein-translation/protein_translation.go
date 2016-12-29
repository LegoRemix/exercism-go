// Package protein allows a user to convert RNA sequences into proteins
package protein

// currently we're on the first version of the test
const testVersion = 1

// here we define the names of proteins
// so we don't have to worry about spelling more than once
const (
	meth   = "Methionine"
	phenyl = "Phenylalanine"
	leuc   = "Leucine"
	ser    = "Serine"
	tyro   = "Tyrosine"
	cys    = "Cysteine"
	trypt  = "Tryptophan"
	stop   = "STOP"
)

//define our map from sequences to proteins
var codonMap = map[string]string{
	"AUG": meth,
	"UUU": phenyl,
	"UUC": phenyl,
	"UUA": leuc,
	"UUG": leuc,
	"UCU": ser,
	"UCC": ser,
	"UCA": ser,
	"UCG": ser,
	"UAU": tyro,
	"UAC": tyro,
	"UGU": cys,
	"UGC": cys,
	"UGG": trypt,
	"UAA": stop,
	"UAG": stop,
	"UGA": stop,
}

// FromCodon takes a single codon and returns the corresponding protein
func FromCodon(codon string) string {
	return codonMap[codon]
}

// FromRNA takes a series of codons and produces the list of resultant proteins
func FromRNA(rna string) []string {
	var result []string
	for i := 0; i < len(rna)-2; i += 3 {
		if protein := FromCodon(rna[i : i+3]); protein != stop {
			result = append(result, protein)
		} else {
			return result
		}
	}
	return result
}
