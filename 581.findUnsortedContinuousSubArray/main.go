package main

import (
	"fmt"
)

func main() {
	// nums := []int{2, 6, 4, 8, 10, 9, 15}
	// nums := []int{1, 2, 3, 4, 7, 5, 3, 2, 4, 5, 6, 7, 4, 6, 4}
	// nums := []int{1, 2, 3, 3, 8, 9, 10, 5, 6, 11}
	// nums := []int{}
	nums := []int{1, 2, 3, 4}
	fmt.Println(findUnsortedSubarray(nums))
}

func findUnsortedSubarray(nums []int) int {
	// guarantee that nums has at least 1 element.
	prevMax := nums[0]
	var start, end int = -1, -1
	for i := 0; i < len(nums); i++ {
		// current num is less than previous maximum
		if nums[i] < prevMax {
			if start < 0 { // only assign start 1 time
				start = i - 1
			}
			end = i
		}

		// adjust start index
		isAdjusted := false
		for start >= 0 && end >= 0 && nums[end] < nums[start] {
			start--
			isAdjusted = true
		}

		// before entering next iteration
		if nums[i] > prevMax {
			prevMax = nums[i]
		}
		if isAdjusted {
			start++
		}
	}

	// fmt.Printf("start = %d, end = %d\n", start, end)
	// fmt.Printf("must sort: %v\n", nums[start:end+1])

	if end < 0 || start < 0 {
		return 0
	}
	return end - start + 1
}
