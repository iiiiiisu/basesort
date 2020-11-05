package methods

type MergeSort struct {
	Data []uint32
	Length int32
}

func (s *MergeSort) SetData(data []uint32)  {
	s.Data = data
	s.SetLength()
}

func (s *MergeSort) SetLength()  {
	s.Length = int32(len(s.Data))
}

func (s *MergeSort) getIndex(low, high int32) int32 {
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

func (s *MergeSort) PartSort(pos, length int32) {
	var (
		posA, posB, posBEnd int32
	)
	if pos + length >= s.Length {
		return
	}
	temp := make([]uint32, length * 2)
	var tempLen int32 = 0
	posA=pos
	posB=pos+length
	if s.Length < pos + 2 * length{
		posBEnd = s.Length
	} else {
		posBEnd = pos + 2 * length
	}
	for {
		if s.Data[posA] > s.Data[posB] {
			temp[tempLen] = s.Data[posB]
			posB ++
		} else {
			temp[tempLen] = s.Data[posA]
			posA ++
		}
		tempLen ++
		if posA >= pos + length || posB >= posBEnd {
			break
		}
	}
	if posA >= pos + length {
		for posB < posBEnd {
			temp[tempLen] = s.Data[posB]
			posB ++
			tempLen ++
		}
	} else if posB >= posBEnd {
		for posA < pos + length {
			temp[tempLen] = s.Data[posA]
			posA ++
			tempLen ++
		}
	}
	for i := int32(0); i < tempLen; i++ {
		s.Data[pos+i] = temp[i]
	}
}

func (s *MergeSort) Sort() {
	var i, j int32
	for i = 1; i < s.Length; i *= 2 {
		for j = 0; j < s.Length; j += 2*i {
			s.PartSort(j, i)
		}
	}
}

func (s *MergeSort) CheckResult() bool {
	var i int32
	for i=1; i < s.Length; i++ {
		if s.Data[i] < s.Data[i-1] {
			return false
		}
	}
	return true
}
