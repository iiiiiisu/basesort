package methods

type QuickSort struct {
	Data []uint32
	Length int32
}

func (s *QuickSort) SetData(data []uint32)  {
	s.Data = data
	s.SetLength()
}

func (s *QuickSort) SetLength()  {
	s.Length = int32(len(s.Data))
}

func (s *QuickSort) getIndex(low, high int32) int32 {
	temp := s.Data[low]
	for low < high {
		for low < high && s.Data[high] >= temp {
			high --
		}
		s.Data[low] = s.Data[high]
		for low < high && s.Data[low] <= temp {
			low ++
		}
		s.Data[high] = s.Data[low]
	}
	s.Data[low] = temp
	return low
}

func (s *QuickSort) Sort() {
	type Stack struct {
		Low int32
		High int32
	}
	var low, high int32 = 0, s.Length - 1
	var index int32 = 0
	var stack = make([]Stack, s.Length/2)
	var length = 0
	for true {
		if length > 0 {
			low, high = stack[length-1].Low, stack[length-1].High
			length --
		}
		index = s.getIndex(low, high)

		if index - 1 > low {
			stack[length] = Stack{low, index-1}
			length ++
		}
		if high > index + 1 {
			stack[length] = Stack{index+1, high}
			length ++
		}
		if length == 0 {
			break
		}
	}
}

func (s *QuickSort) CheckResult() bool {
	var i int32
	for i=1; i < s.Length; i++ {
		if s.Data[i] < s.Data[i-1] {
			return false
		}
	}
	return true
}
