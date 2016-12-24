//implements a simple clock
package clock

import "fmt"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

//a simple clock which keeps time in at most minute precision
type Clock struct {
	minutes int
}

//maximum number of minutes per day
const MINUTES_PER_DAY int = 1440

//Arithmetic mod function that always returns a non negative number
func arithMod(dividend int, divisor int) (remainder int) {
	remainder = dividend % divisor
	if remainder < 0 {
		remainder = divisor + remainder
	}
	return remainder
}

//Creates a new instance of a clock using a specified hour and minute
func New(hour, minute int) Clock {
	//Since our precision is in minutes, just convert everything down to that
	minutes := hour*60 + minute
	//compute the arithmetic modulus of the minutes we have
	minutes = arithMod(minutes, MINUTES_PER_DAY)
	return Clock{minutes}

}

//Creates a string representation of the clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

//Adds some amount of minutes to the clock
func (c Clock) Add(minutes int) Clock {
	//we add the minutes to our value stored in the clock
	mins := (c.minutes + minutes)
	//compute the modulus to wrap everything back into bounds
	mins = arithMod(mins, MINUTES_PER_DAY)
	return Clock{mins}

}
