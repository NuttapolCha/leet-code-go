package main

import "fmt"

func main() {
	for _, row := range generate(5) {
		fmt.Println(row)
	}
}

func generate(numRows int) [][]int {
	ret := [][]int{{1}}

	if numRows == 1 {
		return ret
	}

	for i := 1; i < numRows; i++ {
		nextRow := make([]int, i+1)
		nextRow[0] = 1
		nextRow[i] = 1

		for j := 1; j < len(nextRow)-1; j++ {
			nextRow[j] = ret[i-1][j-1] + ret[i-1][j]
		}
		ret = append(ret, nextRow)
	}

	return ret
}
