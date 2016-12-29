// Package palindrome allows a user to find a palindromic product within a given range of numbers
package palindrome

import "errors"

// testVersion is only 1
const testVersion = 1

// Product represents all of the products to get a palindromic number
type Product struct {
	Product        int      //the palindromic number we've found
	Factorizations [][2]int //the list of all possible factors
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

// reverseNum takes a number and reverses all the digits
func reverseNum(x int) int {
	sgn := 1
	if x < 0 {
		sgn = -1
	}

	var reverse int
	for n := x; n > 0; n /= 10 {
		reverse = reverse*10 + abs(n%10)
	}

	return reverse * sgn
}

// Products computes the factorizations of the minimum and maximum palindromic in a given range
func Products(fmin, fmax int) (pmin, pmax Product, err error) {

	//bail in case of an inverted range
	if fmax < fmin {
		err = errors.New("fmin > fmax")
		return
	}

	prodMap := make(map[int][][2]int)
	for n := fmin; n <= fmax; n++ {
		for m := n; m <= fmax; m++ {
			if n*m == reverseNum(n*m) {
				prodMap[n*m] = append(prodMap[n*m], [2]int{n, m})
			}
		}
	}

	// bail if we have no palindromes
	if len(prodMap) == 0 {
		err = errors.New("No palindromes")
		return
	}

	for k, v := range prodMap {
		if k < pmin.Product || len(pmin.Factorizations) == 0 {
			pmin.Product = k
			pmin.Factorizations = v
		}

		if k > pmax.Product || len(pmax.Factorizations) == 0 {
			pmax.Product = k
			pmax.Factorizations = v
		}

	}

	return
}
