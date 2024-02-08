package main

import (
	"fmt"
	"math"
)

func main() {
	n := 61
	fmt.Println(numSquares(n))
}

var INF = math.MaxInt32
var memo = make(map[int]int)

func numSquares(n int) int {
	return backtracking(n)
}

func backtracking(remaining int) int {
	if remaining == 0 {
		return 0
	}

	if val, ok := memo[remaining]; ok {
		return val
	}

	minSquares := INF

	for i := 1; i*i <= remaining; i++ {
		minSquares = min(minSquares, 1+backtracking(remaining-i*i))
	}

	memo[remaining] = minSquares
	return minSquares
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
