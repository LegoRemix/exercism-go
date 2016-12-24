// Package slice
package slice

import "errors"

// All gives you back the all contiguous substrings of length n, in string s
func All(n int, input string) []string {
	var substrs []string

	//We iterate through the parts of the string
	for start, end := 0, n; end <= len(input); start, end = start+1, end+1 {
		substrs = append(substrs, input[start:end])
	}

	return substrs
}

// UnsafeFirst gives you back the results of All, but without error checking
func UnsafeFirst(n int, input string) string {
	if n > len(input) {
		return ""
	}

	return input[0:n]
}

// First provides error handling on top of UnsafeFirst
func First(n int, input string) (string, error) {
	if n > len(input) {
		return "", errors.New("Invalid substr length parameter")
	}
	return UnsafeFirst(n, input), nil
}
