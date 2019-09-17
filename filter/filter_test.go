package filter

//go:generate fungo -template=Filter -fileName=filter -t1=int
//go:generate fungo -template=Match -fileName=int_equal -functionName=intEqual -t1=int -t2=int

import (
	"errors"
	"testing"
)

func check(a, b []int) bool {
	ok, _ := intEqual(a, b, func(a, b int) (bool, error) {
		return a == b, nil
	})
	return ok
}

func TestFilter(t *testing.T) {
	compare := func(x int) (bool, error) {
		return x % 2 == 0, nil
	}

	data := []int{1, 2, 3, 4, 5, 6}
	result, _ := Filter(data[:], compare)
	expected := []int{2, 4, 6}
	if !check(result, expected) {
		t.Error("Filter -> Expected", expected, "got", result)
	}

	data = []int{1, 3, 5, 7, 9, 11, 13}
	result, _ = Filter(data[:], compare)
	expected = []int{}
	if !check(result, expected) {
		t.Error("Filter -> Expected", expected, "got", result)
	}
}

func TestFilterEmptySlice(t *testing.T) {
	var data []int
	result, _ := Filter(data[:], func(x int) (bool, error) {
		return x % 2 == 0, nil
	})
	var expected []int
	if !check(result, expected) {
		t.Error("Filter -> Expected", expected, "got", result, "'")
	}
}

func TestFilterWithError(t *testing.T) {
	compare := func(x int) (bool, error) {
		if x > 4 {
			return true, errors.New("Too high!")
		}
		return true, nil
	}

	data := []int{1, 2, 3, 4, 5, 6}
	_, err := Filter(data[:], compare)
	if err == nil {
		t.Error("Filter -> Expected an error to be returned")
	}
}
