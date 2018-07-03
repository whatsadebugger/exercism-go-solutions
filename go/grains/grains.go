// Package grains counts all the grains on chess board
package grains

import "errors"

// Square returns number of grains on square
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("Incorrect Input value")
	}
	return (1 << uint(n-1)), nil
}

// Total returns total grains on chess board
func Total() uint64 {
	return 1<<64 - 1
}
