package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// nums := []int{4, 2, 3}
	fmt.Println(trap(nums))
}

func trap(height []int) int {
	// find starting index
	startIndex := 0
	for i, num := range height {
		if num > 0 {
			startIndex = i
			break
		}
	}

	// continue iteration from startIndex
	units := 0
	for i := startIndex; i < len(height); {
		j := i + 1
		unrealizedUnits := 0

		// moving j until next wall
		for ; j < len(height) && height[j] < height[i]; j++ {
			fmt.Printf("height[i=%d] = %d, height[j=%d] = %d\n", i, height[i], j, height[j])
			unrealizedUnits += height[i] - height[j]
		}

		// no wall at another side
		if j >= len(height) {
			i++
		} else {
			fmt.Printf("units += %d\n", unrealizedUnits)
			units += unrealizedUnits
			i = j
		}
	}
	return units
}
