package main

import (
	"fmt"
	"sort"
)

func main() {
	// stones := []int{2, 7, 4, 1, 8, 1}
	stones := []int{1, 3, 5, 6}
	fmt.Println(lastStoneWeight(stones))
	return

	sorted := SortIntByAsc(stones)
	sort.Sort(sorted)
	stones = []int(sorted)

	fmt.Printf("init stones = %v\n", stones)
	stones = destroy(stones, 2, 1)
	fmt.Printf("stones after = %v\n", stones)
}

type SortIntByAsc []int

func (a SortIntByAsc) Len() int           { return len(a) }
func (a SortIntByAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortIntByAsc) Less(i, j int) bool { return a[i] > a[j] }

func lastStoneWeight(stones []int) int {
	sorted := SortIntByAsc(stones)
	sort.Sort(sorted)
	stones = []int(sorted)

	fmt.Printf("stones = %+v\n", stones)

	for len(stones) > 1 {
		for i := 0; i < len(stones)-2; i++ {
			fmt.Printf("stones = %+v\n", stones)
			stones = destroy(stones, i+1, i) // because stones[i+1] always less than or equal to stones[i]
		}
	}

	return stones[0]
}

func destroy(stones []int, ix, iy int) (ret []int) {
	fmt.Printf("destroying (%d,%d) from stones: %+v\n", stones[ix], stones[iy], stones)
	switch {
	case stones[ix] > stones[iy]:
		// constraint violated
		ret = stones

	case stones[ix] == stones[iy]:
		ret = remove(stones, ix)
		ret = remove(ret, ix) // after remove ix at first, index shifted

	case stones[ix] < stones[iy]:
		left := stones[iy] - stones[ix]
		stones[iy] = left
		ret = remove(stones, ix)
	}

	return
}

func remove(arr []int, index int) (ret []int) {
	if index < len(arr)-1 {
		ret = append(arr[0:index], arr[index+1:]...)
	} else {
		ret = arr[0:index]
	}
	return
}
