package methods

import "testing"

func TestMergeSort_Sort(t *testing.T) {
	data := getTestData()
	suffle(data)
	s := MergeSort{Data: data, Length: 10}
	s.Sort()
	if !s.CheckResult() {
		t.Error(s.Data)
		t.Error("Merge Sort Error")
	}
}