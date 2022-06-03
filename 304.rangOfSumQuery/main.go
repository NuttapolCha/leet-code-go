package main

import "fmt"

func main() {
	m := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
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
	preCalculate [][]int
}

// use 1D caching
func Constructor(matrix [][]int) NumMatrix {
	dp := make([][]int, len(matrix), len(matrix))

	// build an 1-D sum cache
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]), len(matrix[i]))
		for j := 0; j < len(dp[i]); j++ {
			if j-1 >= 0 { // has previous in a row
				dp[i][j] = matrix[i][j] + dp[i][j-1]
			} else { // not has prvious in a row
				dp[i][j] = matrix[i][j]
			}
		}
	}

	// fmt.Println("DP = ")
	for _, row := range dp {
		fmt.Println(row)
	}

	return NumMatrix{
		preCalculate: dp,
	}
}

func (matrix *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0

	for i := row1; i <= row2; i++ {
		if col1-1 >= 0 {
			sum += matrix.preCalculate[i][col2] - matrix.preCalculate[i][col1-1]
		} else {
			sum += matrix.preCalculate[i][col2]
		}
	}

	return sum
}
