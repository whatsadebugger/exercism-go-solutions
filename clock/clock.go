// Package clock
package clock

import (
	"fmt"
)

// Clock API:
//
// type Clock                      // define the clock type
// New(hour, minute int) Clock     // a "constructor"
// (Clock) String() string         // a "stringer"
// (Clock) Add(minutes int) Clock

// import "time"

// Clock type holds hours and minutes to keep track of time
type Clock struct {
	hour   int
	minute int
}

// New returns a new Clock initialized with given values
func New(h, m int) Clock {
	if m == -40 {
		fmt.Print("wow")
	}
	//shift extra hours to h
	h += m / 60
	//prune minutes
	m %= 60
	//prune days
	h %= 24
	//if negative remove one hour
	if m < 0 {
		m += 60
		h--
	}
	if h < 0 {
		h += 24
	}

	return Clock{hour: h, minute: m}
}

// Add goes to the future
func (c Clock) Add(minutes int) Clock {
	totalMin := c.minute + minutes
	// add mins
	futureMin := totalMin % 60
	futureMin += c.minute

	//add hours
	futureHour := totalMin / 60
	futureHour += c.hour
	// prune days
	futureHour %= 24

	return Clock{hour: futureHour, minute: futureMin}
}

// Subtract goes back in time
func (c Clock) Subtract(minutes int) Clock {
	// get total hours to remove
	hours := minutes / 60
	pastHour := c.hour - hours
	hours %= 24

	if pastHour < 0 {
		pastHour += 24
	}

	// get total minutes to sub
	minutes %= 60
	pastMin := c.minute - minutes

	if pastMin < 0 {
		pastMin += 60
	}

	return Clock{hour: pastHour, minute: pastMin}
}

// String will return a string representation of the clock
func (c Clock) String() string {
	h, m := fmt.Sprintf("%v", c.hour), fmt.Sprintf("%v", c.minute)
	if c.hour < 10 {
		h = fmt.Sprintf("0%v", c.hour)
	}
	if c.minute < 10 {
		m = fmt.Sprintf("0%v", c.minute)
	}
	return fmt.Sprintf("%v:%v", h, m)
}
