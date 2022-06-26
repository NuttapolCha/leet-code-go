package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	// nums := []int{1, 0}
	// target := 1
	fmt.Println(findTargetSumWays(nums, target))
}

func findTargetSumWays(nums []int, target int) int {
	return targetSum(nums, 0, target)
}

func targetSum(nums []int, startFrom, target int) int {
	if startFrom == len(nums)-1 {
		if nums[startFrom]+target == 0 && nums[startFrom]-target == 0 {
			return 2
		}
		if nums[startFrom]+target == 0 || nums[startFrom]-target == 0 {
			return 1
		}
		return 0
	}

	return targetSum(nums, startFrom+1, target+nums[startFrom]) + targetSum(nums, startFrom+1, target-nums[startFrom])
}
