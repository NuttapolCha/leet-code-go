package main

import "fmt"

func main() {
	// s := "baaba" // expected 2 "baab" -> "a"
	s := "babababbababab"
	fmt.Println(removePalindromeSub(s))
}

func removePalindromeSub(s string) int {
	if isPalindrome(s) {
		return 1
	}

	return 2
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
