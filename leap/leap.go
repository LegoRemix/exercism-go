// Package for computing whether a year was a leap year

// The package name is expected by the test program.
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// Computes if a year is a leap year (every 4th year with exceptions)
func IsLeapYear(year int) bool {
	//Every 4th year is a leap year UNLESS it's a year ending in 00
	//not divisible by 400
	return year%4 == 0 && !(year%100 == 0 && year%400 != 0)
}
