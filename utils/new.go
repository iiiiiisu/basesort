package utils

func NewArray(length int) []uint32 {
	data := make([]uint32, length)
	for i := 0; i < length; i ++ {
		data[i] = uint32(i + 1)
	}
	return data
}
