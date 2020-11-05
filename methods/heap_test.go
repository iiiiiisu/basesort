package methods

import "testing"

func TestHeapSort_Sort(t *testing.T) {
	data := getTestData()
	suffle(data)
	s := HeapSort{}
	s.SetData(data)
	s.Sort()
	if !s.CheckResult() {
		t.Error(s.Data)
		t.Error("Heap Sort Error")
	}
}