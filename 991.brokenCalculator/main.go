package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(brokenCalc(1, 1000000000))
}

func brokenCalc(startValue int, target int) int {
	count := 0
	for startValue != target {
		fmt.Printf("iteration: %d, use multiply: %v\n", count+1, shouldUseMultiply(startValue, target))
		if shouldUseMultiply(startValue, target) {
			startValue *= 2
		} else {
			startValue--
		}
		fmt.Printf("current value = %d\n", startValue)
		count++
		time.Sleep(1 * time.Second)
	}
	return count
}

func shouldUseMultiply(currValue, target int) bool {
	// var countIfUseMultiply, countIfUseMinus int
	if currValue*2 == target {
		return true
	}
	if currValue*2 > target {
		// countIfUseMultiply = currValue*2 - target
	}
	if (currValue-1)*2 > target {
		//
	}
	return false
}

// 1,000,000,000
// 536,870,912
