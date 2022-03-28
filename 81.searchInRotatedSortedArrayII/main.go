package main

import (
	"fmt"
)

func main() {
	input := []int{2, 5, 6, 0, 0, 1, 2}
	// input := []int{5}
	fmt.Println(search(input, 5))
}

func search(nums []int, target int) bool {
	var prev *int
	for i, num := range nums {
		if prev == nil {
			prev = new(int)
		} else {
			if num < *prev {
				return bSearch(append(nums[i:], nums[0:i]...), target)
			}
		}
		*prev = num
	}
	return bSearch(nums, target)
}

func bSearch(sortedNums []int, target int) bool {
	if len(sortedNums) == 0 {
		return false
	}

	leftIndex := 0
	rightIndex := len(sortedNums) - 1
	midIndex := (leftIndex + rightIndex) / 2
	midValue := sortedNums[midIndex]

	if rightIndex < leftIndex {
		return false
	}
	if rightIndex == leftIndex { // e.g. == midIndex
		return target == midValue
	}
	if target < midValue {
		return bSearch(sortedNums[leftIndex:midIndex], target)
	}
	if target > midValue {
		return bSearch(sortedNums[midIndex+1:], target)
	}

	return target == midValue
}
