package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	fmt.Println(maxProduct(words))
}

func maxProduct(words []string) int {
	max := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if !strings.ContainsAny(words[i], words[j]) {
				prod := len(words[i]) * len(words[j])
				if prod > max {
					max = prod
				}
			}
		}
	}
	return max
}
