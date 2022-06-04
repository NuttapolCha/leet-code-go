package main

import (
	"fmt"
)

func main() {
	s := "00110"
	k := 2

	fmt.Println(hasAllCodes(s, k))
}

func hasAllCodes(s string, k int) bool {
	bins := make(map[string]bool)

	for i := 0; i < len(s)-k+1; i++ {
		bins[s[i:i+k]] = true
	}

	// 1 << k is the same as 2^k
	return len(bins) == 1<<k
}
