package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindDuplicateSeedPlantedLocations(t *testing.T) {
	tcs := []struct {
		name                          string
		field                         Field
		duplicateSeedPlantedLocations []Vector
	}{
		{
			name: "basic",
			field: Field{
				rowLen: 1,
				seedPlantedLocations: []Vector{
					{X: 0, Y: 1},
					{X: 1, Y: 0},
				},
			},
			duplicateSeedPlantedLocations: nil,
		},
		{
			name: "row len 2",
			field: Field{
				rowLen: 2,
				seedPlantedLocations: []Vector{
					{X: 0, Y: 0},
					{X: 0, Y: 1},
					{X: 1, Y: 0},
					{X: 1, Y: 1},
				},
			},
			duplicateSeedPlantedLocations: nil,
		},
		{
			name: "row len 2 and has only one seed planted location",
			field: Field{
				rowLen: 2,
				seedPlantedLocations: []Vector{
					{X: 0, Y: 0},
					{X: 0, Y: 1},
					{X: 1, Y: 0},
				},
			},
			duplicateSeedPlantedLocations: []Vector{
				{X: 0, Y: 0},
			},
		},
		{
			name: "row len 2 and has two seed planted location",
			field: Field{
				rowLen: 3,
				seedPlantedLocations: []Vector{
					{X: 0, Y: 0},
					{X: 0, Y: 1},
					{X: 0, Y: 2},
					{X: 0, Y: 3},
					{X: 0, Y: 4},
				},
			},
			duplicateSeedPlantedLocations: []Vector{
				{X: 0, Y: 1},
				{X: 0, Y: 2},
				{X: 0, Y: 3},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			dpliLocs, err := tc.field.FindDuplicateSeedPlantedLocations()
			require.NoError(t, err)

			require.Equal(t, tc.duplicateSeedPlantedLocations, dpliLocs)
		})
	}
}
