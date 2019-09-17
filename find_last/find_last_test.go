package find_last

//go:generate fungo -template=FindLast -fileName=find_last

import (
	"errors"
	"testing"
)

func pred1(x string) (bool, error) {
	return x[0] == 'A', nil
}

func pred2(a, b string) (bool, error) {
	return false, errors.New("this is wrong")
}

func TestMin(t *testing.T) {
}
