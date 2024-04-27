package main

import "fmt"

func main() {
	var numTest int
	fmt.Scan(&numTest)

	for i := 0; i < numTest; i++ {
		var n int
		fmt.Scan(&n)
		fmt.Println(FibonacciCounter(n))
	}
}

func FibonacciCounter(n int) (int, int) {
	if n == 0 {
		return 1, 0
	}
	if n == 1 {
		return 0, 1
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return a, b
}
