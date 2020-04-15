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
