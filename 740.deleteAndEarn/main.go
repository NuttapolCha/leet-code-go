package main

import (
	"fmt"
)

func main() {
	// nums := []int{8, 3, 7, 4, 1, 8, 10, 2, 10, 10}
	nums := []int{6, 5, 10, 2, 8, 6, 6, 5, 2, 9, 9, 4, 6, 3, 3, 7, 7, 8, 9, 5}
	numCount, scoreMap := newCountMap(nums)
	fmt.Printf("numCount: %+v\nscoreMap: %+v\n\n", numCount, scoreMap)

	// checkPickOne(nums)

	maximumPoints := deleteAndEarn(nums)
	fmt.Printf("maximum earns = %d\n", maximumPoints)
}

func deleteAndEarn(nums []int) (cumulativePoints int) {
	if len(nums) <= 0 {
		return 0
	}
	choosen := pickOne(nums)
	fmt.Printf(">> start deleteAndEarn, deleting %d from %v\n", choosen, nums)

	// we delete choosen, we earn choosen
	numsAfterOperation := performOperation(nums, choosen)
	cumulativePoints += choosen

	cumulativePoints += deleteAndEarn(numsAfterOperation)
	return cumulativePoints
}

func performOperation(nums []int, num int) (numsAfterOperation []int) {
	left := num - 1
	right := num + 1

	nums = deleteNumFromList(nums, num, false)
	nums = deleteNumFromList(nums, left, true)
	nums = deleteNumFromList(nums, right, true)

	numsAfterOperation = nums
	return
}

func checkPickOne(nums []int) {
	_, scoreMap := newCountMap(nums)
	for num := range scoreMap {
		fmt.Printf("choosing %d will earns %d and lost %d\n", num, scoreMap[num], scoreMap[num-1]+scoreMap[num+1])
	}
	fmt.Printf("\n")
}

func pickOne(nums []int) (choosen int) {
	// initial choosen
	choosen = nums[0]

	_, scoreMap := newCountMap(nums)

	maxParticipateScore := 0
	maximumEarn := 0
	minimumLost := scoreMap.sumValue()

	for num := range scoreMap {
		participateScore := 0
		earn := scoreMap[num]
		lost := scoreMap[num-1] + scoreMap[num+1]

		if earn > maximumEarn {
			maximumEarn = earn
			participateScore++
		}

		if lost < minimumLost {
			minimumLost = lost
			participateScore++
		}

		if isIsolate(num, scoreMap) {
			participateScore += 10
		}

		if participateScore > maxParticipateScore {
			maxParticipateScore = participateScore
			choosen = num
		}
	}

	return
}

func deleteNumFromList(nums []int, num int, deleteAll bool) []int {
	ret := make([]int, 0, len(nums))

	numIsDeleted := false

	for _, e := range nums {
		if e != num || numIsDeleted {
			ret = append(ret, e)
		} else {
			if !deleteAll {
				numIsDeleted = true
			}
		}
	}

	return ret
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

func (c CountMap) sumValue() int {
	ret := 0
	for _, value := range c {
		ret += value
	}
	return ret
}

func isIsolate(num int, c CountMap) bool {
	for key := range c {
		if key == num+1 || key == num-1 {
			return false
		}
	}
	return true
}
