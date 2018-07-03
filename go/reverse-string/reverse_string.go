// Package reverse sesrever ffuts
package reverse

import "strings"

// String returns a string that is reversed
func String(s string) string {
	runes := []rune(s)
	bob := strings.Builder{}
	for i := len(runes) - 1; i >= 0; i-- {
		bob.WriteRune(runes[i])
	}
	return bob.String()
}
