package methods

type InsertionSort struct {
	Data []uint32
	Length int32
}

func (s *InsertionSort) SetData(data []uint32)  {
	s.Data = data
	s.SetLength()
}

func (s *InsertionSort) SetLength()  {
	s.Length = int32(len(s.Data))
}

func (s *InsertionSort) Sort() {
	var i, j int32
	for i = 1; i < s.Length; i++ {
		for j = i-1; j >= 0; j-- {
			if s.Data[j] > s.Data[j+1] {
				swap(s.Data, j, j+1)
			} else {
				continue
			}
		}
	}
}

func (s *InsertionSort) CheckResult() bool {
	var i int32
	for i=1; i < s.Length; i++ {
		if s.Data[i] < s.Data[i-1] {
			return false
		}
	}
	return true
}
