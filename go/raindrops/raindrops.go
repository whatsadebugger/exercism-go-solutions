// Package raindrops makes it rain
package raindrops

import (
	"fmt"
)

// Convert takes a value and will proceed to make it rain
func Convert(i int) string {
	s := ""
	if i%3 == 0 {
		s += "Pling"
	}
	if i%5 == 0 {
		s += "Plang"
	}
	if i%7 == 0 {
		s += "Plong"
	}
	if len(s) == 0 {
		s = fmt.Sprintf("%v", i)
	}
	return s
}
