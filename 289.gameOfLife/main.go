package main

import "fmt"

func main() {
	input := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	fmt.Printf("before:\n")
	for _, row := range input {
		for _, e := range row {
			fmt.Printf("%d\t", e)
		}
		fmt.Println("")
	}

	gameOfLife(input)

	fmt.Printf("after:\n")
	for _, row := range input {
		for _, e := range row {
			fmt.Printf("%d\t", e)
		}
		fmt.Println("")
	}
}

func gameOfLife(board [][]int) {
	numRows := len(board)

	copy := board
	for i, row := range board {
		numCols := len(row)
		for j, _ := range row {
			switch {
			case i == 0: // first row
				switch {
				case j == 0:
					if i+1 < numRows && j+1 < numCols {
						board[i][j] = liveOrDie(copy[i+1][j], copy[i+1][j+1], copy[i][j+1])
					} else if i+1 < numRows && j+1 >= numCols {
						board[i][j] = liveOrDie(copy[i+1][j])
					} else if i+1 >= numRows && j+1 < numCols {
						board[i][j] = liveOrDie(copy[i][j+1])
					}
				case j == len(row)-1:
				default:
				}
			case i == len(board)-1: // last row
				switch {
				case j == 0:
				case j == len(row)-1:
				default:
				}
			default: // not first and not last row
				switch {
				case j == 0:
				case j == len(row)-1:
				default:
				}
			}
		}
	}
}

// O(1) (n max = 8)
func liveOrDie(neighbors ...int) int {
	sum := 0
	for _, num := range neighbors {
		sum += num
	}

	switch {
	case sum < 2 || sum > 3:
		return 0
	case sum == 2 || sum == 3:
		return 1
	default:
		panic("this should not be occurred")
	}
}
