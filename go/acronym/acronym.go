// Package acronym abbreviates given string
package acronym

import "strings"

// Abbreviate will take the first letter of each string token
// and return an Abbreiviated UPPERCASE string
func Abbreviate(s string) string {
	tkns := strings.Split(strings.Replace(s, "-", " ", -1), " ")
	var ab strings.Builder
	for _, v := range tkns {
		d := v[0]
		if 'a' <= d && d <= 'z' {
			d = d - 'a' + 'A'
		}
		ab.WriteByte(d)
	}
	return ab.String()
}
