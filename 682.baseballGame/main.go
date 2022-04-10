package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
}

func calPoints(ops []string) int {
	sum := make([]int, 0, len(ops))

	for i, op := range ops {
		switch op {
		case "+":
			sum = append(sum, sum[len(sum)-1]+sum[len(sum)-2])
		case "C":
			sum = sum[0 : len(sum)-1]
		case "D":
			sum = append(sum, sum[len(sum)-1]*2)
		default:
			// i.e. is score
			score, err := strconv.Atoi(op)
			if err != nil {
				panic(err)
			}
			sum = append(sum, score)
		}
		fmt.Printf("sum after iteration: %d = %v\n", i, sum)
	}

	ret := 0
	for _, score := range sum {
		ret += score
	}

	return ret
}
