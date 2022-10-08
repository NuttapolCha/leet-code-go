package main

import "fmt"

func main() {
	amount := 3
	// coins := []int{1, 2, 5}
	coins := []int{2}
	fmt.Println(change(amount, coins))
}

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if i-coin >= 0 {
				dp[i] += dp[i-coin]
			}
		}
	}

	return dp[amount]
}
