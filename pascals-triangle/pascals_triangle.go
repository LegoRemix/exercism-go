// Package pascal computes pascals triangle
package pascal

// the index into our memoize dictionary
type index struct {
	row int
	col int
}

// our memoization for the pascal function
var pascalMemo = map[index]int{
	index{}: 1,
}

// pascal computes the pascal value for n, k
func pascal(n, k int) int {
	if n == 0 && k == 0 {
		return 1
	} else if n == 0 || k == 0 {
		return 0
	}

	idx := index{row: n, col: k}

	//try to look up our memo
	if v, ok := pascalMemo[idx]; ok {
		return v
	}

	//calculate it if we don't have it in the memo
	ans := pascal(n-1, k-1) + pascal(n-1, k)

	//save the answer
	pascalMemo[idx] = ans

	return ans

}

// Layer computes the nth row of pascal's triangle
func Layer(row int) []int {

	var layer []int
	for i := 0; i < row; i++ {
		layer = append(layer, pascal(row, i+1))
	}

	return layer
}

// Triangle computes the first n rows of pascal's triangle
func Triangle(nrows int) [][]int {
	var triangle [][]int
	for i := 1; i <= nrows; i++ {
		triangle = append(triangle, Layer(i))
	}
	return triangle
}
