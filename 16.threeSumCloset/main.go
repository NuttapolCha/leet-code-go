package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{-1, 2, 1, -4}
	// nums := []int{0, 1, 2}
	nums := []int{1, 1, 1, 0}
	target := -100

	fmt.Println(threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// sorted = [-4, -1, 1, 2]

	sumCloset := nums[len(nums)-1] + nums[len(nums)-2] + nums[len(nums)-3]
	minDiff := sumCloset + abs(target) // a maximum int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				diff := abs(sum - target)
				if diff < minDiff {
					// fmt.Printf("sumCloset has been updated at (i,j,k) = (%d,%d,%d) of %+v\n", i, j, k, nums)
					minDiff = diff
					sumCloset = sum
				}
			}
		}
	}

	return sumCloset
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
