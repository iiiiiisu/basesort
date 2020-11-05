package methods

import (
	"math/rand"
	"time"
)

type Sorter interface {
	SetData([]uint32)
	SetLength()
	Sort()
	CheckResult() bool
}


func swap(data []uint32, posA, posB int32) {
	data[posA], data[posB] = data[posB], data[posA]
}

func getTestData() []uint32 {
	length := 1000
	data := make([]uint32, length)
	for i:=0; i < length; i++ {
		data[i] = uint32(i + 1)
	}
	return data
}

func suffle(data []uint32) {
	var (
		i, length int32
	)
	length = int32(len(data))
	for i = 0; i < length; i ++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Int31n(length)
		data[i], data[r] = data[r], data[i]
	}
}