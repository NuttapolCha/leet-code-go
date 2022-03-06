package main

import (
	"fmt"
)

func main() {
	nums := []int{8, 7, 3, 8, 1, 4, 10, 10, 10, 2}

	maximumPoints := deleteAndEarn(nums)
	fmt.Printf("maximum earns = %d\n", maximumPoints)
}

func deleteAndEarn(nums []int) (cumulativePoints int) {
	if len(nums) <= 0 {
		return 0
	}
	choosen := pickOne(nums)
	fmt.Printf(">> start deleteAndEarn, deleting %d from %v\n", choosen, nums)

	numsAfterOperation, points := performOperation(nums, choosen)
	cumulativePoints += points

	cumulativePoints += deleteAndEarn(numsAfterOperation)
	return cumulativePoints
}

func performOperation(nums []int, num int) (resultNums []int, points int) {
	left := num - 1
	right := num + 1

	nums, points = deleteNumFromList(nums, num, false)
	nums, _ = deleteNumFromList(nums, left, true)
	nums, _ = deleteNumFromList(nums, right, true)

	resultNums = nums
	return
}

// array related

func pickOne(nums []int) (choosen int) {
	// initial choosen
	choosen = nums[0]

	numCount, minimumLost := newCountMap(nums)

	for num := range numCount {
		fmt.Printf("\tchecking if it worth to choose: %d in nums: %v\n", num, nums)
		// if we choose the current num, compare the points earned and lost
		var earned, lost1, lost2 int
		tmp := nums

		// points earned if we choose to delete this num
		tmp, earned = deleteNumFromList(tmp, num, false)
		fmt.Printf("\twill earns: %d points if choose: %d\n", earned, num)

		// points lost if we choose to delete this num
		tmp, lost1 = deleteNumFromList(tmp, num-1, true)
		_, lost2 = deleteNumFromList(tmp, num+1, true)
		currentLost := lost1 + lost2
		fmt.Printf("\twill lost: %d points if choose: %d\n", currentLost, num)

		// update minimumLost and choosen
		if currentLost < minimumLost {
			choosen = num
			minimumLost = currentLost
			fmt.Printf("\n\t>>> minimumLost was updated to %d, choosen was updated to %d\n", minimumLost, choosen)
		}
		fmt.Printf("\n\n")
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

// count map

type CountMap map[int]int

// newCountMap return a new CountMap and its sum
func newCountMap(nums []int) (CountMap, int) {
	ret := make(map[int]int, len(nums))
	sum := 0

	for _, num := range nums {
		ret[num]++
		sum += num
	}

	return ret, sum
}
