package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoCitySchedCost([][]int{
		{259, 770},
		{448, 54},
		{926, 667},
		{184, 139},
		{840, 118},
		{557, 469},
	},
	))
}

type sortByA [][]int

func (a sortByA) Len() int           { return len(a) }
func (a sortByA) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByA) Less(i, j int) bool { return a[i][0] < a[j][0] }

type sortByB [][]int

func (a sortByB) Len() int           { return len(a) }
func (a sortByB) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByB) Less(i, j int) bool { return a[i][1] < a[j][1] }

func twoCitySchedCost(costs [][]int) int {
	sortedByA := sortByA(costs)
	sort.Sort(sortedByA)

	fmt.Printf("sorted by A: %+v\n", sortedByA)

	people := len(costs)
	cityCap := people / 2

	aAmount := 0
	for i, cost := range sortedByA {
		// Since we already sorted costs,
		// first half should go to city A, and the final half go to city B
		if i < cityCap {
			aAmount += cost[0]
		} else {
			aAmount += cost[1]
		}
	}

	sortedByB := sortByB(costs)
	sort.Sort(sortedByB)

	fmt.Printf("sorted by B: %+v\n", sortedByB)

	bAmount := 0
	for i, cost := range sortedByB {
		if i < cityCap {
			bAmount += cost[1]
		} else {
			bAmount += cost[0]
		}
	}

	return func() int {
		if bAmount < aAmount {
			return bAmount
		}
		return aAmount
	}()
}
