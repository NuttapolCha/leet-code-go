package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 1, 2, 4, 5}
	fmt.Println(sortArrayByParity(nums))
}

func sortArrayByParity(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		switch {
		case nums[i]%2 == 0 && nums[j]%2 == 0: // both are even
			return nums[i] < nums[j]
		case nums[i]%2 == 0: // only nums[j] is even
			return true
		case nums[j]%2 == 0: // only nums[i] is even
			return false
		default: // both are odd
			return nums[i] < nums[j]
		}
	})
	return nums
}
