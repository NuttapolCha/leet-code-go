package main

import (
	"fmt"
)

func main() {
	input := "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"
	//  rever := "abcboa"
	fmt.Println(validPalindrome(input))
}

func validPalindrome(s string) bool {
	return isPalindrome([]byte(s), false)
}

func isPalindrome(chars []byte, isSkipped bool) bool {
	left := 0
	right := len(chars) - 1

	for left < right {
		if chars[left] != chars[right] {
			if isSkipped {
				return false
			} else {
				return isPalindrome(chars[left:right], true) ||
					isPalindrome(func() []byte {
						ret := make([]byte, 0, right-left+1)
						for i := right; i > left; i-- {
							ret = append(ret, chars[i])
						}
						return ret
					}(), true)
			}
		}
		left++
		right--
	}

	return true
}
