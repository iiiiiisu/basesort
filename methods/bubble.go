package methods

type BubbleSort struct {
	Data []uint32
	Length int32
}

func (s *BubbleSort) SetData(data []uint32)  {
	s.Data = data
	s.SetLength()
}

func (s *BubbleSort) SetLength()  {
	s.Length = int32(len(s.Data))
}

func (s *BubbleSort) Sort() {
	var i, j int32
	for i = 0; i < s.Length - 1; i++ {
		for j = 0; j < s.Length - i - 1; j++ {
			if s.Data[j] > s.Data[j+1] {
				swap(s.Data, j, j+1)
			}
		}
	}
}

func (s *BubbleSort) CheckResult() bool {
	var i int32
	for i=1; i < s.Length; i++ {
		if s.Data[i] < s.Data[i-1] {
			return false
		}
	}
	return true
}
