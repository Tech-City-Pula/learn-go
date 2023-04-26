package main

import (
	"testing"
)

func TestSumInts(t *testing.T) {
	sum := SumInts(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if sum != 55 {
		t.Errorf("SumInts(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) = %d; want 55", sum)
	}

	sum = SumInts(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	if sum != 66 {
		t.Errorf("SumInts(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11) = %d; want 66", sum)
	}

	sum = SumInts(1, 2, 3, 4, 5, 6)
	if sum != 21 {
		t.Errorf("SumInts(1, 2, 3, 4, 5, 6) = %d; want 21", sum)
	}

	sum = SumInts(1, 2)
	if sum != 3 {
		t.Errorf("SumInts(1, 2) = %d; want 3", sum)
	}
}
