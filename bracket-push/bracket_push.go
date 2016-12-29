// Package brackets tests whether or not we have a valid closed set of braces
package brackets

// closeToOpen maps terminating characters with opening braces
var closeToOpen = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
}

const testVersion = 4

// Bracket tests whether a string contains correctly enclosed brackets
func Bracket(test string) (bool, error) {
	var stack []rune = make([]rune, 0, len(test))
	for _, r := range test {

		//if we find a  closing character and we have a non-empty stack
		if open, ok := closeToOpen[r]; ok && len(stack) > 0 {
			// ... check to make sure it closes out what's on top of the stack
			if stack[len(stack)-1] != open {
				//if it doesn't, this string fails
				return false, nil
			}
			//if it does, pop the item from the stack
			stack = stack[:len(stack)-1]

			// if we have an empty stack and a closing character, bail
		} else if len(stack) == 0 && ok {
			return false, nil

		} else {
			//otherwise, we push onto the stack
			stack = append(stack, r)
		}
	}

	return len(stack) == 0, nil
}
