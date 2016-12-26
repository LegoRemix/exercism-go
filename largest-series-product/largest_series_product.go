// Package lsproduct computes the largest product possible in a string of numbers over a span of n
package lsproduct

import (
	"errors"
	"strconv"
	"unicode"
)

// the test version is simply 1
const testVersion = 4

// LargestSeriesProduct computes the maximum possible product in a string over a given span
func LargestSeriesProduct(numbers string, spanSize int) (int, error) {
	// sanity check to make sure that we have a valid spanSize
	if spanSize < 0 {
		return 0, errors.New("Invalid spanSize passed as argument")
	}

	// first we need to do a pass where we convert the string into a list of numbers
	nums := make([]int, 0, len(numbers))

	for _, r := range numbers {
		if !unicode.IsDigit(r) {
			return 0, errors.New("Found a non-numeric character in the input string")
		}

		//convert the digit to an int
		n, err := strconv.Atoi(string(r))

		// bubble up any errors
		if err != nil {
			return 0, err
		}

		// append the number to the array
		nums = append(nums, n)
	}

	// make sure our spanSize isn't bigger than our array
	if len(nums) < spanSize {
		return 0, errors.New("SpanSize greater than the size of input numeric string")
	}
	var max int

	//slide along the entire array
	for i := 0; i < len(nums)-(spanSize-1); i++ {
		var curr int = 1
		//compute the product along this span
		for j := 0; j < spanSize; j++ {
			curr *= nums[j+i]
		}

		// check if we have a bigger product
		if curr > max {
			max = curr
		}
	}

	return max, nil

}
