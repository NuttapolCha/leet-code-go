package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	var ret int
	left := 0
	right := len(height) - 1

	for right > left {
		// calculate area
		area := (right - left) * min(height[left], height[right])
		if area > ret {
			ret = area
		}

		// decision should move left or right
		switch {
		case height[left] == max(height[left], height[right]):
			// left is max, we greedy choose move right
			right--
		case height[right] == max(height[left], height[right]):
			// right is max, we greedy chose move left
			left++
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
