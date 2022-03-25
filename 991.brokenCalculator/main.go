package main

import (
	"fmt"
)

func main() {
	fmt.Println(brokenCalc(1, 1000000000))
}

func brokenCalc(startValue int, target int) int {
	count := 0
	for startValue != target {
		target = decision(startValue, target)
		// fmt.Printf("count: %d, target = %d\n", count, target)
		count++
		// time.Sleep(1 * time.Second)
	}

	return count
}

func decision(startValue, target int) int {
	if target < startValue { // because divide cannot reach start value which is greater than target
		return target + 1
	}
	if target%2 == 0 {
		return target / 2
	}
	return target + 1
}

// 1,000,000,000
// 536,870,912
