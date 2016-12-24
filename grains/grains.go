// Package grains counts the number of grains on a chess board
package grains

import "errors"

// numSquares is the number of squares on the board
const numSquares = 64

// testVersion is just one
const testVersion = 1

// pow computes a^b power for ints
func pow(a uint64, b uint64) (uint64, error) {
	//handle the simple case where b == 0
	if b == 0 {
		return 1, nil
		// error on negative exponent
	} else if b < 0 {
		return 0, errors.New("Negative exponent not legal")
	}

	var y uint64 = uint64(1)
	for b > 1 {
		if b%2 == 0 {
			a *= a
			b /= 2
		} else {
			y = a * y
			a *= a
			b = (b - 1) / 2
		}
	}
	return a * y, nil
}

// Square computes the number of grains of rice on a given square
func Square(n int) (uint64, error) {
	if n < 1 || n > numSquares {
		return 0, errors.New("Attempted to access square out of bounds")
	}

	res, _ := pow(2, uint64(n-1))

	return res, nil
}

// Total computes the total number of grains of rice on the board
func Total() uint64 {
	var total uint64
	for i := 1; i <= numSquares; i++ {
		res, _ := Square(i)
		total += res
	}
	return total
}
