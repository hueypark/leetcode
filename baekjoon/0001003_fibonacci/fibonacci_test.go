package main

import "testing"

func TestFibonacciCounter(t *testing.T) {
	tcs := []struct {
		n        int
		counter0 int
		counter1 int
	}{
		{0, 1, 0},
		{1, 0, 1},
		{2, 1, 1},
		{3, 1, 2},
	}

	for _, tc := range tcs {
		counter0, counter1 := FibonacciCounter(tc.n)
		if counter0 != tc.counter0 || counter1 != tc.counter1 {
			t.Errorf("Expected (%v, %v) but got (%v, %v)", tc.counter0, tc.counter1, counter0, counter1)
		}
	}
}
