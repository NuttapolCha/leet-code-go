package main

import "fmt"

func main() {
	// nums := []int{4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4}
	nums := []int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}
	// nums := []int{1, 2, 3, 4}
	fmt.Println(maxOperations(nums, 3))
}

func maxOperations(nums []int, k int) int {
	// construct count map
	countMap := make(map[int]int, len(nums))
	for _, num := range nums {
		if _, ok := countMap[num]; !ok {
			countMap[num] = 1
		} else {
			countMap[num]++
		}
	}

	// fmt.Printf("countMap before counting operations:\n")
	// sameSequence := []int{}
	// for num, count := range countMap {
	// 	sameSequence = append(sameSequence, num)
	// 	fmt.Printf("%d => %d\n", num, count)
	// }

	// start find num operations
	// fmt.Printf("matched operations\n")
	operations := 0
	for _, num := range nums {
		countMap[num]-- // always decrease when picking a num

		// no more usable num left
		if countMap[num] < 0 {
			countMap[num]++
			continue
		}

		// find target
		target := k - num
		if count, ok := countMap[target]; ok && count > 0 {
			// fmt.Printf("\t%d + %d = %d\n", num, target, k)
			operations++
			countMap[target]--
		} else {
			countMap[num]++ // increase available usable num back if not found target
		}
	}

	// fmt.Printf("countMap after counting operations:\n")
	// for _, num := range sameSequence {
	// 	fmt.Printf("%d => %d\n", num, countMap[num])
	// }
	return operations
}
