package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 5, 6, 7, 8, 8, 11, 23, 45}))
}

func findDuplicate(nums []int) int {
	countMap := make(map[int]int, len(nums))

	for _, num := range nums {
		if _, ok := countMap[num]; ok {
			return num
		} else {
			countMap[num] = 1
		}
	}

	return 0
}
