package main

import "fmt"

func main() {
	// fmt.Println(longestPalindrome("a"))
	// fmt.Println(longestPalindrome("cbbbbbbbbdabcdbcaxyz"))
	fmt.Println(longestPalindrome("abcdedcba"))
}

func longestPalindrome(s string) string {
	longestSubStr := ""

	for i := 0; i < len(s); i++ {
		fmt.Printf("iteration: %d\n", i)
		subStr := string(s[i])

		left := i - 1
		right := i + 1

		// same char determined as palindrome
		// checking left side
		for left > -1 && s[left] == s[i] {
			subStr = s[left : i+1]
			fmt.Printf("left = current, s[%d] == s[%d] result in subStr = %s, left--\n", left, i, subStr)
			left--
		}

		// checking right side
		for right < len(s) && s[right] == s[i] {
			subStr = s[i : right+1]
			fmt.Printf("right = current, s[%d] == s[%d] result in subStr = %s, right++\n", right, i, subStr)
			right++
		}

		// no more same char
		// checking for the both side
		for left > -1 && right < len(s) && s[left] == s[right] {
			subStr = s[left : right+1]
			fmt.Printf("left = right, s[%d] == s[%d] result in subStr = %s, left-- and right++\n", left, right, subStr)
			left--
			right++
		}

		if len(subStr) > len(longestSubStr) {
			longestSubStr = subStr
		}
	}
	return longestSubStr
}
