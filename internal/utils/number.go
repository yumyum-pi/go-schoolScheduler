package utils

import (
	"math/rand"
	"time"
)

// GenerateRandomInt return a random int with
// given int and a factor
func GenerateRandomInt(i int, factor int) int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(i * factor)
	n /= factor

	return n
}

// RangeInt is struct that has a two varible-min, max.
// Using it's Random() menthod to get a random number between
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
