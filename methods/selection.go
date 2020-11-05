package methods


type SelectionSort struct {
	Data []uint32
	Length int32
}

func (s *SelectionSort) SetData(data []uint32)  {
	s.Data = data
	s.SetLength()
}

func (s *SelectionSort) SetLength()  {
	s.Length = int32(len(s.Data))
}

func (s *SelectionSort) Sort() {
	var i, j int32
	var min int32
	for i = 0; i < s.Length - 1; i++ {
		min = i
		for j = i; j < s.Length; j++ {
			if s.Data[j] < s.Data[min] {
				min = j
			}
		}
		swap(s.Data, i, min)
	}
}

func (s *SelectionSort) CheckResult() bool {
	var i int32
	for i=1; i < s.Length; i++ {
		if s.Data[i] < s.Data[i-1] {
			return false
		}
	}
	return true
}

