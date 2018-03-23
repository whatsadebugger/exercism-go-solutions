// Package gigasecond has
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds a Gigasecond to the pass in time.Time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
