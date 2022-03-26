package main

import "fmt"

func main() {
	// input := []int{3, 5, 7, 8, 45, 67, 78, 98, 120}
	input := []int{5}

	result := search(input, 5)

	fmt.Printf("input = %+v\n", input)
	fmt.Println(result)
}

func search(nums []int, target int) int {
	return bSearch(nums, target, 0, len(nums)-1)
}

func bSearch(nums []int, target, leftIndex, rightIndex int) int {
	// fmt.Printf("searching %d start from %d to %d\n", target, nums[leftIndex], nums[rightIndex])

	mid := (leftIndex + rightIndex) / 2
	midValue := nums[mid]

	switch {
	case rightIndex == leftIndex:
		if midValue == target {
			return mid
		}
		return -1
	case rightIndex < leftIndex:
		// not fonud
		return -1
	case target == midValue:
		// found
		return mid

		// try again
	case target < midValue:
		return bSearch(nums, target, leftIndex, mid)
	case target > midValue:
		return bSearch(nums, target, mid+1, rightIndex)
	}

	return -1
}
