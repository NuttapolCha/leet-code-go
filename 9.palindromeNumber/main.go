package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome(123321))
}

func isPalindrome(x int) bool {
	s := fmt.Sprintf("%d", x)

	chars := strings.Split(s, "")
	backward := ""

	for i := len(chars) - 1; i >= 0; i-- {
		backward += chars[i]
	}

	return backward == s
}
