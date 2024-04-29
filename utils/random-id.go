package utils

import (
	"math/rand"
)

func RandomID() int {
	return rand.Intn(9000000000000000000)
}
