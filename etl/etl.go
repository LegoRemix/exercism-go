// Package etl takes a given map[int][]string => map[string]int
package etl

import "strings"

// Transform takes a map from ints to a list of strings and inverts it
func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)
	for points, letters := range input {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = points
		}
	}

	return output
}
