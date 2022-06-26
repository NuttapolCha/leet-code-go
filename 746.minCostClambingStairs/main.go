package main

import "fmt"

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}

func minCostClimbingStairs(cost []int) int {
	dp0 := cost[len(cost)-1]
	dp1 := cost[len(cost)-2]

	if len(cost) <= 2 {
		return min(dp0, dp1)
	}

	// bottom up DP but the bottom is the last index
	var next int
	for i := len(cost) - 3; i >= 0; i-- {
		next = cost[i] + min(dp0, dp1)
		dp0 = dp1
		dp1 = next
	}

	// fmt.Println(dp)
	return min(next, dp0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
