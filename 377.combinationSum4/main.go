package main

import "fmt"

func main() {
	nums := []int{2, 1, 3}
	target := 35
	fmt.Println(combinationSum4(nums, target))
}

func combinationSum4(nums []int, target int) int {
	if target == 0 {
		return 1
	}
	if target < 0 {
		return 0
	}

	combs := 0
	for _, num := range nums {
		combs += combinationSum4(nums, target-num)
	}
	return combs
}
