package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrganicCabbage(t *testing.T) {
	tcs := []struct {
		field Field
		count int
	}{
		{
			field: Field{
				Size: Vector{10, 6},
				Cabbages: map[Vector]struct{}{
					{0, 0}: {},
					{1, 0}: {},
					{1, 1}: {},
					{2, 4}: {},
					{3, 4}: {},
					{4, 2}: {},
					{4, 3}: {},
					{4, 5}: {},
					{7, 4}: {},
					{7, 5}: {},
					{8, 4}: {},
					{8, 5}: {},
					{9, 4}: {},
					{9, 5}: {},
				},
			},
			count: 5,
		},
	}

	for _, tc := range tcs {
		count := tc.field.CaculateRequiredPestPreventRobotCount()
		require.Equal(t, tc.count, count)
	}
}
