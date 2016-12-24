// Package pythagorean allows a user to generate pythagorean triplets between two ranges
package pythagorean

import "sort"

// Triplet represents a pythagorean triplet
type Triplet [3]int

// TripSlice is a list of triplets
type TripSlice []Triplet

// Len computes the length of a triplet slice
func (a TripSlice) Len() int { return len(a) }

//Swap trades the position of two elements in the slice
func (a TripSlice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

//Less checks if one Triplet is less than another
func (a TripSlice) Less(i, j int) bool {
	if a[i][0] != a[j][0] {
		return a[i][0] < a[j][0]
	} else if a[i][1] != a[j][1] {
		return a[i][1] < a[j][1]
	}
	return a[i][2] < a[j][2]
}

// min computes the min of two values
func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// maximum computes the max of two values
func maximum(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Range computes all pythagorean triples with sides in a given range [min, max)
func Range(min, max int) []Triplet {

	var trips []Triplet
	// m can be anywhere from 2 (minimum possible value where m > n > 0) and max - 1 (since the minimum value of n is 1)
	for m := 2; m < (max - 1); m++ {

		start := 1
		if m%2 == 1 {
			start = 2
		}

		// n ranges from 1 to m, and since one must be even, we skip by 2
		for n := start; n < m && (m*m+n*n) <= max; n += 2 {
			a := m*m - n*n
			b := 2 * m * n
			c := m*m + n*n
			lowest := minimum(a, b)
			second := maximum(a, b)
			minK := min / lowest
			if min%lowest != 0 {
				minK++
			}

			for k := maximum(1, minK); k*c <= max; k++ {
				trips = append(trips, Triplet{lowest * k, second * k, c * k})
			}
		}
	}

	return trips
}

// Sum returns a list of triples such that the sum of the sides is p
func Sum(perimeter int) []Triplet {

	var result []Triplet

	//just in case we generate duplicates we keep around a list of what we already added
	generated := make(map[Triplet]bool)
	// compute all possible candidates
	candidates := Range(1, perimeter)
	for _, c := range candidates {
		if c[0]+c[1]+c[2] == perimeter && !generated[c] {
			result = append(result, c)
			generated[c] = true
		}
	}

	//sort since we may generate these trips out of lexographically order
	sort.Sort(TripSlice(result))
	return result
}
