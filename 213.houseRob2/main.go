package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 1}
	nums := []int{0}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(houseRob1(nums[1:]), houseRob1(nums[:len(nums)-1]))
}

func houseRob1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	dp0 := nums[0]
	dp1 := max(nums[0], nums[1])

	var next int
	for i := 2; i < len(nums); i++ {
		next = max(nums[i]+dp0, dp1)

		// update notAdjaventMaximumDp
		if dp1 > dp0 {
			dp0 = dp1
		}

		dp0 = dp1
		dp1 = next
	}

	return next
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
