package main

import "fmt"

func main() {
	testCases := [][]int{
		{0, 7, 1, 4, 6},
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{4, 2, 0, 3, 2, 5},
		{4, 2, 3},
	}

	for _, height := range testCases {
		fmt.Println(trap(height))
	}
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

		nextMax := 0
		nextMaxIndex := -1
		nextMin := 0
		nextMinIndex := -1

		// moving j until next wall
		for ; j < len(height) && height[j] < height[i]; j++ {
			// fmt.Printf("height[i=%d] = %d, height[j=%d] = %d\n", i, height[i], j, height[j])
			// assign next max and next min
			if height[j] > nextMax {
				nextMax = height[j]
				nextMaxIndex = j
			}
			if height[j] < nextMin {
				nextMin = height[j]
				nextMinIndex = j
			}

			unrealizedUnits += height[i] - height[j]
		}

		// no grater wall at another side
		hasWallAtAnotherSide := false
		if j >= len(height) {
			if nextMaxIndex > nextMinIndex {
				hasWallAtAnotherSide = true
				// recalculate unrealized units
				j = i + 1
				unrealizedUnits = 0

				for ; j < nextMaxIndex; j++ {
					unrealizedUnits += height[nextMaxIndex] - height[j]
				}
			}
		} else {
			hasWallAtAnotherSide = true
		}

		if hasWallAtAnotherSide {
			// fmt.Printf("units += %d\n", unrealizedUnits)
			units += unrealizedUnits
			i = j
		} else {
			i++
		}
	}
	return units
}
