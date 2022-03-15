package main

import "fmt"

func main() {
	out := twoSum([]int{3, 2, 4}, 6)
	fmt.Printf("index = %d and %d\n", out[0], out[1])
}

func twoSum(nums []int, target int) []int {
	ret := make([]int, 2, 2)

	// use each element to compare with others
	for i, num1 := range nums {

		// check for each element match with the first one
		for j, num2 := range nums {
			if i == j {
				continue // not selecting the same element
			}
			if num1+num2 == target {
				ret[0] = i
				ret[1] = j
				return ret
			}
		}
	}
	return ret
}
