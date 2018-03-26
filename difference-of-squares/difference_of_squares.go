// Package diffsquares finds the difference of squares
package diffsquares

// SumOfSquares finds sum of squares
func SumOfSquares(n int) int {
	sumsq := 0
	for i := 1; i <= n; i++ {
		sumsq += i * i
	}

	return sumsq
}

// SquareOfSums square of the sums
func SquareOfSums(n int) int {
	sqsums := 0
	for i := 1; i <= n; i++ {
		sqsums += i
	}
	return sqsums * sqsums
}

// Difference
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
