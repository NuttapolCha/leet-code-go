package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	// -4 -1 -1 0 1 2
	// input := func() []int {
	// 	return make([]int, 3000, 3000)
	// }()
	fmt.Println(threeSum(input))
}

type intSorter []int

func (a intSorter) Len() int           { return len(a) }
func (a intSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a intSorter) Less(i, j int) bool { return a[i] < a[j] }

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)

	sortedNums := intSorter(nums)
	sort.Sort(sortedNums)

	appended := make(map[string]bool)

	for i, n1 := range sortedNums {
		for j := i + 1; j < len(sortedNums); j++ {
			n2 := sortedNums[j]
			// n1 + n2 + n3 = 0
			n3 := -1 * (n1 + n2)
			if bSearch(sortedNums[j+1:], n3) {
				if !alreadyExist(appended, n1, n2, n3) {
					ret = append(ret, []int{n1, n2, n3})
				}
			}
		}

	}
	return ret
}

func bSearch(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	leftIndex := 0
	rightIndex := len(nums) - 1
	midIndex := (leftIndex + rightIndex) / 2

	if rightIndex == leftIndex { // e.g. == midIndex
		return nums[midIndex] == target
	}
	if rightIndex < leftIndex {
		return false
	}
	if target < nums[midIndex] {
		return bSearch(nums[0:midIndex], target)
	}
	if target > nums[midIndex] {
		// fmt.Printf("nums = %+v, midIndex = %d, len(nums) = %d\n", nums, midIndex, len(nums))
		return bSearch(nums[midIndex+1:], target)
	}

	return target == nums[midIndex]
}

func alreadyExist(appended map[string]bool, num1, num2, num3 int) bool {
	if appended == nil {
		return false
	}

	key := fmt.Sprintf("%d:%d:%d", num1, num2, num3)
	if !appended[key] {
		appended[key] = true
		return false
	}

	return true
}
