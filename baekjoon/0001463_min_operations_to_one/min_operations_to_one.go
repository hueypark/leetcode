package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	calculator := NewMinOperationToOneCalculator()

	fmt.Println(calculator.Calc(x))
}

type MinOperationToOneCalculator struct {
	memo map[int]int
}

func NewMinOperationToOneCalculator() *MinOperationToOneCalculator {
	return &MinOperationToOneCalculator{
		memo: make(map[int]int),
	}
}

func (calc *MinOperationToOneCalculator) Calc(x int) int {
	if x == 1 {
		return 0
	}

	v, ok := calc.memo[x]
	if ok {
		return v
	}

	var res int

	if x%6 == 0 { // If x is divisible by 2 and 3
		res = min(calc.Calc(x/3), calc.Calc(x/2), calc.Calc(x-1)) + 1
	} else if x%3 == 0 {
		res = min(calc.Calc(x/3), calc.Calc(x-1)) + 1
	} else if x%2 == 0 {
		res = min(calc.Calc(x/2), calc.Calc(x-1)) + 1
	} else {
		res = calc.Calc(x-1) + 1
	}

	calc.memo[x] = res
	return res
}
