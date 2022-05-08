package main

import (
	"fmt"
	"sort"
)

func main() {
	testCases := [][]int{
		{1, 0, 1, 1, 1, 1, 1, 3, 2, 1, 0, -1}, // true
		{1, 2, 3, 4},                          // false
		{3, 1, 4, 2},                          // true
		{3, 5, 0, 3, 4},                       // true
		{-2, 1, 2, -2, 1, 2},                  // true
	}
	for _, tc := range testCases {
		fmt.Println(find132pattern(tc))
	}

}

func find132pattern(nums []int) bool {
	// fmt.Printf("unsorted = %v\n", unsort)
	// fmt.Printf("nums = %v\n", nums)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sorted := make([]int, len(nums))
			copy(sorted, nums)
			sort.Slice(sorted, func(i, j int) bool {
				return sorted[i] < sorted[j]
			})

			found := bSearchBetween(nums, sorted[i], sorted[j])
			if found {
				return true
			}
		}
	}

	return false
}

// search a num which grater than 'one' but less than 'three'
func bSearchBetween(nums []int, one, three int) bool {
	if len(nums) < 1 {
		return false
	}
	left := 0
	right := len(nums) - 1
	midIndex := (left + right) / 2
	target := nums[midIndex]

	fmt.Printf("searching %d from nums: %v, checking if %d > %d and %d < %d?\n", target, nums, target, one, target, three)

	if right < left {
		return false
	}
	if right == left {
		return target > one && target < three
	}

	// [1,2,3,4,5,6,7,8] 2,5,...

	switch {
	case target > one && target < three:
		// fmt.Printf("found target = %d\n", target)
		return true
	case target > three:
		return bSearchBetween(nums[:midIndex], one, three)
	case target < one:
		return bSearchBetween(nums[midIndex+1:], one, three)
	}

	// fmt.Printf("found target = %d\n", target)
	return target > one && target < three
}
