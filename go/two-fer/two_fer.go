// Package twofer should get one for both of us
package twofer

// ShareWith should get one for both of us
func ShareWith(name string) string {
	switch {
	case name == "":
		return "One for you, one for me."
	default:
		return "One for " + name + ", one for me."
	}

}
