package methods
import "testing"

func TestShellSort_Sort(t *testing.T) {
	data := getTestData()
	suffle(data)
	s := ShellSort{Data: data, Length: 10}
	s.Sort()
	if !s.CheckResult() {
		t.Error(s.Data)
		t.Error("Shell Sort Error")
	}
}