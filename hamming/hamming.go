//Package hamming computes the hamming between two strings
package hamming

import "errors"

const testVersion = 5

//computes the hamming distance of two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("hamming : non-equal string lengths")

	}

	var diffCount int = 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffCount++
		}
	}

	return diffCount, nil
}
