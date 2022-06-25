package main

import (
	"fmt"
)

func main() {
	coins := []int{3, 4, 5}
	fmt.Println(coinChange(coins, 7))
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	// bottom up DP from 1 -> amount
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		var min *int
		for _, coinVal := range coins {
			if i >= coinVal && dp[i-coinVal] != -1 {
				numCoins := 1 + dp[i-coinVal]
				if min == nil || (min != nil && numCoins < *min) {
					min = &numCoins
				}
			}
		}
		if min == nil {
			dp[i] = -1
		} else {
			dp[i] = *min
		}
		// fmt.Printf("dp[%d] = %d\n", i, dp[i])
	}

	return dp[amount]
}
