package die

import (
	"math/rand"
)

func Roll(size int) int {
	return rand.Intn(size) + 1
}
