package main

import (
	"fmt"
	"strings"
)

func main() {
	num := uint32(11)
	fmt.Println(hammingWeight(num))
}

func hammingWeight(num uint32) int {
	bi := strings.Split(toBinary(num), "")

	count := 0
	for _, c := range bi {
		if c == "1" {
			count++
		}
	}

	return count
}

func toBinary(num uint32) string {
	reverse := ""
	for num != 0 {
		r := num % 2
		num /= 2
		reverse += fmt.Sprintf("%d", r)
	}
	revChars := strings.Split(reverse, "")

	bi := ""
	for i := len(revChars) - 1; i >= 0; i-- {
		bi += revChars[i]
	}

	return bi
}
