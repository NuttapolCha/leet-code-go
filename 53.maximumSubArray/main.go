package main

import "fmt"

func main() {
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{5, 4, -1, 7, 8}
	// nums := []int{0, -1, -2}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	// using sliding window by ignoring left side nagative portion
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxVal := -10000

	curr := 0
	for i := 0; i < len(nums); i++ {
		curr += nums[i]
		maxVal = max(curr, maxVal)

		if curr < 0 && i+1 < len(nums) {
			maxVal = max(maxVal, maxSubArray(nums[i+1:]))
			break
		}
	}

	return maxVal
}
