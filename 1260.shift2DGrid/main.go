package main

import "fmt"

func main() {
	original := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for _, row := range shiftGrid(original, 91) {
		fmt.Println(row)
	}
}

type MatrixElem map[int]int

func shiftGrid(grid [][]int, k int) [][]int {
	elems := make(MatrixElem)
	m := len(grid)
	var n int
	for i, row := range grid {
		n = len(row) // should be constant
		for j, e := range row {
			shiftedOneDIndex := (j + n*i + k)
			if shiftedOneDIndex > m*n-1 {
				shiftedOneDIndex %= m * n
			}
			elems[shiftedOneDIndex] = e
			fmt.Printf("(val: %d, index: %d)\t", e, shiftedOneDIndex)
		}
		fmt.Println("")
	}

	fmt.Printf("result matrix =>\n")

	index := 0
	ret := make([][]int, m, m)
	for i := range ret {
		ret[i] = make([]int, n, n)
		for j := range ret[i] {
			ret[i][j] = elems[index]
			fmt.Printf("(val: %d, index: %d)\t", elems[index], index)
			index++
		}
		fmt.Println("")
	}

	return ret
}
