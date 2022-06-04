package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "11100100101010100000000000000000011001010110010101001100100110"
	k := 16

	fmt.Println(hasAllCodes(s, k))
}

func hasAllCodes(s string, k int) bool {
	allBins := genAllBins(make([]string, 0), make([]byte, k, k), 0)

	for _, bin := range allBins {
		if !strings.Contains(s, bin) {
			fmt.Printf("biniary code: %s is of length %d and does not exist in %s\n", bin, k, s)
			return false
		}
	}

	return true
}

func genAllBins(curr []string, digits []byte, index int) []string {
	if index == len(digits) { // out of range. i.e. no more fill
		curr = append(curr, string(digits))
	} else {
		digits[index] = byte('0')
		curr = genAllBins(curr, digits, index+1)

		digits[index] = byte('1')
		curr = genAllBins(curr, digits, index+1)
	}

	return curr
}
