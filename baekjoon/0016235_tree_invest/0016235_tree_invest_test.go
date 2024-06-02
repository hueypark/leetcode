package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpring(t *testing.T) {
	tcs := []struct {
		trees             []Tree
		land              map[Vector]int
		expectedTrees     []Tree
		expectedDeadTrees []Tree
	}{
		{
			trees:             []Tree{{Location: Vector{0, 0}, Age: 1}, {Location: Vector{0, 1}, Age: 3}, {Location: Vector{1, 0}, Age: 2}},
			land:              map[Vector]int{{0, 0}: 5, {0, 1}: 5, {1, 0}: 5},
			expectedTrees:     []Tree{{Location: Vector{0, 0}, Age: 2}, {Location: Vector{0, 1}, Age: 4}, {Location: Vector{1, 0}, Age: 3}},
			expectedDeadTrees: nil,
		},
		{
			trees:             []Tree{{Location: Vector{0, 0}, Age: 1}, {Location: Vector{0, 1}, Age: 3}, {Location: Vector{1, 0}, Age: 2}},
			land:              map[Vector]int{{0, 0}: 5, {0, 1}: 5},
			expectedTrees:     []Tree{{Location: Vector{0, 0}, Age: 2}, {Location: Vector{0, 1}, Age: 4}},
			expectedDeadTrees: []Tree{{Location: Vector{1, 0}, Age: 2}},
		},
		{
			trees:             []Tree{{Location: Vector{0, 0}, Age: 4}, {Location: Vector{1, 1}, Age: 5}},
			land:              map[Vector]int{{0, 0}: 10, {1, 1}: 10},
			expectedTrees:     []Tree{{Location: Vector{0, 0}, Age: 5}, {Location: Vector{1, 1}, Age: 6}},
			expectedDeadTrees: nil,
		},
	}

	for _, tc := range tcs {
		trees, deadTrees := spring(tc.trees, tc.land)

		require.Equal(t, sortTree(tc.expectedTrees), sortTree(trees), "Spring: Trees did not match expected value")
		require.Equal(t, sortTree(tc.expectedDeadTrees), sortTree(deadTrees), "Spring: Dead trees did not match expected value")
	}
}

func TestSummer(t *testing.T) {
	tcs := []struct {
		deadTrees    []Tree
		land         map[Vector]int
		expectedLand map[Vector]int
	}{
		{
			deadTrees:    []Tree{{Location: Vector{0, 1}, Age: 3}},
			land:         map[Vector]int{{0, 0}: 5, {0, 1}: 5, {1, 0}: 5},
			expectedLand: map[Vector]int{{0, 0}: 5, {0, 1}: 6, {1, 0}: 5},
		},
		{
			deadTrees:    []Tree{{Location: Vector{0, 0}, Age: 2}, {Location: Vector{1, 1}, Age: 4}},
			land:         map[Vector]int{{0, 0}: 5, {1, 1}: 5},
			expectedLand: map[Vector]int{{0, 0}: 6, {1, 1}: 7},
		},
	}

	for _, tc := range tcs {
		summer(tc.deadTrees, tc.land)
		require.Equal(t, tc.expectedLand, tc.land, "Summer: Land did not match expected value")
	}
}

func TestFall(t *testing.T) {
	tcs := []struct {
		trees            []Tree
		landSize         int
		expectedNewTrees []Tree
	}{
		{
			trees:            []Tree{{Location: Vector{0, 0}, Age: 5}},
			landSize:         2,
			expectedNewTrees: []Tree{{Location: Vector{0, 1}, Age: 1}, {Location: Vector{1, 0}, Age: 1}, {Location: Vector{1, 1}, Age: 1}},
		},
		{
			trees:            []Tree{{Location: Vector{1, 1}, Age: 10}},
			landSize:         3,
			expectedNewTrees: []Tree{{Location: Vector{0, 0}, Age: 1}, {Location: Vector{0, 1}, Age: 1}, {Location: Vector{0, 2}, Age: 1}, {Location: Vector{1, 0}, Age: 1}, {Location: Vector{1, 2}, Age: 1}, {Location: Vector{2, 0}, Age: 1}, {Location: Vector{2, 1}, Age: 1}, {Location: Vector{2, 2}, Age: 1}},
		},
	}

	for _, tc := range tcs {
		newTrees := fall(tc.trees, tc.landSize)
		require.Equal(t, tc.expectedNewTrees, newTrees, "Fall: New trees did not match expected value")
	}
}

