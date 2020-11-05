package methods

import "testing"

func TestSelectionSort_Sort(t *testing.T) {
	data := getTestData()
	suffle(data)
	s := SelectionSort{Data: data, Length: 10}
	s.Sort()
	if !s.CheckResult() {
		t.Error(s.Data)
		t.Error("Selection Sort Error")
	}
}