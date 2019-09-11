package apply

//go:generate fungo -template=Apply -fileName=apply

import (
	"errors"
	fc "fungo/contains"
	"strings"
	"testing"
)

var list = []string{"aaa", "bbb", "ccc", "ddd", "eee"}
var expectedList = []string{"AAA", "BBB", "CCC", "DDD", "EEE"}

func pred1(x string) (string, error) {
	return string(strings.ToUpper(string(x))), nil
}

func equal(a string, b string) (bool, error) {
	return a == b, nil
}

func pred2(x string) (string, error) {
	if x[0] == 'd' {
		return "", errors.New("this is wrong") // true should be overridden
	}
	return string(strings.ToUpper(string(x))), nil
}

func TestApply(t *testing.T) {
	result, _ := Apply(list, pred1)
	contained, _, _ := fc.Contains(result, expectedList, equal)
	if !contained || len(result) != 5 {
		t.Error("Apply -> Failed! Expected", expectedList, "got", result)
	}

	_, err := Apply(list, pred2)
	if err == nil {
		t.Error("Apply -> Failed! Expected an error to be returned")
	}
}
