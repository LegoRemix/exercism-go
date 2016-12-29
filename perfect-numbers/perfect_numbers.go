// Package perfect checks whether a number is perfect, abundant, or deficient
package perfect

import "errors"

// Classification represents what type of number is this (Deficient, Perfect, Abundant)
type Classification int

const testVersion = 1

const (
	// ClassificationDeficient means that the sum of factors < n
	ClassificationDeficient = iota
	// ClassificationPerfect means that the sum of factors == n
	ClassificationPerfect
	// ClassificationAbundant means that the sum of actors > n
	ClassificationAbundant
)

// ErrOnlyPositive is returned when 0 is passed to the function
var ErrOnlyPositive = errors.New("Cannot classify non-positive numbers")

// Classify takes a number and returns whether it is abundant. deficient, or perfect
func Classify(n uint64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}

	sum := uint64(0)
	for i := uint64(1); i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	if sum < n {
		return ClassificationDeficient, nil
	} else if sum == n {
		return ClassificationPerfect, nil
	}

	return ClassificationAbundant, nil
}
