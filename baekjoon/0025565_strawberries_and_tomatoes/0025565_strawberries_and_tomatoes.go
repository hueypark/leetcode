package main

import (
	"errors"
	"fmt"
	"slices"
)

func main() {
	f := readInput()

	duplicateSeedPlantedLocations, err := f.FindDuplicateSeedPlantedLocations()
	if err != nil {
		panic(err)
	}

	fmt.Println(len(duplicateSeedPlantedLocations))
	for _, loc := range duplicateSeedPlantedLocations {
		fmt.Println(loc.X, loc.Y)
	}
}

type Field struct {
	rowLen               int
	seedPlantedLocations []Vector
}

type Vector struct {
	X int
	Y int
}

type Row struct {
	RowGroup  RowGroup
	Locations []Vector
}

type RowGroup int

const (
	RowGroupSameX RowGroup = iota
	RowGroupSameY
)

var (
	ErrNoRowsFound       = errors.New("no rows found")
	ErrInvalidRowLength  = errors.New("invalid row length")
	ErrInvalidRowsLength = errors.New("invalid rows length")
)

func (f *Field) FindDuplicateSeedPlantedLocations() ([]Vector, error) {
	rows := f.findRows()
	if len(rows) == 0 {
		return nil, ErrNoRowsFound
	}

	switch len(rows) {
	case 1:
		row := rows[0]

		slices.SortFunc(row.Locations, func(a, b Vector) int {
			if a.X == b.X {
				return a.Y - b.Y
			}

			return a.X - b.X
		})

		locataionsLen := len(row.Locations)
		if locataionsLen == f.rowLen {
			return row.Locations, nil
		} else if f.rowLen < locataionsLen {
			diff := locataionsLen - f.rowLen
			halfDiff := diff / 2

			return row.Locations[halfDiff : locataionsLen-halfDiff], nil
		} else {
			return nil, ErrInvalidRowLength
		}
	case 2:
		row1 := rows[0]
		row2 := rows[1]

		if row1.RowGroup == row2.RowGroup {
			return nil, nil
		}

		for _, loc1 := range row1.Locations {
			for _, loc2 := range row2.Locations {
				if loc1 == loc2 {
					return []Vector{loc1}, nil
				}
			}
		}

		return nil, nil
	default:
		return nil, ErrInvalidRowsLength
	}
}

func (f *Field) findRows() []Row {
	sameXCandidates := make(map[int][]Vector)
	sameYCandidates := make(map[int][]Vector)
	for _, loc := range f.seedPlantedLocations {
		sameXCandidates[loc.X] = append(sameXCandidates[loc.X], loc)
		sameYCandidates[loc.Y] = append(sameYCandidates[loc.Y], loc)
	}

	var rows []Row
	for _, row := range sameXCandidates {
		if len(row) >= f.rowLen {
			rows = append(rows, Row{
				RowGroup:  RowGroupSameX,
				Locations: row,
			})
		}
	}

	if f.rowLen == 1 {
		return rows
	}

	if len(rows) >= 2 {
		return rows
	}

	// TODO: Handle more then 2 rows case.

	for _, row := range sameYCandidates {
		if len(row) >= f.rowLen {
			rows = append(rows, Row{
				RowGroup:  RowGroupSameY,
				Locations: row,
			})
		}
	}

	return rows
}

func readInput() Field {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	f := Field{
		rowLen: n,
	}

	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			var exists bool
			fmt.Scan(&exists)

			if !exists {
				continue
			}

			f.seedPlantedLocations = append(f.seedPlantedLocations, Vector{
				X: x,
				Y: y,
			})
		}

	}

	return f
}
