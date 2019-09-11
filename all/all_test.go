package all

//go:generate fungo -template=All -fileName=all

import (
	"errors"
	"testing"
)

var list1 = []string{"aaa", "bbb", "ccc", "ddd", "eee"}
var list2 = []string{"aaa", "bb", "ccc", "ddd", "eee"}
var expLen = 3

func pred1(x string) (bool, error) {
	return len(x) == expLen, nil
}

func pred2(x string) (bool, error) {
	if x[0] == 'd' {
		return true, errors.New("this is wrong") // true should be overridden
	}
	return len(x) == expLen, nil
}

func pred3(x string) (bool, error) {
	if x[0] == 'd' {
		return false, errors.New("this is wrong")
	}
	return len(x) == expLen, nil
}

func TestAllPositive(t *testing.T) {
	result, _ := All(list1, pred1)
	if !result {
		t.Error("All -> Should have found that all items in", list1, "have a length of", expLen)
	}
}

func TestAllNegative(t *testing.T) {
	result, _ := All(list2, pred1)
	if result {
		t.Error("All -> Should have found that at least one item in", list2, "have a length different than", expLen)
	}
}

func TestAllWithError(t *testing.T) {
	result, err := All(list1, pred2)
	if result {
		t.Error("All -> When an error is returned, the result of All should always be false")
	}
	if err == nil {
		t.Error("All -> Expected an error to be returned")
	}

	result, err = All(list1, pred3)
	if result {
		t.Error("All -> When an error is returned, the result of All should always be false")
	}
	if err == nil {
		t.Error("All -> Expected an error to be returned")
	}
}
