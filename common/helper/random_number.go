package helpers

import (
	"time"

	"math/rand"
)

func RandomNumber(length int) int {
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(9*(10^length)) + 10 ^ length
	return randomNumber
}
