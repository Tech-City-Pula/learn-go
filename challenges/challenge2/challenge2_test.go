package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	var sum1 = Sum(2, 3)
	if sum1 != 5 {
		t.Errorf("Sum(2, 3) = %d; want 5", sum1)
	}

	var sum2 = Sum(7, -2)
	if sum2 != 5 {
		t.Errorf("Sum(7, -2) = %d; want 5", sum2)
	}
}
