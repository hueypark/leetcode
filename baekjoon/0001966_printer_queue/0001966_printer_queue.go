package main

import (
	"fmt"
	"slices"
)

func main() {

}

func CalcPrintOrder(priorities []int, idx int, iterCount int) (int, error) {
	if len(priorities) == 0 {
		return 0, fmt.Errorf("priorities is empty")
	}

	if len(priorities) == 1 {
		return iterCount + 1, nil
	}

	firstPrior := priorities[0]

	hasMorePrior := slices.ContainsFunc(priorities[1:], func(p int) bool {
		return p > firstPrior
	})
	if !hasMorePrior {
		if idx == 0 {
			return iterCount + 1, nil
		} else {
			idx--
			return CalcPrintOrder(priorities[1:], idx, iterCount+1)
		}
	}

	priorities = append(priorities[1:], firstPrior)

	if idx == 0 {
		idx = len(priorities) - 1
	} else {
		idx--
	}

	return CalcPrintOrder(priorities, idx, iterCount)
}
