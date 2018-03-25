// Package isogram deals with isograms
package isogram

import "unicode"

// IsIsogram determines if a word is an isogram
// and it should happen pretty fast
func IsIsogram(s string) bool {
	runes := map[rune]bool{}

	for _, v := range s {
		v = unicode.ToLower(v)
		if unicode.IsLetter(v) && runes[v] {
			return false
		}
		runes[v] = true
	}
	return true
}
