package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}

	fmt.Println(topKFrequent(nums, 2))
}

type numCountPair struct {
	num   int
	count int
}

type SortNumCountPairByMostCount []*numCountPair

func (a SortNumCountPairByMostCount) Len() int           { return len(a) }
func (a SortNumCountPairByMostCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortNumCountPairByMostCount) Less(i, j int) bool { return a[i].count > a[j].count }

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)

	// O(n)
	for _, num := range nums {
		if _, ok := countMap[num]; !ok {
			countMap[num] = 1
		} else {
			countMap[num]++
		}
	}

	// O(n)
	numsCount := []*numCountPair{}
	for num, count := range countMap {
		numsCount = append(numsCount, &numCountPair{
			num:   num,
			count: count,
		})
	}

	// O(nlogn)
	sorted := SortNumCountPairByMostCount(numsCount)
	sort.Sort(sorted)

	// O(k)
	ret := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ret = append(ret, sorted[i].num)
	}

	return ret
}
