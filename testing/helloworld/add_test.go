package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := add(1, 2)
	if a != 3 {
		t.Errorf("a := add(1,2) : is not 3. Got %d", a)
	}

}
func TestAdd2(t *testing.T) {
	a := add(2, 2)
	if a != 4 {
		t.Errorf("a := add(1,2) : is not 3. Got %d", a)
	}

}
