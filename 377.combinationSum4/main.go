package main

import "fmt"

func main() {
	nums := []int{2, 1, 3}
	target := 35
	fmt.Println(combinationSum4(nums, target))
}

func combinationSum4(nums []int, target int) int {
	// bottom up DP
	dp := make([]int, target+1)

	// smallest problem => dp[0] = 0

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num == 0 {
				dp[i] += 1
			} else if i-num > 0 {
				dp[i] += dp[i-num]
			}
		}
	}

	// fmt.Println(dp)
	return dp[target]
}
