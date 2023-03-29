package main

import (
	"testing"
)

func TestAddComplex(t *testing.T) {

	tests := []struct {
		lhs      int
		rhs      int
		expected int
	}{
		{lhs: 1, rhs: 2, expected: 3},
		{lhs: 2, rhs: 2, expected: 4},
		{lhs: 2, rhs: 2, expected: 6},
		{lhs: 10, rhs: -5, expected: 5},
	}

	for _, test := range tests {
		ans := add(test.lhs, test.rhs)
		if ans != test.expected {
			t.Errorf("add(%d,%d) = %d . Wanted %d", test.lns, test.rhs, ans, test.expected)
		}

	}

}
func TestAddComplex2(t *testing.T) {
	a := add(2, 2)
	if a != 4 {
		t.Errorf("a := add(1,2) : is not 3. Got %d", a)
	}

}
