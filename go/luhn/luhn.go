// Package luhn validates per luhn formula
package luhn

import (
	"strconv"
	"strings"
)

// Valid validates if a number is valid per Luhn formula
func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)

	if _, err := strconv.ParseInt(s, 10, 0); err != nil || len(s) < 2 {
		return false
	}

	return luhn(s)
}

func luhn(s string) bool {
	nums := strings.Split(s, "")
	sum := 0
	double := false
	for i := len(s) - 1; i >= 0; i-- {
		v, _ := strconv.ParseInt(nums[i], 10, 8)
		n := int(v)

		if double {
			if n*2 >= 10 {
				n = n*2 - 9
			} else {
				n = n * 2
			}
		}

		double = !double
		sum += n
	}

	return sum%10 == 0
}
