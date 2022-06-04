package main

import "fmt"

func main() {
	printAllBins(3)
}

func printAllBins(digits int) {
	d := make([]int, digits, digits)
	bCounter(d, 0)
}

// slice d already have k filled
func bCounter(d []int, k int) {
	if k == len(d) {
		fmt.Println(d)
	} else {
		d[k] = 0
		bCounter(d, k+1)

		d[k] = 1
		bCounter(d, k+1)
	}
}
