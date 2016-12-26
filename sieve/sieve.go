// Package sieve implements the Sieve of Eratosthenes to find primes from [2, N]
package sieve

import "github.com/willf/bitset"

// Sieve implements the Sieve of Eratosthenes to find prime numbers
func Sieve(limit int) []int {
	// we have no prime numbers less than two, so just bail now
	if limit < 2 {
		return nil
	}

	length := uint(limit - 2)
	primeSet := bitset.New(length)
	var primes []int
	//for each open prime number we counter ...
	for idx, found := uint(0), true; found && idx < length; idx, found = primeSet.NextClear(idx + 1) {
		n := idx + 2
		// we found a prime!
		primes = append(primes, int(n))
		//go through and mark off all multiples
		for k := uint(2); k*n-2 < length; k++ {
			primeSet.Set(k*n - 2)
		}

	}

	return primes
}
