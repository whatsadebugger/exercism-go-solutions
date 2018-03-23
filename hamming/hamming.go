// Package hamming will find ham
package hamming

import "errors"

// Distance should return hamming distance
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("a is longer than b")
	}

	distance := 0

	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
