// Package summultiples lets a user sum all of the multiples of a given number
package summultiples

//Summer is a function which finds the sum of numbers
type Summer func(int) int

// MultipleSummer returns a closure which when provided an int, sums all numbers [0,N) divisible by the provided factors
func MultipleSummer(args ...int) Summer {
	return func(n int) int {
		var sum int
		for i := 0; i < n; i++ {
			for _, num := range args {
				if i%num == 0 {
					sum += i
					break
				}
			}
		}
		return sum
	}
}
