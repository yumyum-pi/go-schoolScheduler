package utils

import (
	"math/rand"
	"time"
)

// RangeInt is struct that has a two variable-min, max.
// Using it's Random() method to get a random number between
// min and max
type RangeInt struct {
	Min int
	Max int
}

// Random return a random value between the range
func (r *RangeInt) Random() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(r.Max-r.Min) + r.Min
}
