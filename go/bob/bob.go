// Package bob can respond to any remarks!
package bob

import (
	"strings"
)

// Hey will respond appropriately
// All credit goes to tjkomplex genious answer. My solution sucked
func Hey(remark string) string {
	r := strings.TrimSpace(remark)

	switch {
	case r == "":
		return "Fine. Be that way!"
	case strings.ToUpper(r) == r && strings.ToLower(r) != r && strings.HasSuffix(r, "?"):
		return "Calm down, I know what I'm doing!"
	case strings.ToUpper(r) == r && strings.ToLower(r) != r:
		return "Whoa, chill out!"
	case strings.HasSuffix(r, "?"):
		return "Sure."
	default:
		return "Whatever."
	}
}
