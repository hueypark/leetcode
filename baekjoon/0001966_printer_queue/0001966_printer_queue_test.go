package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrintOrder(t *testing.T) {
	tcs := []struct {
		priorities []int
		idx        int
		printOrder int
	}{
		{
			priorities: []int{5},
			idx:        0,
			printOrder: 1,
		},
		{
			priorities: []int{1, 2, 3, 4},
			idx:        2,
			printOrder: 2,
		},
		{
			priorities: []int{1, 1, 9, 1, 1, 1},
			idx:        0,
			printOrder: 5,
		},
	}

	for _, tc := range tcs {
		printOrder, err := CalcPrintOrder(tc.priorities, tc.idx, 0)
		require.NoError(t, err)
		require.Equal(t, tc.printOrder, printOrder, "priorities: %v, idx: %d", tc.priorities, tc.idx)
	}

}
