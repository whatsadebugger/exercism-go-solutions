// Package raindrops makes it rain
package raindrops

import (
	"fmt"
)

// Convert takes a value and will proceed to make it rain
func Convert(i int) string {
	s := ""
	hasFactor := false
	if i%3 == 0 {
		s += "Pling"
		hasFactor = true
	}
	if i%5 == 0 {
		s += "Plang"
		hasFactor = true
	}
	if i%7 == 0 {
		s += "Plong"
		hasFactor = true
	}
	if !hasFactor {
		s = fmt.Sprintf("%v", i)
	}
	return s
}
