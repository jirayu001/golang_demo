package coveragedemo

import (
	"math"
	"testing"
)

func TestBasicOperation(t *testing.T) {
	x := operation(1, 2, '+')
	if x != 3 {
		t.Errorf("operation(%f,%f,+) = %f. Wanted : %f", 1.0, 2.0, x, 3.0)
	}
	x = operation(1, 2, '-')
	if x != -1 {
		t.Errorf("operation(%f,%f,-) = %f. Wanted : %f", 1.0, 2.0, x, -1.0)
	}
	x = operation(1.0, 3.0, '/')
	epsilon := math.Nextafter(1.0, 2.0) - 1.0
	if math.Abs(x-(1.0/3.0)) > epsilon { //0.3333333 > 0.00333333
		t.Errorf("operation(%f,%f,/) = %f. Wanted : %f", 1.0, 3.0, x, 0.3333)
	}
}
