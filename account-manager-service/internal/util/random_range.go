package util

import "math/rand"

func GetRandomNumberBetween(min, max int) int {
	return rand.Intn(max-min) + min
}
