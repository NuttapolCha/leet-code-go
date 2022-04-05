package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	var ret int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			// guarantee that j always grater than i
			area := (j - i) * min(height[i], height[j])
			if area > ret {
				ret = area
			}
		}
	}

	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
