// Package diffsquares finds the difference of squares
package diffsquares

// SumOfSquares finds sum of squares
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// SquareOfSums square of the sums
func SquareOfSums(n int) int {
	return ((n * (n + 1)) / 2 * (n * (n + 1)) / 2)
}

// Difference finds the difference of sqsums ans sumofsqs
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
