package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(123321))
}

func isPalindrome(x int) bool {
	s := fmt.Sprintf("%d", x)
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
