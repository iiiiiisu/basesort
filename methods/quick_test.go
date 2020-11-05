package methods

import "testing"

func TestQuickSort_Sort(t *testing.T) {
	data := getTestData()
	suffle(data)
	s := QuickSort{Data: data, Length: 10}
	s.Sort()
	if !s.CheckResult() {
		t.Error(s.Data)
		t.Error("Quick Sort Error")
	}
}