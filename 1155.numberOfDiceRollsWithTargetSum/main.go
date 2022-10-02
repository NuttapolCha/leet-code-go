package main

import "fmt"

func main() {
	n := 30
	k := 30
	target := 500

	fmt.Println(numRollsToTarget(n, k, target))
}

func numRollsToTarget(n int, k int, target int) int {
	// smallest condition
	// 1. no more dice
	left := n - 1
	if left == 0 {
		if target <= k {
			return 1
		} else {
			return 0
		}
	}
	// or 2. we not use exactly n dices
	if target <= 0 {
		return 0
	}

	possibility := 0
	for i := 1; i <= k; i++ {
		possibility += numRollsToTarget(n-1, k, target-i)
	}

	return possibility
}
