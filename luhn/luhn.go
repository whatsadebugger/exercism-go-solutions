// Package luhn validates per luhn formula
package luhn

import (
	"strconv"
	"strings"
)

// Valid validates if a number is valid per Luhn formula
func Valid(s string) bool {
	s = strings.TrimSpace(s)

	if _, err := strconv.ParseInt(s, 10, 0); err != nil || len(s) < 2 {
		return false
	}

	return luhn(s)
}

func luhn(s string) bool {

	strings.Split(s, "")
	return false

}
