package main

import "fmt"

func main() {
	nums := []int{2, 7, 9, 3, 1}
	// nums := []int{1, 2, 3, 1}
	// nums := []int{2, 1, 1, 2}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]

	// notAdjaventMaximumDp is the max dp[j] where j < i-1
	notAdjacentMaximumDp := dp[0]
	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i]+notAdjacentMaximumDp, dp[i-1])

		// update notAdjaventMaximumDp
		if dp[i-1] > notAdjacentMaximumDp {
			notAdjacentMaximumDp = dp[i-1]
		}
	}

	// fmt.Println(dp)
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
