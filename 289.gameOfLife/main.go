package main

import (
	"fmt"
)

func main() {
	input := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	fmt.Printf("before:\n")
	printArray(input)

	gameOfLife(input)

	fmt.Printf("after:\n")
	printArray(input)
}

func printArray(input [][]int) {
	for _, row := range input {
		for _, e := range row {
			fmt.Printf("%d\t", e)
		}
		fmt.Println("")
	}
}

func gameOfLife(board [][]int) {
	if len(board) == 0 {
		return
	}
	if len(board[0]) == 0 {
		return
	}

	var copied [][]int
	for i := 0; i < len(board); i++ {
		var row []int
		for j := 0; j < len(board[i]); j++ {
			row = append(row, board[i][j])
		}
		copied = append(copied, row)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = liveOrDie(copied, i, j)
		}
	}
}

// O(1) (n max = 8)
func liveOrDie(board [][]int, i, j int) int {
	hasAbove := i > 0
	hasBelow := i+1 < len(board)
	hasLeft := j > 0
	hasRight := j+1 < len(board[i])

	sum := 0

	// above row
	if hasAbove {
		if hasLeft {
			sum += board[i-1][j-1]
		}

		sum += board[i-1][j]

		if hasRight {
			sum += board[i-1][j+1]
		}
	}

	// left and right
	if hasLeft {
		sum += board[i][j-1]
	}
	if hasRight {
		sum += board[i][j+1]
	}

	// below row
	if hasBelow {
		if hasLeft {
			sum += board[i+1][j-1]
		}

		sum += board[i+1][j]

		if hasRight {
			sum += board[i+1][j+1]
		}
	}

	curr := board[i][j]
	switch {
	case curr == 1 && (sum < 2 || sum > 3):
		return 0
	case curr == 1 && sum == 2 || sum == 3:
		return 1
	case curr == 0 && sum == 3:
		return 1
	default:
		return curr
	}
}
