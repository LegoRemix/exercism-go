// Package diffsquares computes the difference of two squares
package diffsquares

// SumOfSquares computes the sum of all squares from [1,N]
func SumOfSquares(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		ans += i * i
	}

	return ans
}

// SquareOfSums computes the square of the sum from [1,N]
func SquareOfSums(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		ans += i
	}

	return ans * ans
}

// Difference computes Sum(i^2, 1, n) - Sum(i, 1, n)^2
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
