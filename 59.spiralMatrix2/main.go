package main

import (
	"fmt"
)

func main() {
	// m := generateMatrix(6)
	// fmt.Println(m)
	// for _, row := range m {
	// 	for _, e := range row {
	// 		fmt.Printf("%d\t", e)
	// 	}
	// 	fmt.Printf("\n\n")
	// }
	directionMap(5)
}

type direction int

const (
	RIGHT direction = iota
	DOWN
	LEFT
	UP
)

func (d direction) next() direction {
	var ret direction
	switch d {
	case RIGHT:
		ret = DOWN
	case DOWN:
		ret = LEFT
	case LEFT:
		ret = UP
	case UP:
		ret = RIGHT
	}

	return ret
}

func generateMatrix(n int) [][]int {
	ret := make([][]int, n, n)

	// initial state
	d := RIGHT
	m := directionMap(n)
	i := 0
	j := 0

	for running := 1; running <= n*n; running++ {
		i, j = next(n, running, d, m)
		if len(ret[i]) == 0 {
			ret[i] = make([]int, n, n)
		}
		ret[i][j] = running
	}
	return ret
}

func next(n, running int, d direction, m map[int]direction) (i, j int) {
	if nextDirection, isTriggerred := m[running]; isTriggerred {
		d = nextDirection
	}
	switch d {
	case RIGHT:
		i = running
	case DOWN:
	case LEFT:

	case UP:
	}
	// running--
	// i = running / n
	// j = running % n
	// fmt.Printf("running = %d, i = %d, j = %d\n", running+1, i, j)
	return
}

func directionMap(n int) map[int]direction {
	ret := make(map[int]direction, n)

	delta := n - 1
	d := RIGHT

	count := 0
	limit := 3
	for running := 1; running <= n*n && delta != 0; running += delta {
		ret[running] = d
		fmt.Printf("n = %d => direction: %v\n", running, d)
		if count >= limit {
			limit = 2
			delta--
			count = 1
		} else {
			count++
		}
		d = d.next()
	}

	return ret
}
