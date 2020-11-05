package methods

import "testing"

func TestBubbleSort_Sort(t *testing.T) {
	data := getTestData()
	suffle(data)
	s := BubbleSort{Data: data, Length: 10}
	s.Sort()
	if !s.CheckResult() {
		t.Error(s.Data)
		t.Error("Bubble Sort Error")
	}
}