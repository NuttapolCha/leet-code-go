package main

import "fmt"

func main() {
	input := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Println(searchMatrix(input, 3))
}

func searchMatrix(matrix [][]int, target int) bool {
	nums := []int{}
	for _, row := range matrix {
		nums = append(nums, row...)
	}
	return bSearch(nums, target)
}

func bSearch(nums []int, target int) bool {
	leftIndex := 0
	rightIndex := len(nums) - 1
	midIndex := (leftIndex + rightIndex) / 2

	// check index
	if rightIndex < leftIndex {
		return false
	}
	if leftIndex == rightIndex { // e.g. == midIndex
		return target == nums[midIndex]
	}

	// check value
	if target < nums[midIndex] {
		return bSearch(nums[leftIndex:midIndex], target)
	}

	if target > nums[midIndex] {
		return bSearch(nums[midIndex+1:], target)
	}

	return target == nums[midIndex]
}
