// Package secret implements our secret handshake protocol
package secret

// Our test version is 1
const testVersion = 1

// Our number of actions
const numActions = 4

// a mapping from number to actions
var actions = map[uint]string{
	1 << 0: "wink",
	1 << 1: "double blink",
	1 << 2: "close your eyes",
	1 << 3: "jump",
}

// reverse takes a slice and reverses the order of the contents
func reverse(s []string) {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
}

// Handshake computes the handshake protocol
func Handshake(n uint) []string {
	var response []string = nil
	for i := uint(0); i <= numActions; i++ {
		mask := uint(1) << i
		if mask&n != 0 && i != numActions {
			response = append(response, actions[mask])
		} else if mask&n != 0 && i == numActions {
			reverse(response)
		}
	}

	return response
}
