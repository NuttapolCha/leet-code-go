package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fib(i))
	}
}

func fib(n int) int {
	// Use bottom up dynamic programing approach

	// bottom condition
	prev0 := 0 // F(0)
	prev1 := 1 // F(1)

	if n == 0 {
		return prev0
	}
	if n == 1 {
		return prev0 + prev1
	}

	var ret int
	for i := 2; i <= n; i++ {
		ret = prev0 + prev1
		prev0 = prev1
		prev1 = ret
	}

	return ret
}
