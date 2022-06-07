package bytegen

import (
	"math/rand"
	"time"
)

func GenerateRandomSize() int {
	min := 50   // byte
	max := 8192 // 8kbyte
	return GenerateRandomBetween(min, max)
}
func GenerateRandomBetween(min, max int) int {
	if min > max {
		min, max = max, min
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
