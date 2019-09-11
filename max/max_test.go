package max

//go:generate fungo -template=Max -fileName=max -t1=int

import (
	"errors"
	"testing"
)

func pred1(a int, b int) (int, error) {
	if a == b {
		return 0, nil
	}
	if a < b {
		return -1, nil
	}
	return 1, nil
}

func pred2(a int, b int) (int, error) {
	return 0, errors.New("this is wrong")
}

func TestMin(t *testing.T) {
	var list []int
	_, err := Max(list, pred1)
	if err == nil {
		t.Error("Max -> Expected an error to be returned")
	}

	list = []int{1, 2, 3, 3, 4, 5}
	res, err := Max(list, pred1)
	if res != 5 {
		t.Error("Max -> Expected 1 to be returned but got", res)
	}

	list = []int{4, 2, 3, 4, 5, 1}
	res, err = Max(list, pred1)
	if res != 5 {
		t.Error("Max -> Expected 1 to be returned but got", res)
	}

	list = []int{2, 3, 1, 1, 1, 4, 5}
	res, err = Max(list, pred1)
	if res != 5 {
		t.Error("Max -> Expected 1 to be returned but got", res)
	}

	res, err = Max(list, pred2)
	if err == nil {
		t.Error("Max -> Expected an error to be returned")
	}
}
