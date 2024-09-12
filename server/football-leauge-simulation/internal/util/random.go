package util

import (
	"math/rand"
	"time"
)

// Init initializes the random number generator.
func Init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max.
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// RandomFloat generates a random float64 between min and max.
func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandomChoice returns a random element from the given slice.
func RandomChoice[T any](choices []T) T {
	return choices[rand.Intn(len(choices))]
}
