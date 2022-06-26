package main

import "fmt"

func main() {
	fmt.Println(tribonacci(25))
}

func tribonacci(n int) int {
	dp0 := 0
	dp1 := 1
	dp2 := 1

	if n == 0 {
		return dp0
	}
	if n <= 2 {
		return dp1
	}

	var next int
	for i := 3; i <= n; i++ {
		next = dp0 + dp1 + dp2
		dp0 = dp1
		dp1 = dp2
		dp2 = next
	}

	return next
}
