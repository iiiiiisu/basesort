package utils

import (
	"math/rand"
)

func Shuffle(data []uint32, seed ...int64) {
	var (
		i, length int32
	)
	length = int32(len(data))
	for i = 0; i < length; i ++ {
		if len(seed) > 0 {
			rand.Seed(seed[0])
		}
		r := rand.Int31n(length)
		data[i], data[r] = data[r], data[i]
	}
}
