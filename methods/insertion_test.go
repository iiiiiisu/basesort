package methods

import "testing"

func TestInsertionSort_Sort(t *testing.T) {
	data := getTestData()
	suffle(data)
	s := InsertionSort{Data: data, Length: 10}
	s.Sort()
	if !s.CheckResult() {
		t.Error(s.Data)
		t.Error("Insertion Sort Error")
	}
}