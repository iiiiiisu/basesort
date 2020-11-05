package methods

type ShellSort struct {
	Data []uint32
	Length int32
}

func (s *ShellSort) SetData(data []uint32)  {
	s.Data = data
	s.SetLength()
}

func (s *ShellSort) SetLength()  {
	s.Length = int32(len(s.Data))
}

func (s *ShellSort) Sort() {
	var i, j, m, n int32
	for i = s.Length / 2; i > 0; i /= 2 {
		for j = 0; j < i; j ++ {
			for m = j + i; m < s.Length; m += i {
				for n = m; n > j; n -= i {
					if s.Data[n] < s.Data[n-i] {
						swap(s.Data, n, n-i)
					} else {
						continue
					}
				}
			}
		}
	}
}

func (s *ShellSort) CheckResult() bool {
	var i int32
	for i=1; i < s.Length; i++ {
		if s.Data[i] < s.Data[i-1] {
			return false
		}
	}
	return true
}
