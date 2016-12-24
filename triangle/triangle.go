// Package triangle computes the type of triangle by side length
package triangle

import (
	"math"
)

const testVersion = 3

type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// isInvalid computes whether or not a number is invalid (Inf or NaN or negative)
func isInvalid(num float64) bool {
	return math.IsNaN(num) || math.IsInf(num, 0) || num <= 0.0
}

// KindFromSides computes the whether or not a
func KindFromSides(a, b, c float64) Kind {
	if isInvalid(a) || isInvalid(b) || isInvalid(c) {
		return NaT
	} else if (a+b) < c || (b+c) < a || (c+a) < b {
		return NaT
	} else if a == b && b == c {
		return Equ
	} else if a == b || b == c || c == a {
		return Iso
	} else {
		return Sca
	}
}
