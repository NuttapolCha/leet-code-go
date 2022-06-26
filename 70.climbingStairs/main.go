package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	dp1 := 1
	dp2 := 2

	if n == 0 {
		return 0
	}
	if n == 1 {
		return dp1
	}
	if n == 2 {
		return dp2
	}

	// bottom up dynamic programming
	var next int
	for i := 3; i <= n; i++ {
		next = dp2 + dp1
		dp1 = dp2
		dp2 = next
	}

	return next
}
