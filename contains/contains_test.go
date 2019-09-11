package contains

//go:generate fungo -template=Contains -fileName=contains

import (
	"errors"
	"testing"
)

func eqFunc(a, b string) (bool, error) {
	return a == b, nil
}

func eqFuncErr(a, b string) (bool, error) {
	return false, errors.New("this is wrong")
}

func TestContains(t *testing.T) {
	// Empty list 1
	r, i, e := Contains([]string{}, []string{"aa", "bb"}, eqFunc)
	if r || i >= 0 || e != nil {
		t.Error("Contains -> Empty list 1 failed")
	}

	// Empty list 2
	r, i, e = Contains([]string{"aa", "bb", "cc"}, []string{}, eqFunc)
	if r || i >= 0 || e != nil {
		t.Error("Contains -> Empty list 2 failed")
	}

	// One element in list 1 and 2 positive
	r, i, e = Contains([]string{"aa"}, []string{"aa"}, eqFunc)
	if !r || i < 0 || e != nil {
		t.Error("Contains -> One element in list 1 and 2 failed positive failed")
	}

	// One element in list 1 and 2 negative
	r, i, e = Contains([]string{"aa"}, []string{"bb"}, eqFunc)
	if r || i >= 0 || e != nil {
		t.Error("Contains -> One element in list 1 and 2 failed negative failed")
	}

	// Many elements in list 1 one element in list 2 positive
	r, i, e = Contains([]string{"aa", "bb", "cc"}, []string{"aa"}, eqFunc)
	if !r || i < 0 || e != nil {
		t.Error("Contains -> Many elements in list 1 one element in list 2 positive failed")
	}
	r, i, e = Contains([]string{"aa", "bb", "cc"}, []string{"bb"}, eqFunc)
	if !r || i < 0 || e != nil {
		t.Error("Contains -> Many elements in list 1 one element in list 2 positive failed")
	}
	r, i, e = Contains([]string{"aa", "bb", "cc"}, []string{"cc"}, eqFunc)
	if !r || i < 0 || e != nil {
		t.Error("Contains -> Many elements in list 1 one element in list 2 positive failed")
	}

	// Many elements in list 1 one element in list 2 negative
	r, i, e = Contains([]string{"aa", "bb", "cc"}, []string{"dd"}, eqFunc)
	if r || i >= 0 || e != nil {
		t.Error("Contains -> Many elements in list 1 one element in list 2 negative failed")
	}

	// Many elements in list 1 and 2 positive
	r, i, e = Contains([]string{"aa", "bb", "cc", "dd", "ee"}, []string{"aa", "bb", "cc"}, eqFunc)
	if !r || i != 0 || e != nil {
		t.Error("Contains -> Many elements in list 1 and 2 positive failed")
	}
	r, i, e = Contains([]string{"aa", "bb", "cc", "dd", "ee"}, []string{"bb", "cc", "dd"}, eqFunc)
	if !r || i != 1 || e != nil {
		t.Error("Contains -> Many elements in list 1 and 2 positive failed")
	}
	r, i, e = Contains([]string{"aa", "bb", "cc", "dd", "ee"}, []string{"cc", "dd", "ee"}, eqFunc)
	if !r || i != 2 || e != nil {
		t.Error("Contains -> Many elements in list 1 and 2 positive failed")
	}
	r, i, e = Contains([]string{"aa", "bb", "cc", "dd", "ee"}, []string{"aa", "bb", "cc", "dd", "ee"}, eqFunc)
	if !r || i != 0 || e != nil {
		t.Error("Contains -> Many elements in list 1 and 2 positive failed")
	}
	r, i, e = Contains([]string{"aa", "aa", "aa", "aa", "bb"}, []string{"aa", "bb"}, eqFunc)
	if !r || i != 3 || e != nil {
		t.Error("Contains -> Many elements in list 1 and 2 positive failed")
	}

	// Many elements in list 1 and 2 negative
	r, i, e = Contains([]string{"aa", "bb", "cc", "dd", "ee"}, []string{"aa", "bb", "dd"}, eqFunc)
	if r || i >= 0 || e != nil {
		t.Error("Contains -> Many elements in list 1 and 2 negative failed")
	}
	r, i, e = Contains([]string{"aa", "bb", "cc", "dd", "ee"}, []string{"bb", "cc", "ee"}, eqFunc)
	if r || i >= 0 || e != nil {
		t.Error("Contains -> Many elements in list 1 and 2 negative failed")
	}
	r, i, e = Contains([]string{"aa", "bb", "cc", "dd", "ee"}, []string{"cc", "dd", "ee", "ff"}, eqFunc)
	if r || i >= 0 || e != nil {
		t.Error("Contains -> Many elements in list 1 and 2 negative failed")
	}

	// Negative with error
	r, i, e = Contains([]string{"aa", "bb", "cc", "dd", "ee"}, []string{"bb", "cc", "ee"}, eqFuncErr)
	if e == nil {
		t.Error("Contains -> Negative with error failed")
	}
}
