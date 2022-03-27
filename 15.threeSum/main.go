package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(input))
}

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)

	for i, num1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			num2 := nums[j]
			for k := j + 1; k < len(nums); k++ {
				num3 := nums[k]
				if num1+num2+num3 == 0 {
					toAppend := []int{num1, num2, num3}
					if !contains(ret, toAppend) {
						ret = append(ret, toAppend)
					}
				}
			}
		}
		// }
	}
	return ret
}

func contains(threeSums [][]int, s []int) bool {
	for _, eachResult := range threeSums {
		if isSame(eachResult, s) {
			return true
		}
	}
	return false
}

type intSorter []int

func (a intSorter) Len() int           { return len(a) }
func (a intSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a intSorter) Less(i, j int) bool { return a[i] < a[j] }

func isSame(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	sorted1 := intSorter(s1)
	sort.Sort(sorted1)

	sorted2 := intSorter(s2)
	sort.Sort(sorted2)

	for i := 0; i < len(s1); i++ {
		if sorted1[i] != sorted2[i] {
			return false
		}
	}

	return true
}
