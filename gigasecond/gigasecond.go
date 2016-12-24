// Package Gigasecond helps users compute when they will have lived a gigasecond
package gigasecond

import "time"

// Test version for gigasecond
const testVersion = 4

//A gigasecond is *exactly* 10^9 seconds
const GIGASEC time.Duration = 1000000000 * time.Second

//Computes when someone has lived an entire gigasecond
func AddGigasecond(birthDate time.Time) time.Time {
	return birthDate.Add(GIGASEC)
}
