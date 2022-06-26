package main

import "fmt"

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))

	dp[len(cost)-1] = cost[len(cost)-1]
	dp[len(cost)-2] = cost[len(cost)-2]

	if len(cost) <= 2 {
		return min(dp[0], dp[1])
	}

	// bottom up DP but the bottom is the last index
	for i := len(cost) - 3; i >= 0; i-- {
		dp[i] = cost[i] + min(dp[i+1], dp[i+2])
	}

	// fmt.Println(dp)
	return min(dp[0], dp[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
