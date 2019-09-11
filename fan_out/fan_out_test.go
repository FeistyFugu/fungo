package fan_out

//go:generate fungo -template=FanOut -fileName=fan_out

import (
	"strings"
	"testing"
)

func CompareT(a, b []string) bool {
	for i := range a {
		if string(a[i]) != string(b[i]) {
			return false
		}
	}
	return true
}

func TestFanOut(t *testing.T) {
	// Perform tests 100 times to make sure no timing issues occur

	// Concurrent mapping test
	for i := 0; i < 100; i ++ {
		source := [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		var destination [10]string
		dispatch := func(index int, item string) (int, string, error) {
			return index, string(strings.ToUpper(string(item))), nil
		}
		merge := func(index int, item string) (bool, error) {
			destination[index] = item
			return false, nil
		}
		if err := FanOut(source[:], dispatch, merge); err != nil {
			t.Error("FanOut -> No error was expected")
		}
		expected := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
		if !CompareT(expected, destination[:]) {
			t.Error("FanOut -> Expected", source, "but got", destination)
		}
	}

	// Merge cancellation test
	for i := 0; i < 100; i ++ {
		source := [10]string{"aa", "bbbb", "c", "ddd", "eeeee", "f", "gg", "hhh", "iiiii", "jj"}
		dispatch := func(index int, item string) (int, string, error) {
			return index, string(strings.ToUpper(string(item))), nil
		}
		pos := 0
		merge := func(index int, item string) (bool, error) {
			pos = index
			return len(item) == 4, nil
		}
		if err := FanOut(source[:], dispatch, merge); err != nil {
			t.Error("FanOut -> No error was expected")
		}
		if pos != 1 {
			t.Error("FauOut -> Expected 1 but got", pos)
		}
	}


}
