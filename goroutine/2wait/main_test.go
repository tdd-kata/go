package main

import "testing"

// go test .
func TestSumAtoB(t *testing.T) {
	sum := SumAtoB(2, 4)
	expected := 9

	if sum != expected {
		t.Errorf("기대하는 값은 %d, 실제 값은 %d", expected, sum)
	}
}
