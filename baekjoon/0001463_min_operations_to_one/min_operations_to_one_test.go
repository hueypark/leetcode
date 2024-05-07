package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinOperationToOne(t *testing.T) {
	tcs := []struct {
		x        int
		expected int
	}{
		{2, 1},
		{10, 3},
	}

	for _, tc := range tcs {
		calculator := NewMinOperationToOneCalculator()
		actual := calculator.Calc(tc.x)
		require.Equalf(t, tc.expected, actual, "x: %d", tc.x)
	}
}
