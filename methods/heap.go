package methods

type HeapSort struct {
	Data []uint32
	Length int32
}

func (s *HeapSort) SetData(data []uint32)  {
	s.Data = data
	s.SetLength()
}

func (s *HeapSort) SetLength()  {
	s.Length = int32(len(s.Data))
}

func (s *HeapSort) Sort() {
	var i int32 = 0
	s.Build()
	for i = s.Length - 1; i > 0; i-- {
		swap(s.Data, 0, i)
		s.Adjust(0, i - 1)
	}
}

func (s *HeapSort) Build() {
	var i int32
	for i = (s.Length - 1) / 2; i >= 0; i -- {
		s.Adjust(i, s.Length - 1)
	}
}

func (s *HeapSort) Adjust(pos, last int32) {
	var i, node int32
	i = pos
	node = i*2+1
	for node <= last {
		if node + 1 <= last && s.Data[node + 1] > s.Data[node] {
			node ++
		}
		if s.Data[i] < s.Data[node] {
			swap(s.Data, i, node)
		} else {
			break
		}
		i = node
		node = i * 2 + 1
	}
}

func (s *HeapSort) CheckResult() bool {
	var i int32
	for i=1; i < s.Length; i++ {
		if s.Data[i] < s.Data[i-1] {
			return false
		}
	}
	return true
}