func TestWinter(t *testing.T) {
	tcs := []struct {
		land         map[Vector]int
		A            [][]int
		expectedLand map[Vector]int
	}{
		{
			land:         map[Vector]int{{0, 0}: 5, {0, 1}: 5, {1, 0}: 5, {1, 1}: 5},
			A:            [][]int{{1, 2}, {3, 4}},
			expectedLand: map[Vector]int{{0, 0}: 6, {0, 1}: 7, {1, 0}: 8, {1, 1}: 9},
		},
		{
			land:         map[Vector]int{{0, 0}: 0, {0, 1}: 0, {1, 0}: 0, {1, 1}: 0},
			A:            [][]int{{10, 20}, {30, 40}},
			expectedLand: map[Vector]int{{0, 0}: 10, {0, 1}: 20, {1, 0}: 30, {1, 1}: 40},
		},
	}

	for _, tc := range tcs {
		winter(tc.land, tc.A)
		require.Equal(t, tc.expectedLand, tc.land, "Winter: Land did not match expected value")
	}
}

func TestTreeCountAfterNYears(t *testing.T) {
	tcs := []struct {
		landSize      int
		fertilizers   [][]int
		trees         []Tree
		year          int
		expectedCount int
	}{
		{
			landSize: 1,
			fertilizers: [][]int{
				{1},
			},
			trees:         []Tree{{Location: Vector{0, 0}, Age: 1}},
			year:          1,
			expectedCount: 1,
		},
		{
			landSize: 1,
			fertilizers: [][]int{
				{1},
			},
			trees:         []Tree{{Location: Vector{0, 0}, Age: 1}},
			year:          4,
			expectedCount: 0,
		},
		{
			landSize: 5,
			fertilizers: [][]int{
				{2, 3, 2, 3, 2},
				{2, 3, 2, 3, 2},
				{2, 3, 2, 3, 2},
				{2, 3, 2, 3, 2},
				{2, 3, 2, 3, 2},
			},
			trees:         []Tree{{Location: Vector{1, 0}, Age: 3}, {Location: Vector{2, 1}, Age: 3}},
			year:          1,
			expectedCount: 2,
		},
		{
			landSize: 5,
			fertilizers: [][]int{
				{2, 3, 2, 3, 2},
				{2, 3, 2, 3, 2},
				{2, 3, 2, 3, 2},
				{2, 3, 2, 3, 2},
				{2, 3, 2, 3, 2},
			},
			trees:         []Tree{{Location: Vector{1, 0}, Age: 3}, {Location: Vector{2, 1}, Age: 3}},
			year:          6,
			expectedCount: 85,
		},
		{
			landSize: 10,
			fertilizers: [][]int{
				{100, 100, 100, 100, 100, 100, 100, 100, 100, 100},
				{100, 100, 100, 100, 100, 100, 100, 100, 100, 100},
				{100, 100, 100, 100, 100, 100, 100, 100, 100, 100},
				{100, 100, 100, 100, 100, 100, 100, 100, 100, 100},
				{100, 100, 100, 100, 100, 100, 100, 100, 100, 100},
				{100, 100, 100, 100, 100, 100, 100, 100, 100, 100},
				{100, 100, 100, 100, 100, 100, 100, 100, 100, 100},
				{100, 100, 100, 100, 100, 100, 100, 100, 100, 100},
				{100, 100, 100, 100, 100, 100, 100, 100, 100, 100},
				{100, 100, 100, 100, 100, 100, 100, 100, 100, 100},
			},
			trees:         []Tree{{Location: Vector{1, 1}, Age: 1}},
			year:          1000,
			expectedCount: 5258,
		},
	}

	for _, tc := range tcs {
		land := make(map[Vector]int)
		for x := range tc.landSize {
			for y := range tc.landSize {
				land[Vector{x, y}] = 5
			}
		}

		count := treeCountAfterNYears(tc.trees, tc.landSize, land, tc.fertilizers, tc.year)
		require.Equal(t, tc.expectedCount, count, "Tree count after %d years did not match expected value", tc.year)
	}
}

func sortTree(trees []Tree) []Tree {
	sort.Slice(trees, func(i, j int) bool {
		if trees[i].Location.X == trees[j].Location.X {
			return trees[i].Location.Y < trees[j].Location.Y
		}

		return trees[i].Location.X < trees[j].Location.X
	})

	return trees
}
