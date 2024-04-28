package main

import (
	"fmt"
)

func main() {
	var numTest int
	fmt.Scan(&numTest)

	for i := 0; i < numTest; i++ {
		field := readInput()

		fmt.Println(field.CaculateRequiredPestPreventRobotCount())
	}
}

type Field struct {
	Size     Vector
	Cabbages map[Vector]struct{}
}

type Vector struct {
	X int
	Y int
}

func (field Field) CaculateRequiredPestPreventRobotCount() int {
	var count int
	visited := make(map[Vector]struct{})

	for cabbageLoc := range field.Cabbages {
		_, ok := visited[cabbageLoc]
		if ok {
			continue
		}

		count++

		field.markVisited(visited, cabbageLoc)
	}

	return count
}

func (field Field) markVisited(visited map[Vector]struct{}, loc Vector) {
	if loc.X < 0 || loc.X >= field.Size.X || loc.Y < 0 || loc.Y >= field.Size.Y {
		return
	}

	_, ok := visited[loc]
	if ok {
		return
	}

	visited[loc] = struct{}{}

	_, ok = field.Cabbages[loc]
	if !ok {
		return
	}

	field.markVisited(visited, Vector{loc.X - 1, loc.Y})
	field.markVisited(visited, Vector{loc.X + 1, loc.Y})
	field.markVisited(visited, Vector{loc.X, loc.Y - 1})
	field.markVisited(visited, Vector{loc.X, loc.Y + 1})
}

func readInput() Field {
	field := Field{
		Cabbages: make(map[Vector]struct{}),
	}
	var cabbagesNum int
	fmt.Scan(&field.Size.X, &field.Size.Y, &cabbagesNum)

	for i := 0; i < cabbagesNum; i++ {
		var cabbagesLocation Vector
		fmt.Scan(&cabbagesLocation.X, &cabbagesLocation.Y)
		field.Cabbages[cabbagesLocation] = struct{}{}
	}

	return field
}
