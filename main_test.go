package main

import "testing"

func TestSqrt(t *testing.T) {
	const in, out = 4, 2
	if x := sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	}
}