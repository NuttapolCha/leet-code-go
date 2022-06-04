package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(transpose(matrix))
}

func transpose(matrix [][]int) [][]int {
	// len of ret's row should be len of matrix's column
	ret := make([][]int, len(matrix[0]), len(matrix[0]))

	for i, row := range matrix {
		for j := 0; j < len(row); j++ {
			if len(ret[j]) == 0 {
				// len of ret's column should be len of matrix's row
				ret[j] = make([]int, len(matrix), len(matrix))
			}
			ret[j][i] = row[j]
		}
	}

	return ret
}
