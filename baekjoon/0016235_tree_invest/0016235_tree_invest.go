package main

import (
	"fmt"
	"sort"
)

type Tree struct {
	Location Vector
	Age      int
}

type Vector struct {
	X int
	Y int
}

func spring(trees []Tree, land map[Vector]int) ([]Tree, []Tree) {
	var deadTrees []Tree
	sort.Slice(trees, func(i, j int) bool {
		return trees[i].Age < trees[j].Age
	})

	var aliveTrees []Tree
	for _, t := range trees {
		vec := t.Location
		if land[vec] >= t.Age {
			land[vec] -= t.Age
			t.Age++
			aliveTrees = append(aliveTrees, t)
		} else {
			deadTrees = append(deadTrees, t)
		}
	}
	return aliveTrees, deadTrees
}

func summer(deadTrees []Tree, land map[Vector]int) {
	for _, t := range deadTrees {
		vec := t.Location
		land[vec] += t.Age / 2
	}
}

func fall(trees []Tree, landSize int) []Tree {
	var newTrees []Tree
	for _, t := range trees {
		if t.Age%5 == 0 {
			for _, d := range []Vector{
				{-1, -1}, {-1, 0}, {-1, 1},
				{0, -1}, {0, 1},
				{1, -1}, {1, 0}, {1, 1},
			} {
				nx, ny := t.Location.X+d.X, t.Location.Y+d.Y
				if nx >= 0 && nx < landSize && ny >= 0 && ny < landSize {
					newTrees = append(newTrees, Tree{Location: Vector{nx, ny}, Age: 1})
				}
			}
		}
	}
	return newTrees
}

func winter(land map[Vector]int, fertilizers [][]int) {
	for i := 0; i < len(fertilizers); i++ {
		for j := 0; j < len(fertilizers[i]); j++ {
			vec := Vector{i, j}
			land[vec] += fertilizers[i][j]
		}
	}
}

func main() {
	var N, M, K int
	fmt.Scanf("%d %d %d", &N, &M, &K)

	fertilizers := make([][]int, N)
	for i := range fertilizers {
		fertilizers[i] = make([]int, N)
		for j := range fertilizers[i] {
			fmt.Scanf("%d", &fertilizers[i][j])
		}
	}

	land := make(map[Vector]int)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			land[Vector{i, j}] = 5
		}
	}

	trees := make([]Tree, M)
	for i := 0; i < M; i++ {
		var x, y, age int
		fmt.Scanf("%d %d %d", &x, &y, &age)
		trees[i] = Tree{Location: Vector{x - 1, y - 1}, Age: age}
	}

	var deadTrees []Tree
	for year := 0; year < K; year++ {
		trees, deadTrees = spring(trees, land)
		summer(deadTrees, land)
		newTrees := fall(trees, N)

		trees = append(trees, newTrees...)
		winter(land, fertilizers)
	}

	fmt.Println(len(trees))
}
