package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoCitySchedCost([][]int{
		{259, 770},
		{448, 54},
		{840, 118},
		{926, 667},
		{184, 139},
		{557, 469},
	},
	))
}

type flyingCost [][]int

func (f flyingCost) Len() int {
	return len(f)
}

func (f flyingCost) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f flyingCost) Less(i, j int) bool {
	return f[i][0]-f[i][1] < f[j][0]-f[j][1]
}

func (f flyingCost) costs() int {
	sum := 0

	for i, cost := range f {
		if i < f.Len()/2 {
			sum += cost[0]
		} else {
			sum += cost[1]
		}
	}

	return sum
}

func twoCitySchedCost(costs [][]int) int {
	f := flyingCost(costs)
	sort.Sort(f)

	return f.costs()
}
