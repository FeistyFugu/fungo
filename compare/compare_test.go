package compare

import "testing"

//go:generate fungo -template=Compare -fileName=compare -t1=string -t2=int

func TestCompare(t *testing.T) {
	a := []string{"Alex", "Chris", "Stephanie", "Jenny", "Bob"}
	b := []int{4, 5, 9, 5, 3}
	eq := func(x string, y int) (bool, error) {
		return len(x) == y, nil
	}
	r, _ := Compare(a, b, eq)
	if !r {
		t.Error("Equals -> Expected", a, "to match", b)
	}

	b = []int{1, 5, 9, 5, 3}
	r, _ = Compare(a, b, eq)
	if r {
		t.Error("Equals -> Expected", a, "to not match", b)
	}

	b = []int{4, 5, 9, 5, 2}
	r, _ = Compare(a, b, eq)
	if r {
		t.Error("Equals -> Expected", a, "to not match", b)
	}

	b = []int{4, 5, 9, 6, 3}
	r, _ = Compare(a, b, eq)
	if r {
		t.Error("Equals -> Expected", a, "to not match", b)
	}

	b = []int{4, 5, 9, 6}
	r, _ = Compare(a, b, eq)
	if r {
		t.Error("Equals -> Expected", a, "to not match", b)
	}
}