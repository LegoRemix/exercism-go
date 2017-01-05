// Package robotname has a list of serial numbers for robots
package robotname

import (
	"fmt"
	"math/rand"
)

// Robot is a uniquely named robot instance
type Robot string

// usedNames is a set of all of our previously used names
var usedNames = make(map[string]bool)

// letters is all of the upper case latin letters
var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// randomLetter gets a random letter from the uppercase latin alphabet
func randomLetter() string {
	return string(letters[rand.Intn(len(letters))])
}

// genName gets a random name for the robot
func genName() string {
	for {
		name := randomLetter() + randomLetter() + fmt.Sprintf("%3d", rand.Intn(1000))
		if !usedNames[name] {
			usedNames[name] = true
			return name
		}
	}
}

// Name returns the unique name for the Robot, even if it has to be generated
func (r *Robot) Name() string {
	if *r == "" {
		*r = Robot(genName())
	}
	return string(*r)
}

// Reset clears the name of the robot so that it will get a new one on next request
func (r *Robot) Reset() {
	*r = Robot("")
}
