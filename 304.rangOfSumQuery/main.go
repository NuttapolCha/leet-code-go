package main

import "fmt"

func main() {
	m := [][]int{
		{3, 0, 1, 4, 2},
		{6, 5, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	a := Constructor(m)
	fmt.Println(a.SumRegion(2, 1, 4, 3))
	fmt.Println(a.SumRegion(1, 1, 2, 2))
	fmt.Println(a.SumRegion(1, 2, 2, 4))
	fmt.Println(a.SumRegion(0, 0, 4, 4))
}

type NumMatrix struct {
	martrix [][]int
}

// TODO: apply caching in constructor and we are able to
// do O(1) query time complexity
func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{
		martrix: matrix,
	}
}

func (matrix *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	m := matrix.martrix

	sum := 0
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			sum += m[i][j]
		}
	}

	return sum
}
