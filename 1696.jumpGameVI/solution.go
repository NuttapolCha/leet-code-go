package solution

func maxResult(nums []int, k int) int {
	dp := make([]int, len(nums))

	// the smallest problem
	dp[len(nums)-1] = nums[len(nums)-1]
	if len(dp) == 1 {
		return dp[0]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		// we guarantee that i+1 will not out of range
		maxOnJump := dp[i+1]

		// then we check all possible score when we jump from current index i to the possible maximum i+k
		for j := 2; j <= k; j++ {
			if i+j >= len(nums) {
				break
			}
			if dp[i+j] > maxOnJump {
				maxOnJump = dp[i+j]
			}
		}

		// assign to current dp
		dp[i] = nums[i] + maxOnJump
	}

	// fmt.Printf("dp = %+v\n", dp)
	return dp[0]
}
