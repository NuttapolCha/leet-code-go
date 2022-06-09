package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) []int {
	result := make([]int, 0)
	for i, num := range numbers {
		toFind := target - num
		foundIndex := bSearch(numbers, toFind, i+1, len(numbers)-1)
		if foundIndex != -1 {
			result = append(result, []int{i + 1, foundIndex + 1}...)
			break
		}
	}
	return result
}

func bSearch(nums []int, target, left, right int) int {
	// not found at any index
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	if left == right {
		if nums[mid] == target {
			return mid
		}
		return -1
	}

	// found target
	if nums[mid] == target {
		return mid
	}

	// search the right side
	if nums[mid] < target {
		return bSearch(nums, target, mid+1, right)
	}

	// search the left side
	if nums[mid] > target {
		return bSearch(nums, target, left, mid)
	}

	return -1
}
