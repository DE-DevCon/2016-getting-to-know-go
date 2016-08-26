// Adapted from the Exorcism Difference of Squares Golang exercise http://exercism.io/exercises/go/difference-of-squares/readme
// https://github.com/exercism/xgo/blob/master/exercises/difference-of-squares/difference_of_squares_test.go

package diffsquares

// You can also try this here https://play.golang.org/p/VwaPLhXYx0

// SquareOfSums returns (1 + 2 + 3 + ... + n) squared
func SquareOfSums(n int) int {
	// Write Code here
}

// SumOfSquares returns (1 * 1 + 2 * 2 + 3 * 3 + ... + n * n)
func SumOfSquares(n int) (s int) {
	// Write Code here
}

// Difference returns the difference between SquareOfSums(n) and SumOfSquares(n)
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
