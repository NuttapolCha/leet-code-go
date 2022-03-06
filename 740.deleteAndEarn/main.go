package main

import (
	"fmt"
)

func main() {
	nums := []int{8, 3, 7, 4, 1, 8, 10, 2, 10, 10}
	// nums := []int{1, 2, 2, 3, 4, 4, 4, 5}
	// nums := []int{2, 2, 3, 3, 3, 4}

	maximumPoints := deleteAndEarn(nums)
	fmt.Printf("maximum earns = %d\n", maximumPoints)
}

func deleteAndEarn(nums []int) (cumulativePoints int) {
	if len(nums) <= 0 {
		return 0
	}
	choosen := pickOne(nums)
	// fmt.Printf(">> start deleteAndEarn, deleting %d from %v\n", choosen, nums)

	// we delete choosen, we earn choosen
	numsAfterOperation := performOperation(nums, choosen)
	cumulativePoints += choosen

	cumulativePoints += deleteAndEarn(numsAfterOperation)
	return cumulativePoints
}

func performOperation(nums []int, num int) (numsAfterOperation []int) {
	left := num - 1
	right := num + 1

	nums, _ = deleteNumFromList(nums, num, false)
	nums, _ = deleteNumFromList(nums, left, true)
	nums, _ = deleteNumFromList(nums, right, true)

	numsAfterOperation = nums
	return
}

func pickOne(nums []int) (choosen int) {
	// initial choosen
	choosen = nums[0]

	numCount, scoreMap := newCountMap(nums)

	for num := range numCount {

		left := num - 1
		right := num + 1

		lost := left*numCount[left] + right*numCount[right]
		earn := scoreMap[num]

		if lost < earn {
			choosen = num
		}
	}

	return
}

func deleteNumFromList(nums []int, num int, deleteAll bool) ([]int, int) {
	ret := make([]int, 0, len(nums))

	numIsDeleted := false
	sumOfDeletedNums := 0

	for _, e := range nums {
		if e != num || numIsDeleted {
			ret = append(ret, e)
		} else {
			if !deleteAll {
				numIsDeleted = true
			}
			sumOfDeletedNums += e
		}
	}

	return ret, sumOfDeletedNums
}

type CountMap map[int]int

// newCountMap return new occurances and scores maps
func newCountMap(nums []int) (CountMap, CountMap) {
	countMap := make(map[int]int, len(nums))
	scoreMap := make(map[int]int, len(nums))

	for _, num := range nums {
		countMap[num]++
	}

	for num, count := range countMap {
		scoreMap[num] = num * count
	}

	return countMap, scoreMap
}
